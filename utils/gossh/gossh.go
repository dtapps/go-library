package gossh

import (
	"context"
	"io"
	"log/slog"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

// SshConfig 配置结构体，仅支持密码认证。
type SshConfig struct {
	Username   string // SSH 登录用户名
	Password   string // SSH 登录密码（明文，注意安全）
	ServerAddr string // SSH 服务器地址，格式: "host:port"，例如 "192.168.1.10:22"
	RemoteAddr string // 远程目标地址（在 SSH 服务器内部可访问），例如 "localhost:3306"
	LocalAddr  string // 本地监听地址，例如 "127.0.0.1:8080"
}

// Ssh 表示一个 SSH 隧道实例。
type Ssh struct {
	config *SshConfig
}

// NewSsh 创建一个新的 SSH 隧道实例。
func NewSsh(config *SshConfig) *Ssh {
	return &Ssh{config: config}
}

// Tunnel 启动本地端口转发隧道。
// 它会在 LocalAddr 上监听连接，并将每个连接通过 SSH 转发到 RemoteAddr。
// 函数会一直运行，直到 context 被取消或发生不可恢复错误。
func (s *Ssh) Tunnel(ctx context.Context) error {
	slog.Info("正在初始化 SSH 隧道配置...",
		slog.String("用户名", s.config.Username),
		slog.String("SSH服务器", s.config.ServerAddr),
		slog.String("本地监听地址", s.config.LocalAddr),
		slog.String("远程目标地址", s.config.RemoteAddr),
	)

	// 构建 SSH 客户端配置（仅密码认证）
	sshConfig := &ssh.ClientConfig{
		User: s.config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.config.Password),
		},
		Timeout: 30 * time.Second,
		// ⚠️ 警告：跳过主机密钥验证！仅用于测试环境。
		// 生产环境应验证服务器指纹以防止中间人攻击。
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		ClientVersion:   "SSH-Tunnel/Go",
	}

	// 在本地绑定监听地址
	localListener, err := net.Listen("tcp", s.config.LocalAddr)
	if err != nil {
		slog.Error("❌ 无法在本地启动监听器",
			slog.String("本地地址", s.config.LocalAddr),
			slog.Any("错误", err),
		)
		return err
	}
	defer func() {
		slog.Info("🔌 本地监听器已关闭")
		localListener.Close()
	}()

	slog.Info("✅ SSH 隧道已成功启动并开始监听",
		slog.String("本地地址", s.config.LocalAddr),
	)

	// 监听 context 取消信号，用于优雅退出
	go func() {
		<-ctx.Done()
		slog.Info("🛑 收到退出信号，正在关闭监听器...")
		localListener.Close()
	}()

	// 主循环：接受本地连接
	for {
		localConn, err := localListener.Accept()
		if err != nil {
			// 如果是 context 取消导致的 Accept 失败，属于正常退出
			if ctx.Err() != nil {
				slog.Info("👋 SSH 隧道已正常停止")
				return ctx.Err()
			}
			slog.Error("❌ 接受本地连接时发生错误",
				slog.Any("错误", err),
			)
			continue
		}

		// 获取本地连接的远程地址（即客户端地址）
		clientAddr := localConn.RemoteAddr().String()
		slog.Info("📥 收到新的本地连接",
			slog.String("客户端地址", clientAddr),
			slog.String("本地监听地址", s.config.LocalAddr),
		)

		// 为每个连接启动独立的转发 goroutine
		go s.forwardConnection(localConn, sshConfig, clientAddr)
	}
}

// forwardConnection 负责处理单个连接的双向数据转发。
func (s *Ssh) forwardConnection(localConn net.Conn, sshConfig *ssh.ClientConfig, clientAddr string) {
	defer func() {
		localConn.Close()
		slog.Debug("📤 本地连接已关闭", slog.String("客户端地址", clientAddr))
	}()

	// 步骤1: 连接到 SSH 服务器
	slog.Debug("📡 正在连接 SSH 服务器...", slog.String("服务器地址", s.config.ServerAddr))
	sshClient, err := ssh.Dial("tcp", s.config.ServerAddr, sshConfig)
	if err != nil {
		slog.Error("💥 无法连接到 SSH 服务器",
			slog.String("服务器地址", s.config.ServerAddr),
			slog.String("客户端地址", clientAddr),
			slog.Any("错误详情", err),
		)
		return
	}
	defer func() {
		sshClient.Close()
		slog.Debug("🔌 SSH 客户端连接已关闭", slog.String("服务器地址", s.config.ServerAddr))
	}()

	slog.Info("✅ 成功连接到 SSH 服务器",
		slog.String("服务器地址", s.config.ServerAddr),
		slog.String("客户端地址", clientAddr),
	)

	// 步骤2: 通过 SSH 隧道连接到远程目标
	slog.Debug("➡️ 正在通过 SSH 隧道连接远程目标...",
		slog.String("远程地址", s.config.RemoteAddr),
		slog.String("客户端地址", clientAddr),
	)
	remoteConn, err := sshClient.Dial("tcp", s.config.RemoteAddr)
	if err != nil {
		slog.Error("💥 无法通过 SSH 隧道连接远程目标",
			slog.String("远程地址", s.config.RemoteAddr),
			slog.String("客户端地址", clientAddr),
			slog.Any("错误详情", err),
		)
		return
	}
	defer func() {
		remoteConn.Close()
		slog.Debug("🔚 远程目标连接已关闭",
			slog.String("远程地址", s.config.RemoteAddr),
			slog.String("客户端地址", clientAddr),
		)
	}()

	slog.Info("🔗 隧道已建立：本地 ↔ SSH服务器 ↔ 远程目标",
		slog.String("客户端地址", clientAddr),
		slog.String("本地地址", s.config.LocalAddr),
		slog.String("远程地址", s.config.RemoteAddr),
	)

	// 步骤3: 启动双向数据转发

	// 方向1: 本地 → 远程
	go func() {
		n, err := io.Copy(remoteConn, localConn)
		slog.Debug("📤 本地 → 远程 数据传输结束",
			slog.String("客户端地址", clientAddr),
			slog.Int64("传输字节数", n),
		)
		if err != nil && err != io.EOF {
			slog.Warn("⚠️ 本地 → 远程 传输异常",
				slog.String("客户端地址", clientAddr),
				slog.Any("错误", err),
			)
		}
		// 触发对方关闭
		localConn.Close()
		remoteConn.Close()
	}()

	// 方向2: 远程 → 本地
	go func() {
		n, err := io.Copy(localConn, remoteConn)
		slog.Debug("📥 远程 → 本地 数据传输结束",
			slog.String("客户端地址", clientAddr),
			slog.Int64("传输字节数", n),
		)
		if err != nil && err != io.EOF {
			slog.Warn("⚠️ 远程 → 本地 传输异常",
				slog.String("客户端地址", clientAddr),
				slog.Any("错误", err),
			)
		}
		// 触发对方关闭
		localConn.Close()
		remoteConn.Close()
	}()
}
