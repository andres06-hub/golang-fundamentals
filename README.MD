<p align="center">
  <img width="200" height="200" src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png"/>
</p>

These are the basic concepts of golang syntax, variables, functions, among others.
There will also be examples of how to create a basic server.

## Install golang witn asdf

You can use [asdf](https://asdf-vm.com/) managing multiple runtime versions with a single CLI tool

> - Intall asdf [read more here](https://asdf-vm.com/guide/getting-started.html)

# Install

```bash
asdf plugin-add golang https://github.com/kennyp/asdf-golang.git
```

> Then download the version of goland that you require

```bash
asdf install golang <version>
#or
asdf install golang lts
```

> Then list the versions and set it global [read more here](https://asdf-vm.com/guide/getting-started.html#_6-set-a-version)

- List the versions

```bash
asdf list
```

- Set as global

```bash
asdf global golang <version>
```
