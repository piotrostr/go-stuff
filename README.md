# go stuff

me studying the art of gopher

## great resources

- [Effective Go](https://go.dev/doc/effective_go)

  a nice read, outdated a notch but was quite informative

- [Uber Go Guide](https://github.com/uber-go/guide/blob/master/style.md)

  excellent, a lot of good principles and common pitfalls

## keeping the code clean

- coc-go for nvim, based on gopls

```bash
go install golang.org/x/tools/gopls@latest
:CocInstall coc-go
```

- errcheck

```bash
go install github.com/kisielk/errcheck@latest
errcheck ./...
```

- go vet (built-in)

```bash
go vet
```

- staticcheck

```bash
brew install staticcheck
staticcheck [flags] [package]
```
