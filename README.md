# kubesafe
Verify expected cluster before running Kubernetes commands to prevent accidental changes to the wrong cluster.

## Install

### MacOS

#### Homebrew (recommended)

```bash
brew tap djmarkoz/djmarkoz
brew install kubesafe
```

### MacOS / Linux

#### Go get

```bash
go get -u github.com/djmarkoz/kubesafe
```

## Usage

```bash
kubesafe
Verify expected cluster before running Kubernetes commands

Usage:
  kubesafe [command]

Available Commands:
  explain     Explain which directory expects which cluster
  get         Returns the currently expected cluster
  help        Help about any command
  set         Create a '.kubesafe-expected-cluster' file in the current directory specifying the currently active cluster
  unset       Remove the '.kubesafe-expected-cluster' file in the current directory
  verify      Verify if the current cluster is the expected cluster and optionally run a command
```

## Aliases

```bash
alias ks='kubesafe'
alias kss='kubesafe set'
alias ksu='kubesafe unset'
alias ksg='kubesafe get'
alias kse='kubesafe explain'
alias ksv='kubesafe verify'
alias kubectl='kubesafe verify -- \kubectl'
alias helm='kubesafe verify -- \helm'
```
