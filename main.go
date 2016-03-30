package main

import (
	"errors"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]

	is_gui := false
	g_pos, err := findOption(args, "-g")
	if err == nil {
		is_gui = true
		args = remove(args, g_pos)
	}

	is_pipe := false
	_, err = os.Stdin.Stat()
	if err == nil {
		is_pipe = true
	}
	has_dash := false
	_, err = findOption(args, "-")
	if err == nil {
		has_dash = true
	}
	if is_pipe && !has_dash {
		// if "--" exists,
		// appending "-" at the last is cause losing control.
		args = append([]string{"-"}, args...)
	}

	app := "vim.exe"
	if is_gui {
		app = "gvim.exe"
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
