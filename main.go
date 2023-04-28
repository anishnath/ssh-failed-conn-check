package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"golang.org/x/crypto/ssh"
)

func main() {

	var days int
	var host string
	var username string
	flag.IntVar(&days, "days", 7, "number of days to search for failed SSH connections")
	flag.StringVar(&host, "host", "", "SSH host to connect to")
	flag.StringVar(&username, "username", "ubuntu", "SSH username")
	flag.Parse()

	if host == "" {
		cmd := fmt.Sprintf("sudo journalctl _SYSTEMD_UNIT=ssh.service --since '%d days ago' | sudo grep 'Connection closed by authenticating user'", days)
		out, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			log.Fatalf("unable to execute command locally: %v", err)
		}
		fmt.Printf("%s", out)
		return
	}

	key, err := ioutil.ReadFile("/Users/anish/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatalf("unable to connect to remote server: %v", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("unable to create SSH session: %v", err)
	}
	defer session.Close()

	cmd := fmt.Sprintf("sudo journalctl _SYSTEMD_UNIT=ssh.service --since '%d days ago' | sudo grep 'Connection closed by authenticating user'", days)
	out, err := session.Output(cmd)
	if err != nil {
		log.Fatalf("unable to execute command on remote server: %v", err)
	}

	fmt.Printf("%s", out)
}
