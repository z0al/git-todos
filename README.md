<div align="center">
  <img src="./docs/logo.png" width="200" alt="Git Todos Logo"/>
  <p>A Git based Todos App for Developers</p>

  <hr>

  <a href="https://github.com/ahmed-taj/git-todos/releases">
    <img src="https://img.shields.io/github/downloads/ahmed-taj/git-todos/latest/total.svg?style=flat-square" alt="downloads count"/>
  </a>

  <a href="https://github.com/ahmed-taj/git-todos/releases">
    <img src="https://img.shields.io/github/release/ahmed-taj/git-todos.svg?style=flat-square" alt="current release number"/>
  </a>

  <a href="https://saythanks.io/to/ahmed-taj">
    <img src="https://img.shields.io/badge/Say%20Thanks-💖-CA1F7B.svg?style=flat-square" alt="say thanks to the author"/>
  </a>

  <a href="http://twitter.com/home?status=Check out this cool project by @ah_tajelsir https://git.io/todos">
    <img src="https://img.shields.io/badge/Share-with%20friends-blue.svg?logo=twitter&style=flat-square" alt="share with friends on twitter"/>
  </a>
</div>

Git-todos is Command Line Interface (CLI) that helps to manage local To-Dos
with ease. It aims to help you get things done rather than organizing them.

You can use Git-todos to add/remove To-Do items locally, import issues from
remote repositories and to automate the generation of Git commits. Git-todos
stores To-Dos list locally, and per-repository.

## Features

* :book: Uses plain text file as storage
* :raised_hands: Developers friendly
* :rocket: Helps to get things done
* :octocat: Easily import issues from GitHub
* :v: Adds a little bit of encouragements

Interrested to learn more? Read the original [blog post](https://medium.com/ahmed-t-ali/local-to-dos-for-developers-d871682a069).

## Installation

[![Latest release](https://img.shields.io/github/release/ahmed-taj/git-todos.svg?style=flat-square)](https://github.com/ahmed-taj/git-todos/releases)

You can grab the latest binary from the [releases](https://github.com/ahmed-taj/git-todos/releases) page.

### Install on Linux

On Linux you can install or update to the latest released version with [snap](https://snapcraft.io/):

```bash
$ snap install git-todos
$ snap refresh git-todos
```

> **Note:** Because snaps with [strict](https://docs.snapcraft.io/reference/confinement) confinement can't access the global `~/.gitconfig` file, you may need to configure `user.name` and `user.email` per repository in order for git-todos to perform commits.
>
> Run these commands to resolve the issue:
>
> ```bash
> $ git config --local user.email "you@example.com"
> $ git config --local user.name "Your Name"
> ```

### Install on macOS

On macOS you can install or upgrade to the latest released version with Homebrew:

```bash
$ brew tap ahmed-taj/git-todos
$ brew install git-todos
$ brew upgrade git-todos
```

Do you prefer another way to install? [Let me know](https://github.com/ahmed-taj/git-todos/issues/new)

## Usage

The project provides a simple Command Line Interface (CLI). To use the tool simply
run `git-todos`. If the tool binary is available in your system `PATH` you may
also run it without the `-` e.g: `git todos <command>`, thanks to Git.

```bash
$ git-todos help
A Git based Todos App for Developers ⚡

Usage:
  git-todos [command]

Available Commands:
  add         Add a new Todo
  finish      Finish a Todo and commit staged changes
  help        Help about any command
  import      Import an issue from remote Provider (ie. GitHub) as Todo
  list        List available Todos
  mark        Mark a single Todo
  remove      Remove existing Todo
  show        Show Todo details
  version     Print the version number

Flags:
  -h, --help   help for git-todos

Use "git-todos [command] --help" for more information about a command.
```

## Like it?

Give it a star(:star:) :point_up_2: and share it with your friends. Enjoy!

## License

The project is released under the Apache 2.0 license. See [LICENSE](./LICENSE)
