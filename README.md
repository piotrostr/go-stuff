# go stuff

me studying the art of gopher

## great resources

- [Effective Go](https://go.dev/doc/effective_go)

  a nice read, outdated a notch but was quite informative

- [Uber Go Guide](https://github.com/uber-go/guide/blob/master/style.md)

  excellent, a lot of good principles and common pitfalls

## keeping the code clean

- coc-go for nvim, based on gopls

```sh
go install golang.org/x/tools/gopls@latest
:CocInstall coc-go
```

- golangci-lint

```sh
brew tap golangci/tap
brew install golangci/tap/golangci-lint
golangci-lint run
```

- errcheck

```sh
go install github.com/kisielk/errcheck@latest
errcheck ./...
```

- go vet (built-in)

```sh
go vet
```

- staticcheck

```sh
brew install staticcheck
staticcheck [flags] [package]
```
