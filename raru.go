package raru

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/ArtemKulyabin/yax/osx"
	"github.com/ArtemKulyabin/yax/syscallx"
)

type Executor struct {
	name   string
	args   []string
	id     int
	chroot string
	path   string
}

func NewExecutor() (exer *Executor, err error) {
	id, err := RandomID()
	if err != nil {
		return
	}
	return &Executor{id: id}, nil
}

func (exer *Executor) Exec(name string, arg ...string) (err error) {
	exer.name = name
	exer.args = arg
	return exer.execute()
}

func (exer *Executor) SetChrootDir(path string) {
	exer.chroot = path
}

func (exer *Executor) Spawn(cmd *exec.Cmd) (err error) {
	exer.Prepare(cmd)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Not able to execute: %s", err)
	}
	return
}

func (exer *Executor) Prepare(cmd *exec.Cmd) (err error) {
	exer.path = cmd.Path
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	if exer.chroot != "" {
		if cmd.Dir == "" {
			cmd.Dir = "/"
		}
		cmd.SysProcAttr.Chroot = exer.chroot
		if err = exer.MkJail(); err != nil {
			return
		}
	}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(exer.id), Gid: uint32(exer.id)}
	return
}

func (exer *Executor) execute() (err error) {
	path, err := exec.LookPath(exer.name)
	if err != nil {
		return fmt.Errorf("Executable not found: %s", err)
	}
	exer.path = path

	env := os.Environ()

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getcwd() failed: %s", err)
	}

	if exer.chroot != "" {
		if err = exer.MkJail(); err != nil {
			return
		}
		if err = os.Chdir(exer.chroot); err != nil {
			return fmt.Errorf(`Unable to chdir("/"): %s`, err)
		}
		if err = syscall.Chroot("."); err != nil {
			return
		}
	}

	err = syscallx.Setgid(exer.id)
	if err != nil {
		return fmt.Errorf("Unable to setgid(), aborting: %s", err)
	}
	err = syscallx.Setuid(exer.id)
	if err != nil {
		return fmt.Errorf("Unable to setuid(), aborting: %s", err)
	}

	if exer.chroot == "" {
		err = os.Chdir(cwd)
		if err != nil {
			err = os.Chdir("/")
			if err != nil {
				return fmt.Errorf(`Unable to chdir("/"): %s`, err)
			}
		}
	}

	err = syscall.Exec(exer.path, append([]string{exer.name}, exer.args...), env)
	if err != nil {
		return fmt.Errorf("Not able to execute: %s", err)
	}
	return
}

func (exer *Executor) MkJail() (err error) {
	err = MkJail(exer.chroot, []string{exer.path})
	if err != nil {
		return
	}
	// Very important! Apply recursive chown for jail.
	// Only current random user has access to jail.
	// suid and sgid binaries have no effect for break security.
	err = osx.Chown(exer.chroot, exer.id, exer.id)
	return
}
