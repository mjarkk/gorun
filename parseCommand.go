package main

import (
	"bytes"
	"errors"
	"strings"
)

func (c *Config) parseCommand(program string) (parsedCommand, error) {
	command := (*c)[program]

	output := parsedCommand{
		env:  []string{},
		args: []string{},
	}

	buff := bytes.NewBuffer(nil)
	isInsideQoutes := ""
	isEscaped := false
	i := 0

	writeBuff := func() {
		defer buff.Reset()
		s := buff.String()
		if len(s) == 0 {
			return
		}
		if output.program == "" && strings.Contains(s, "=") {
			output.env = append(output.env, s)
		} else if output.program == "" {
			output.program = s
		} else {
			output.args = append(output.args, s)
		}
	}

	for {
		if len(command) == i {
			break
		}

		returnEscape := isEscaped

		char := command[i]
		switch char {
		case '\\':
			if isEscaped {
				buff.WriteByte(char)
			} else {
				isEscaped = true
			}
		case '`', '"', '\'':
			if isInsideQoutes == string(char) && !isEscaped {
				isInsideQoutes = ""
			} else if isInsideQoutes == "" && !isEscaped {
				isInsideQoutes = string(char)
			} else {
				buff.WriteByte(char)
			}
		case ' ', '\t', '\n':
			if isInsideQoutes != "" || isEscaped {
				buff.WriteByte(char)
			} else {
				writeBuff()
			}
		default:
			buff.WriteByte(char)
		}

		if returnEscape {
			isEscaped = false
		}

		i++
	}
	writeBuff()

	if output.program == "" {
		return output, errors.New("No source to run defined in launch command")
	}

	return output, nil
}
