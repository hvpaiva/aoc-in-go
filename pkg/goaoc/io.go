package goaoc

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

type Env struct {
	Stdin  io.Reader
	Stdout io.Writer
	Args   []string
}

var defaultConsoleEnv = Env{
	Stdin:  os.Stdin,
	Stdout: os.Stdout,
	Args:   os.Args[1:],
}

type DefaultConsoleManager struct {
	Env Env
}

func (m DefaultConsoleManager) Input(part *int) error {
	getters := []func(part *int) error{
		m.Env.getPartInFlag,
		m.Env.getPartInEnv,
		m.Env.getPartInStdin,
	}

	for _, get := range getters {
		err := get(part)
		if err != nil {
			return err
		}

		if *part != 0 {
			break
		}
	}

	return ErrIORead{Err: ErrMissingPart}
}

func (m DefaultConsoleManager) OutputResults(result int) error {
	if _, err := fmt.Fprintf(m.Env.Stdout, "The challenge result is %d\n", result); err != nil {
		return ErrIOWrite{Err: err}
	}

	toClipboard(strconv.Itoa(result), m.Env.Stdout)

	return nil
}

func (e Env) getPartInFlag(part *int) error {
	fs := flag.NewFlagSet("goaoc", flag.ContinueOnError)
	fs.SetOutput(e.Stdout)

	fs.Usage = func() {
		_, _ = fmt.Fprintf(fs.Output(), "Usage: %s [options]\n", fs.Name())

		fs.PrintDefaults()
	}

	fs.IntVar(part, "part", 0, "Part of the challenge, valid values are (1/2)")

	if err := fs.Parse(e.Args); err != nil {
		return ErrIORead{Err: err}
	}

	return nil
}

func (e Env) getPartInEnv(part *int) error {
	env := os.Getenv("GOAOC_CHALLENGE_PART")

	if env == "" {
		return nil
	}

	if p, err := strconv.Atoi(env); err != nil {
		return ErrInvalidPartType
	} else {
		*part = p
	}

	return nil
}

func (e Env) getPartInStdin(part *int) error {
	_, err := fmt.Fprintln(e.Stdout, "Which part do you want to run? (1/2)")
	if err != nil {
		return ErrIORead{Err: err}
	}

	_, err = fmt.Fscanln(e.Stdin, &part)
	if err != nil && errors.Is(err, io.EOF) {
		return ErrIORead{Err: ErrMissingPart}
	}

	return nil
}

func toClipboard(value string, stdout io.Writer) {
	envVar := os.Getenv("GOAOC_DISABLE_COPY_CLIPBOARD")
	if envVar == "true" {
		return
	}

	c := clipboard.New()
	if err := c.CopyText(value); err != nil {
		_, _ = fmt.Fprintf(stdout, "Error copying to clipboard: %s\n", err)

		return
	}

	_, _ = fmt.Fprintf(stdout, "Copied to clipboard: %s\n", value)
}
