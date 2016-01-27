package raru

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ArtemKulyabin/bre/ldd"
	"github.com/ArtemKulyabin/yax/osx"
)

func MkJail(name string, paths []string) (err error) {
	for _, path := range paths {
		path, err = exec.LookPath(path)
		if err != nil {
			return
		}
		newpath := filepath.Join(name, path)
		err = os.MkdirAll(filepath.Dir(newpath), 0770)
		if err != nil {
			return fmt.Errorf("mkdir() failed: %s", err)
		}

		if _, err = os.Stat(newpath); os.IsNotExist(err) {
			err = osx.CopyFile(newpath, path)
			if err != nil {
				return
			}
			libs, err := ldd.GetDynLibs(path)
			if err != nil {
				return err
			}
			libs = append(libs, ldd.GetDynLoader())
			for _, lib := range libs {
				newpath = filepath.Join(name, lib)
				if _, err = os.Stat(newpath); os.IsNotExist(err) {
					err = os.MkdirAll(filepath.Dir(newpath), 0770)
					if err != nil {
						return fmt.Errorf("mkdir() failed: %s", err)
					}
					err = osx.CopyFile(newpath, lib)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return
}
