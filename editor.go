package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func pickEditor() string {
	if e := os.Getenv("EDITOR"); e != "" {
		return e
	}
	if e := os.Getenv("VISUAL"); e != "" {
		return e
	}
	cands := []string{"vim", "nano"}
	if runtime.GOOS == "windows" {
		cands = append(cands, "notepad")
	}
	for _, c := range cands {
		if p, err := exec.LookPath(c); err == nil {
			return p
		}
	}
	return "vim"
}

func editorCmd(path string) *exec.Cmd {
	ed := pickEditor()
	cmd := exec.Command(ed, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func runEditor(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	cmd := editorCmd(path)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("editor %s: %w", cmd.Path, err)
	}
	return nil
}
