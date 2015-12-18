package cdn

import (
	"fmt"
	"io"
	"os"
)

type LocalProvider struct {
	Root string
}

func (p *LocalProvider) Write(dir, name string, fn func(wrt io.Writer) error) error {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", p.Root, dir), 0700); err != nil {
		return err
	}
	fd, err := os.OpenFile(fmt.Sprintf("%s/%s/%s", p.Root, dir, name), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	if err = fn(fd); err != nil {
		return err
	}
	return fd.Sync()
}
