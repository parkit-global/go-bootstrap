package gocli

import (
	"os"
	"os/exec"
)

type Mod struct {
	Name string
	Dir  string
}

func (m *Mod) Init() error {
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

	cmd := exec.Command("go", "mod", "init", m.Name)
	err = cmd.Run()
	if err != nil {
		return err
	}

	err = os.Chdir(currentDir)

	return err
}

func (m *Mod) Tidy() error {
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

	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	if err != nil {
		return err
	}

	err = os.Chdir(currentDir)

	return err
}
