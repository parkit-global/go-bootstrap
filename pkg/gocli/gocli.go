package gocli

import (
	"runtime"
	"strings"

	"os"
	"os/exec"
)

type Mod struct {
	Name string
	Dir  string
}

func (m *Mod) Version() string {
	goVersion := runtime.Version()
	version := strings.TrimPrefix(goVersion, "go")

	parts := strings.Split(version, ".")
	majorMinor := strings.Join(parts[:2], ".")

	return majorMinor
}

func (m *Mod) Init() error {
	return m.chdir(func() error {
		cmd := exec.Command("go", "mod", "init", m.Name)
		return cmd.Run()
	})
}

func (m *Mod) Tidy() error {
	return m.chdir(func() error {
		cmd := exec.Command("go", "mod", "tidy")
		return cmd.Run()
	})
}

func (m *Mod) chdir(doWork func() error) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.MkdirAll(m.Dir, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Chdir(m.Dir)
	if err != nil {
		return err
	}

	err = doWork()

	err = os.Chdir(currentDir)

	return err

}
