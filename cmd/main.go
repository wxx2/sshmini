package main

import "sshmini/lib"

func main() {
	var sshinfo = lib.SshInfo{
		"172.17.6.57:22",
		"root",
		"ceiec",
	}
	var sshinfo1 = lib.SshInfo{
		"172.17.6.61:22",
		"root",
		"ceiec",
	}
	go lib.NewTerminal(sshinfo)
	lib.NewTerminal(sshinfo1)
}
