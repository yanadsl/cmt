package cmd

import (
	"fmt"
	"net/url"
)

type Cmd interface {
	Run(name string, args ...string) (string, string, error)
	Start(name string, args ...string) (Cmd, error)
	Wait() error
	Output(name string, args ...string) (string, string, error)
	URL(path string) *url.URL
}

func Scp(src, dest *url.URL) error {
	scpCmd := NewLocal()
	_, _, err := scpCmd.Run("scp", "-3", formatCopyURL(src), formatCopyURL(dest))

	return err
}

func Rsync(src, dest *url.URL) error {
	rsyncCmd := NewLocal()
	_, _, err := rsyncCmd.Run("rsync", "-r", "--size-only", formatCopyURL(src), formatCopyURL(dest))

	return err
}

func formatCopyURL(u *url.URL) string {
	if u.Host == "" {
		return u.String()
	}
	return fmt.Sprintf("%s@%s:%s", u.User.Username(), u.Host, u.Path)
}
