package raru

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"os/exec"
	"syscall"
)

func Spawn(cmd *exec.Cmd) (err error) {
	id, err := RandomID()
	if err != nil {
		return
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(id), Gid: uint32(id)}
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Not able to execute: %s", err)
	}
	return
}

func Exec(name string, arg ...string) (err error) {
	id, err := RandomID()
	if err != nil {
		return
	}

	path, err := exec.LookPath(name)
	if err != nil {
		return fmt.Errorf("Not able to execute: %s", err)
	}

	env := os.Environ()

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getcwd() failed: %s", err)
	}

	err = Setgid(id)
	if err != nil {
		return fmt.Errorf("Unable to setgid(), aborting: %s", err)
	}
	err = Setuid(id)
	if err != nil {
		return fmt.Errorf("Unable to setuid(), aborting: %s", err)
	}

	err = os.Chdir(cwd)
	if err != nil {
		err = os.Chdir("/")
		if err != nil {
			return fmt.Errorf(`Unable to chdir("/"): %s`, err)
		}
	}

	err = syscall.Exec(path, append([]string{name}, arg...), env)
	if err != nil {
		return fmt.Errorf("Not able to execute: %s", err)
	}
	return
}

func RandomID() (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint16))
	if err != nil {
		return 0, fmt.Errorf("Random generator error: %s", err)
	}
	return 31337 + int(r.Int64()), nil
}
