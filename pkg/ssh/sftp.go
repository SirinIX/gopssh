package ssh

import "github.com/pkg/sftp"

func (s *SSH) NewSftpClient() (*sftp.Client, error) {
	cli, err := s.NewSSHClient()
	if err != nil {
		return nil, err
	}

	return sftp.NewClient(cli)
}
