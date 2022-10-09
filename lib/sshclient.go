package lib

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

type SshInfo struct {
	Address  string
	Username string
	Password string
}

func NewTerminal(sshinfo SshInfo) {
	// 设置配置信息
	config := ssh.ClientConfig{
		User:            sshinfo.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(sshinfo.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 与服务器建立连接
	clinet, err := ssh.Dial("tcp", sshinfo.Address, &config)
	if err != nil {
		log.Fatal("SSH dial error: %s", err.Error())
	}

	// 创建一个会话
	session, err := clinet.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatal("new session error: %s", err.Error())
	}

	// 设置Terminal Mode
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     //关闭回显
		ssh.TTY_OP_ISPEED: 14400, //设置传输速率
		ssh.TTY_OP_OSPEED: 14400,
	}

	// 设置输入输出
	session.Stdout = os.Stdout // 会话输出关联到系统标准输出设备
	session.Stdin = os.Stdin   // 会话错误输出关联到系统标准错误输出设备
	session.Stderr = os.Stderr // 会话输入关联到系统标准输入设备

	// 请求伪终端
	if err = session.RequestPty("linux", 32, 160, modes); err != nil {
		log.Fatal("Request pty error: %s", err.Error())
	}

	if err = session.Shell(); err != nil {
		log.Fatal("start shell error: %s", err.Error())
	}
	if err = session.Wait(); err != nil {
		log.Fatal("return error: %s", err.Error())
	}
}
