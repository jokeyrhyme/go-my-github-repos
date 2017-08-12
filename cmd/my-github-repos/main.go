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
	cfg, err := config.NewConfig("")
	if err != nil {
		fmt.Printf("error initialising config: %v", err)
		os.Exit(1)
	}

	cfg.Read()

	if cfg.GithubToken == "" {
		fmt.Println(`
Create a GitHub token here:
https://github.com/settings/tokens
`)

		fmt.Print("Enter GitHub token: ")
		var input []byte // avoid declaring a shadow `err` below
		input, err = terminal.ReadPassword(syscall.Stdin)
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
