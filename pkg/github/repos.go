package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

/*
ForEachOwnPublicRepository pages through GitHub responses,
and calls `handleRepo()` for each eligible repository,
which are public and owned by me instead of an organisation
*/
func ForEachOwnPublicRepository(
	ctx context.Context,
	client *github.Client,
	handleRepo func(*github.Repository),
) error {
	username, err := GetCurrentUsername(ctx, client)
	if err != nil {
		return err
	}

	opt := &github.RepositoryListOptions{
		Affiliation: "owner",
		Visibility:  "public",
	}
	for {
		repos, resp, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			return fmt.Errorf("error listing GitHub repositories: %v", err)
		}

		for _, repo := range repos {
			if repo.GetPrivate() ||
				repo.Organization != nil ||
				repo.Owner.GetLogin() != username {
				continue // ensure Affiliation and Visibility options
			}
			handleRepo(repo)
		}

		if resp.NextPage == 0 {
			break // no more pages
		}
		opt.Page = resp.NextPage
	}
	return nil
}
