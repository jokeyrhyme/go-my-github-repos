package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

// GetCurrentUsername returns a string based on a call to GitHub's API
func GetCurrentUsername(ctx context.Context, client *github.Client) (string, error) {
	me, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return "", fmt.Errorf("error requesting GitHub user: %v", err)
	}
	if me.GetLogin() == "" {
		return "", fmt.Errorf("unexpected GitHub user data: %v", err)
	}
	return me.GetLogin(), nil
}
