package test

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

const (
	address  = "172.17.6.57"
	username = "root"
	password = "ceiec"
)

func main() {
	// 设置配置信息
	config := ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 与服务器建立连接
	clinet, err := ssh.Dial("tcp", address, &config)
	if err != nil {
		log.Println(err)
		return
	}

	// 创建一个会话
	session, _ := clinet.NewSession()
	defer session.Close()

	// 设置Terminal Mode
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     //关闭回显
		ssh.TTY_OP_ISPEED: 14400, //设置传输速率
		ssh.TTY_OP_OSPEED: 14400,
	}

	// 请求伪终端
	err = session.RequestPty("linux", 32, 160, modes)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置输入输出
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	session.Shell()
	session.Wait()
}
