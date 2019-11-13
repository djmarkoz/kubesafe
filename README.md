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
