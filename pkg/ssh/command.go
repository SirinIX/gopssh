package ssh

import (
	"bufio"
	"io"
	"io/ioutil"
)

func (s *SSH) Command(cmd string) (*SSHResult, error) {
	session, err := s.NewSSHSession()
	if err != nil {
		return nil, err
	}

	// Run the command
	stdout, err := session.StdoutPipe()
	if err != nil {
		s.Logger.Error("failed to get stdout pipe, error: %v", err)
		return nil, err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		s.Logger.Error("failed to get stderr pipe, error: %v", err)
		return nil, err
	}
	if err := session.Start(cmd); err != nil {
		s.Logger.Error("failed to start command, error: %v", err)
		return nil, err
	}

	// Read from stdout, stderr
	sshRes := &SSHResult{
		Addr:    s.Address.String(),
		Command: cmd,
	}
	go func() {
		stderrRes, err := readFromPipe(stderr)
		if err != nil {
			s.Logger.Error("read stderr error: %s", err)
		}
		sshRes.Stderr = string(stderrRes)
	}()
	go func() {
		stdoutRes, err := readFromPipe(stdout)
		if err != nil {
			s.Logger.Error("read stderr error: %s", err)
		}
		sshRes.Stdout = string(stdoutRes)
	}()
	if err := session.Wait(); err != nil {
		s.Logger.Error("wait error: %s", err)
	}

	return sshRes, nil
}

func readFromPipe(pipe io.Reader) ([]byte, error) {
	rd := bufio.NewReader(pipe)
	return ioutil.ReadAll(rd)
}
