package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"syscall"
)

func Shell(cmd string, args ...string) error {
	bin, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}
	return syscall.Exec(bin, append([]string{cmd}, args...), os.Environ())
}

func PkgRoot(o interface{}) string {
	return fmt.Sprintf("%s/src/%s", os.Getenv("GOPATH"), reflect.TypeOf(o).Elem().PkgPath())
}

func FuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func Mkdirs(d string, m os.FileMode) error {
	fi, err := os.Stat(d)
	if err == nil {
		if fi.IsDir() {
			return nil
		}
		return errors.New(fmt.Sprintf("%s is a file", d))
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(d, m)
	}
	return err

}
