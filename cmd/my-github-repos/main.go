package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"

	"github.com/jokeyrhyme/go-my-github-repos/pkg/config"
	gh "github.com/jokeyrhyme/go-my-github-repos/pkg/github"
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

Be sure to enable:
-   repo.public_repo
-   admin:repo_hook
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
		cfg.IsDirty = true
	} else {
		fmt.Println("using previous GitHub token")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	err = gh.ForEachOwnPublicRepository(ctx, client, handleRepo)
	if err != nil {
		fmt.Printf("error iterating over repositories: %v", err)
		os.Exit(1)
	}

	if cfg.IsDirty {
		err = cfg.Write()
		if err != nil {
			fmt.Printf("error writing config: %v", err)
			os.Exit(1)
		}
	}
}

func handleRepo(repo *github.Repository) {
	if repo.GetFork() {
		return // only want to process non-fork repos
	}
	fmt.Printf("%v\n", repo.GetFullName())
}
