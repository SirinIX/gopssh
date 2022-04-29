package ssh

import (
	"bufio"
	"gopssh/log"
	"io"
	"io/ioutil"
)

func (s *SSH) Command(cmd string) (*SSHResult, error) {
	session, err := s.NewSSHSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// Run the command
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Error("failed to get stdout pipe, error: %v", err)
		return nil, err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		log.Error("failed to get stderr pipe, error: %v", err)
		return nil, err
	}
	if err := session.Start(cmd); err != nil {
		log.Error("failed to start command, error: %v", err)
		return nil, err
	}

	// Read from stderr
	stderrDone := make(chan bool, 1)
	stderrRes := ""
	go func() {
		pipeRes, err := readFromPipe(stderr)
		if err != nil {
			log.Error("read stderr error: %s", err)
		}
		stderrRes = string(pipeRes)
		stderrDone <- true
	}()
	<-stderrDone
	
	// Read from stdout
	stdoutDone := make(chan bool, 1)
	stdoutRes := ""
	go func() {
		pipeRes, err := readFromPipe(stdout)
		if err != nil {
			log.Error("read stderr error: %s", err)
		}
		stdoutRes = string(pipeRes)
		stdoutDone <- true
	}()
	<-stdoutDone

	if err := session.Wait(); err != nil {
		log.Error("wait error: %s", err)
	}

	return &SSHResult{
		Address: s.Address,
		Command: cmd,
		Stdout:  stdoutRes,
		Stderr:  stderrRes,
	}, nil
}

func readFromPipe(pipe io.Reader) ([]byte, error) {
	rd := bufio.NewReader(pipe)
	return ioutil.ReadAll(rd)
}
