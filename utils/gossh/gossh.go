package gossh

import (
	"context"
	"io"
	"log/slog"
	"net"
	"sync"
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
func (s *Ssh) Tunnel(ctx context.Context) error {
	slog.Info("正在初始化 SSH 隧道配置...",
		slog.String("用户名", s.config.Username),
		slog.String("SSH服务器", s.config.ServerAddr),
		slog.String("本地监听地址", s.config.LocalAddr),
		slog.String("远程目标地址", s.config.RemoteAddr),
	)

	sshConfig := &ssh.ClientConfig{
		User:    s.config.Username,
		Auth:    []ssh.AuthMethod{ssh.Password(s.config.Password)},
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// 启动本地监听
	localListener, err := net.Listen("tcp", s.config.LocalAddr)
	if err != nil {
		slog.Error("无法在本地启动监听器",
			slog.String("本地地址", s.config.LocalAddr),
			slog.Any("错误", err),
		)
		return err
	}
	defer func() {
		slog.Info("🔌 本地监听器已关闭")
		localListener.Close()
	}()

	slog.Info("✅ SSH 隧道已启动",
		slog.String("本地地址", s.config.LocalAddr),
	)

	go func() {
		<-ctx.Done()
		slog.Info("🛑 收到退出信号，正在关闭监听器...")
		localListener.Close()
	}()

	for {
		localConn, err := localListener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				slog.Info("👋 SSH 隧道已正常停止")
				return nil
			}
			slog.Error("❌ 接受本地连接时发生错误",
				slog.Any("错误", err),
			)
			continue
		}

		clientAddr := localConn.RemoteAddr().String()
		slog.Info("📥 收到新的本地连接",
			slog.String("客户端地址", clientAddr),
			slog.String("本地监听地址", s.config.LocalAddr),
		)

		go s.forwardConnection(localConn, sshConfig, clientAddr)
	}
}

// forwardConnection 负责单个连接的数据转发。
func (s *Ssh) forwardConnection(localConn net.Conn, sshConfig *ssh.ClientConfig, clientAddr string) {
	defer localConn.Close()

	// Step 1: 建立 SSH 连接
	sshClient, err := ssh.Dial("tcp", s.config.ServerAddr, sshConfig)
	if err != nil {
		slog.Error("💥 无法连接 SSH 服务器",
			slog.String("服务器地址", s.config.ServerAddr),
			slog.String("客户端地址", clientAddr),
			slog.Any("错误", err),
		)
		return
	}
	defer sshClient.Close()

	// Step 2: 建立远程连接
	remoteConn, err := sshClient.Dial("tcp", s.config.RemoteAddr)
	if err != nil {
		slog.Error("💥 无法通过 SSH 隧道连接远程目标",
			slog.String("远程地址", s.config.RemoteAddr),
			slog.String("客户端地址", clientAddr),
			slog.Any("错误", err),
		)
		return
	}
	defer remoteConn.Close()

	slog.Info("🔗 隧道建立成功：本地 ↔ SSH服务器 ↔ 远程目标",
		slog.String("客户端地址", clientAddr),
		slog.String("本地地址", s.config.LocalAddr),
		slog.String("远程地址", s.config.RemoteAddr),
	)

	// Step 3: 双向转发 + 安全关闭（防止重复关闭）
	var wg sync.WaitGroup
	closeOnce := func() {
		localConn.Close()
		remoteConn.Close()
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		n, err := io.Copy(remoteConn, localConn)
		slog.Debug("📤 本地 → 远程 传输结束",
			slog.String("客户端地址", clientAddr),
			slog.Int64("传输字节数", n),
			slog.Any("错误", err),
		)
		closeOnce()
	}()

	go func() {
		defer wg.Done()
		n, err := io.Copy(localConn, remoteConn)
		slog.Debug("📥 远程 → 本地 传输结束",
			slog.String("客户端地址", clientAddr),
			slog.Int64("传输字节数", n),
			slog.Any("错误", err),
		)
		closeOnce()
	}()

	wg.Wait()
	slog.Debug("🔚 隧道连接关闭完成", slog.String("客户端地址", clientAddr))
}
