package main

import "sshmini/lib"

func main() {
	sshinfo := lib.SshInfo{
		"172.17.6.57:22",
		"root",
		"ceiec",
	}
	lib.NewTerminal(sshinfo)
}
