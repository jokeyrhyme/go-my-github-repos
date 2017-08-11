# go-my-github-repos [![Build Status](https://travis-ci.org/jokeyrhyme/go-my-github-repos.svg?branch=master)](https://travis-ci.org/jokeyrhyme/go-my-github-repos)

teaching myself Go whilst tidying and standardising my GitHub repos


## Usage

```sh
$ my-github-repos --help
```


## Development

```sh
make setup # to install any required tooling

make test
make build
```


## Roadmap

- [ ] use [golang/dep](https://github.com/golang/dep) for dependency management

- [ ] prompt for GitHub API token and persist to ~/.config/jokeyrhyme/go-my-github-repos/config.toml

- [ ] list all GitHub repos that I have access to

- [ ] identify GitHub repos that are public and in my namespace

- [ ] ensure hook for [facebook/mention-bot](https://github.com/facebook/mention-bot) is installed
