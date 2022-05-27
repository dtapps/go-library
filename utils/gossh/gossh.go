package gossh

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"time"
)

// SshConfig 配置
type SshConfig struct {
	Username   string
	Password   string
	ServerAddr string
	RemoteAddr string
	LocalAddr  string
}

type Ssh struct {
	SshConfig
}

// NewSsh 实例化
func NewSsh(config *SshConfig) *Ssh {
	app := &Ssh{}
	app.Username = config.Username
	app.Password = config.Password
	app.ServerAddr = config.ServerAddr
	app.RemoteAddr = config.RemoteAddr
	app.LocalAddr = config.LocalAddr
	return app
}

func (app *Ssh) Tunnel() {

	// 设置log配置
	log.Printf("%s，服务器：%s；远程：%s；本地：%s\n", "设置SSH配置", app.ServerAddr, app.RemoteAddr, app.LocalAddr)

	config := &ssh.ClientConfig{
		User: app.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(app.Password),
		},
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// 设置本地监听器
	localListener, err := net.Listen("tcp", app.LocalAddr)
	if err != nil {
		log.Printf("设置本地监听器失败: %v\n", err)
	}

	for {
		// 设置本地
		localConn, err := localListener.Accept()
		if err != nil {
			log.Printf("设置本地失败: %v\n", err)
		}
		go sForward(app.ServerAddr, app.RemoteAddr, localConn, config)
	}
}

// 转发
func sForward(serverAddr string, remoteAddr string, localConn net.Conn, config *ssh.ClientConfig) {

	// 设置sshClientConn
	sshClientConn, err := ssh.Dial("tcp", serverAddr, config)
	if err != nil {
		log.Printf("连接失败: %s", err)
	}

	// 设置Connection
	sshConn, err := sshClientConn.Dial("tcp", remoteAddr)

	// 将localConn.Reader复制到sshConn.Writer
	go func() {
		_, err = io.Copy(sshConn, localConn)
		if err != nil {
			log.Printf("复制失败: %v", err)
		}
	}()

	// 将sshConn.Reader复制到localConn.Writer
	go func() {
		_, err = io.Copy(localConn, sshConn)
		if err != nil {
			log.Printf("复制失败: %v", err)
		}
	}()
}
