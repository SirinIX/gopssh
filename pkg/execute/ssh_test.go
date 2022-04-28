package execute

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
	"time"

	"golang.org/x/crypto/ssh"
)

func TestSSH(t *testing.T) {
	addr := "172.16.8.83:22"
	cfg := &ssh.ClientConfig{
		User:    "root",
		Auth:    []ssh.AuthMethod{ssh.Password("Abc!@#135")},
		Timeout: 5 * time.Second,
		Config: ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	cli, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()

	session, err := cli.NewSession()
	if err != nil {
		t.Error(err)
	}
	defer session.Close()

	out, err := session.CombinedOutput("ls -l")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(out))
}

func TestSSHAsync(t *testing.T) {
	addr := "172.16.8.83:22"
	cfg := &ssh.ClientConfig{
		User:    "root",
		Auth:    []ssh.AuthMethod{ssh.Password("Abc!@#135")},
		Timeout: 5 * time.Second,
		// Config: ssh.Config{
		// 	// Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		// },
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	cli, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		t.Error(err)
	}
	defer cli.Close()

	session, err := cli.NewSession()
	if err != nil {
		t.Error(err)
	}
	defer session.Close()

	//
	stdout, err := session.StdoutPipe()
	if err != nil {
		t.Error(err)
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		t.Error(err)
	}
	if err := session.Start("ls -l /lll; sleep 5; echo xxxxx"); err != nil {
		t.Error(err)
	}
	// doneout := make(chan bool, 1)
	// doneerr := make(chan bool, 1)
	// go func() {
		// fmt.Println("stderr >")
		res, err := readPipe(stderr)
		if err != nil {
			t.Error(err)
		}
		fmt.Println("stderr: \n" + string(res))
		// doneerr <- true
	// }()
	// go func() {
		// fmt.Println("stdout >")
		resout, err := readPipe(stdout)
		if err != nil {
			t.Error(err)
		}
		fmt.Println("stdout: \n" + string(resout))
		// doneout <- true
	// }()
	// <-doneerr
	// <-doneout
	session.Wait()
}

func readPipe(pipe io.Reader) ([]byte, error) {
	r := bufio.NewReader(pipe)
	// for {
	// 	line, _, err := r.ReadLine()
	// 	if line == nil {
	// 		return nil
	// 	}
	// 	if err != nil {
	// 		return nil
	// 	}
	// 	fmt.Println("out: \n" + string(line))
	// }


	// res, err := ioutil.ReadAll(r)
	// fmt.Println(string(res))

	// return err

	return ioutil.ReadAll(r)
}
