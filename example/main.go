package main

import (
	"github.com/dxwy2001/sshexec"
	"log"
	"time"
)

func main() {
	sshExecAgent := sshexec.SSHExecAgent{}
	sshExecAgent.Worker = 10
	sshExecAgent.TimeOut = time.Duration(120) * time.Second
	//s, err := sshExecAgent.SftpHostByKey([]string{"193.168.2.1","193.168.2.2"}, 22,"root", "/data/test.log", "/data/test.log")
	//log.Println("res:",s)
	//log.Println("err:",err)

	s1, err1 := sshExecAgent.SshHostByKey([]string{"172.16.10.126:8093", "172.16.10.125:8093"}, 22, "root", "ls -l")
	log.Println("res:", s1)
	log.Println("err:", err1)
	for {
		time.Sleep(1000)
	}
}
