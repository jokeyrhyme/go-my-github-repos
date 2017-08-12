package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/jokeyrhyme/go-my-github-repos/pkg/config"
)

func main() {
	cfg := &config.Config{}

	err := cfg.Read()
	if err != nil {
		fmt.Printf("error reading config: %v", err)
		os.Exit(1)
	}

	if cfg.GithubToken == "" {
		fmt.Println(`
Create a GitHub token here:
https://github.com/settings/tokens
`)

		fmt.Print("Enter GitHub token: ")
		input, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			fmt.Printf("error reading password: %v", err)
			os.Exit(1)
		}
		fmt.Println()

		token := strings.Trim(string(input), " ")
		cfg.GithubToken = token
	} else {
		fmt.Println("using previous GitHub token")
	}

	err = cfg.Write()
	if err != nil {
		fmt.Printf("error writing config: %v", err)
		os.Exit(1)
	}
}
