package main

import (
	"os"
)

func main() {
	programs := getConfig()

	errChan := make(chan error)
	defer close(errChan)

	for programName := range programs {
		go func(programName string) {
			errChan <- programs.Exec(programName)
		}(programName)
	}

	count := 0
	for {
		count++
		err := <-errChan
		if count == len(programs) {
			os.Exit(0)
		}
		if err != nil {
			os.Exit(1)
		}
	}
}
