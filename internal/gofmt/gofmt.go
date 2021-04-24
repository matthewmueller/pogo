package gofmt

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

// Format the output using goimports
func Format(input string) (output string, err error) {
	goimports, err := exec.LookPath("goimports")
	if err != nil {
		return "", err
	}

	cmd := exec.Command(goimports)
	stdin, err := cmd.StdinPipe()

	if err != nil {
		return output, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return output, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return output, err
	}

	reader := bytes.NewBufferString(input)

	if e := cmd.Start(); e != nil {
		return output, e
	}

	io.Copy(stdin, reader)
	stdin.Close()

	formatted, err := ioutil.ReadAll(stdout)
	if err != nil {
		return output, err
	}

	formattingError, err := ioutil.ReadAll(stderr)
	if err != nil {
		return output, err
	}

	stderr.Close()
	stdout.Close()

	if e := cmd.Wait(); e != nil {
		return output, errors.New(string(formattingError))
	}

	return string(formatted), nil
}

// FormatAll formats by a path overwriting the original file
func FormatAll(dir string) (err error) {
	cmd := exec.Command("goimports", "-w", dir)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	defer stderr.Close()

	if e := cmd.Start(); e != nil {
		return fmt.Errorf("error running goimports: %w", err)
	}

	formattingError, err := ioutil.ReadAll(stderr)
	if err != nil {
		return fmt.Errorf("stderr error: %w: %s", err, formattingError)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error running goimports: %w: %s", err, formattingError)
	}

	return nil
}
