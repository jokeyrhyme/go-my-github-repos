# go-my-github-repos

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

- [ ] Travis CI integration

- [ ] `make test` exits with an error code if [golang/lint](https://github.com/golang/lint) or [`go vet`](https://golang.org/src/cmd/vet/README) identify problems

- [ ] `make fmt` uses [sqs/goreturns](https://github.com/sqs/goreturns), and `make test` triggers `make fmt`

- [ ] use [golang/dep](https://github.com/golang/dep) for dependency management

- [ ] prompt for GitHub API token and persist to ~/.config/jokeyrhyme/go-my-github-repos/config.toml

- [ ] list all GitHub repos that I have access to

- [ ] identify GitHub repos that are public and in my namespace

- [ ] ensure hook for [facebook/mention-bot](https://github.com/facebook/mention-bot) is installed
