package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Config contains information about what to run
type Config map[string]string

func getConfig() Config {
	data, err := ioutil.ReadFile("./.gorun")
	if err != nil {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, last := path.Split(wd)
		secondaryFileOption := path.Clean(path.Join(wd, "..", "."+last))
		data, err = ioutil.ReadFile(secondaryFileOption)
		if os.IsNotExist(err) {
			fmt.Println("No config file found, create .gorun or ../." + last + " to use this program")
			os.Exit(1)
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var fullConfig map[string]Config
	err = json.Unmarshal(data, &fullConfig)
	if err != nil {
		fmt.Println("Invalid config file, error:", err)
		os.Exit(1)
	}

	configToUse := ""
	if len(os.Args) > 1 {
		configToUse = os.Args[len(os.Args)-1]
	}

	config, ok := fullConfig[configToUse]
	if !ok {
		if configToUse == "" {
			fmt.Println("could not find default config, a default config will look something like this")
			fmt.Println("{")
			fmt.Println("  \"\": {")
			fmt.Println("    \"App\": \"path/to/your/go/program/.\"")
			fmt.Println("  }")
			fmt.Println("}")
		} else {
			fmt.Println("Config not found, available configs:")
			for configName := range fullConfig {
				if configName == "" {
					configName = "[default]"
				}
				fmt.Print(configName + " ")
			}
			fmt.Println()
		}
		os.Exit(1)
	}

	if config == nil || len(config) == 0 {
		fmt.Println("No programs to run, exiting..")
		os.Exit(1)
	}

	return config
}
