package cdn

import (
	"io"
)

type Provider interface {
	Write(dir, name string, fn func(io.Writer) error) error
}
