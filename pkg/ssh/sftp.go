package ssh

import (
	"fmt"
	"gopssh/pkg/file"
	"path/filepath"

	"github.com/pkg/sftp"
)

const (
	kilobyte  = 1024
	megabytes = 1024 * 1024
)

func (s *SSH) CopyFile(localPath, remotePath string) error {
	sftpCli, err := s.NewSftpClient()
	if err != nil {
		return err
	}
	defer sftpCli.Close()

	localFile, err := file.OpenFile(localPath)
	if err != nil {
		return err
	}
	defer localFile.Close()
	localStat, _ := localFile.Stat()
	total := int(localStat.Size())

	// Return, if remote file exists
	if _, err := sftpCli.Stat(remotePath); err == nil {
		e := fmt.Errorf("remote file already exists")
		s.Logger.Error("remote file %s already exists, error: %v", remotePath, e)
		return e
	}

	remoteDir := filepath.Dir(remotePath)
	if err := sftpCli.MkdirAll(remoteDir); err != nil {
		s.Logger.Error("failed to create remote dir %s error: %v", remoteDir, err)
		return err
	}

	remoteFile, err := sftpCli.Create(remotePath)
	if err != nil {
		s.Logger.Error("failed to create remote file, error: %v", err)
		return err
	}
	defer remoteFile.Close()

	// Transfer
	s.Logger.Info("start to upload file %s", remotePath)
	buf := make([]byte, 10*megabytes)
	current := 0
	for {
		n, _ := localFile.Read(buf)
		if n == 0 {
			break
		}
		length, _ := remoteFile.Write(buf[0:n])
		// log
		current += length
		s.Logger.Info("transfer size: %.2f KB, %.2f %%", integerDivision(current, kilobyte), integerDivision(current, total)*100)
	}
	s.Logger.Info("succeed to upload file %s", remotePath)

	return nil
}

func (s *SSH) NewSftpClient() (*sftp.Client, error) {
	cli, err := s.NewSSHClient()
	if err != nil {
		return nil, err
	}

	return sftp.NewClient(cli)
}

func integerDivision(a, b int) float64 {
	return float64(a) / float64(b)
}
