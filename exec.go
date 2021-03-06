package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"unicode/utf8"
)

type printer struct {
	name string
	buff bytes.Buffer
	lock sync.Mutex
}

func (s *printer) Write(p []byte) (n int, err error) {
	parts := strings.Split(string(p), "\n")
	for i, part := range parts {
		if len(part) > 0 {
			part = "[" + s.name + "] " + part
		}
		parts[i] = part
	}

	fmt.Print(strings.Join(parts, "\n"))

	s.lock.Lock()
	n, err = s.buff.Write(p)
	s.lock.Unlock()
	return
}

func (s *printer) String() (out string) {
	s.lock.Lock()
	out = s.buff.String()
	s.lock.Unlock()
	return
}

type parsedCommand struct {
	program string
	args    []string
	env     []string
	isNotGo bool // If the command is prefixed with $ it will be executed as a normal command
}

// Exec executes a program defined in the config
func (c *Config) Exec(programName string) error {
	parsedCommand, err := c.parseCommand(programName)
	if err != nil {
		fmt.Println("Can't parse start argument, err:", err)
	}

	fullCommand := append([]string{
		"go",
		"run",
		parsedCommand.program,
	}, parsedCommand.args...)
	if parsedCommand.isNotGo {
		fullCommand = fullCommand[2:]
	}

	cmd := exec.Command(fullCommand[0], fullCommand[1:]...)
	cmd.Env = append(os.Environ(), parsedCommand.env...)

	output := &printer{name: programName}
	cmd.Stderr = output
	cmd.Stdout = output

	err = cmd.Run()
	if err != nil {
		return errors.New(output.String())
	}

	return nil
}

// scanWordsWithNewLines is a copy of bufio.ScanWords but this also captures new lines
// For specific comments about this function take a look at: bufio.ScanWords
func scanWordsWithNewLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isSpace(r) {
			return i + width, data[start:i], nil
		}
	}
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}

// isSpace is also copied from the bufio package and has been modified to also captures new lines
// For specific comments about this function take a look at: bufio.isSpace
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		switch r {
		case ' ', '\t', '\v', '\f':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}
