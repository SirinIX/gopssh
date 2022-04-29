package ssh

import (
	"gopssh/pkg/base64"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	protocol = "tcp"
	timeout  = 5
)

func (s *SSH) NewSSHSession() (*ssh.Session, error) {
	cli, err := s.NewSSHClient()
	if err != nil {
		return nil, err
	}
	// defer cli.Close()

	session, err := cli.NewSession()
	if err != nil {
		s.Logger.Error("failed to new session, error: %v", err)
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		s.Logger.Error("failed to request pty xterm, error: %v", err)
		return nil, err
	}

	return session, nil
}

func (s *SSH) NewSSHClient() (*ssh.Client, error) {
	auth, err := s.newSSHAuth()
	if err != nil {
		return nil, err
	}

	sshCfg := &ssh.ClientConfig{
		User:    s.Username,
		Auth:    auth,
		Timeout: timeout * time.Second,

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshCli, err := ssh.Dial(protocol, s.Address.String(), sshCfg)
	if err != nil {
		s.Logger.Error("failed to connect, error: %s", err)
		return nil, err
	}

	return sshCli, nil
}

func (s *SSH) newSSHAuth() ([]ssh.AuthMethod, error) {
	pwd, err := base64.Decode(s.Password)
	if err != nil {
		return nil, err
	}

	return []ssh.AuthMethod{ssh.Password(pwd)}, nil
}
