# injest -- wraps your input in jokes

[![golangci-lint](https://github.com/conacademy/injest/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/conacademy/injest/actions/workflows/golangci-lint.yml)
[![CodeQL](https://github.com/conacademy/injest/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/conacademy/injest/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/conacademy/injest)](https://goreportcard.com/report/github.com/conacademy/injest)

`injest` wraps input into a joke and emits it to `stdout`.

Origin Story.... I do amateur nerd-code freestyle and used a double-entendre of "in jest" and "ingest" in a rhyme (no memory of the whole phrasing, that's how it goes).   I thought that was really funny, and [ConAcademy](https://github.com/conacademy) needed some real code, so here we are.

----

## Installing

Binaries for multiple platforms are [released on GitHub](https://github.com/conacademy/injest/releases) through [GitHub Actions](https://github.com/conacademy/injest/actions).

You can also install for various platforms with [Homebrew](https://brew.sh) from [`conacademy/homebrew-tap`](https://github.com/conacademy/homebrew-tap):

```
brew tap conacademy/homebrew-tap
brew install conacademy/tap/injest
```

----

## Example Usage

```
> injest --list
yomamma
knock

> dig +short pi.neomantra.com | injest --stdin -j knock
Knock knock
Who's there?
Your data.
Your data, who?
Your data is not a joke:
3.141.59.27

> injest -j yomamma I seen her in the back of Taco Bell with handcuffs
Yo momma so fat that I seen her in the back of Taco Bell with handcuffs
```

----

## Building

Building is performed with [task](https://taskfile.dev/):

```
$ task
task: [build] go build -o injest cmd/injest/main.go
```

----

## Contribution and Conduct

As with all ConAcademy projects, pull requests are welcome.  Or fork it.  You do you.

Either way, obey our [Code of Conduct](./CODE_OF_CONDUCT.md).  Be shady, but don't be a jerk.

----

## Credits and License

Copyright (c) 2023 Neomantra BV.  Authored by Evan Wies for [ConAcademy](https://github.com/conacademy).

Released under the [MIT License](https://en.wikipedia.org/wiki/MIT_License), see [LICENSE.txt](./LICENSE.txt).
