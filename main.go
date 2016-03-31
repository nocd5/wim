package main

import (
	"errors"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	app := "vim.exe"

	if g_pos, err := findOption(args, "-g"); err == nil {
		args = remove(args, g_pos)
		app = "gvim.exe"
	}

	is_pipe := false
	if _, err := os.Stdin.Stat(); err == nil {
		is_pipe = true
	}
	has_dash := false
	if _, err := findOption(args, "-"); err == nil {
		has_dash = true
	}
	if is_pipe && !has_dash {
		// if "--" exists,
		// appending "-" at the last is cause losing control.
		args = append([]string{"-"}, args...)
	}

	cmd := exec.Command(app, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func findOption(a []string, e string) (int, error) {
	for i, _e := range a {
		if _e == "--" {
			break
		} else if _e == e {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func remove(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}
