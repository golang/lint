[![wercker status](https://app.wercker.com/status/c390b2a2d61a54339208ab6534382534/m "wercker status")](https://app.wercker.com/project/bykey/c390b2a2d61a54339208ab6534382534)

Original: https://github.com/golang/lint

Golintx is a linter for Go source code.

# Differences From Original golint

* Support per-directory config files
    * if a config file on lint-target directory does not exists, search files in ancestor diretories recursively.
* exit(1) if any problem exists

# Installation

Golintx requires Go 1.5 or later.

    go get -u github.com/haruyama/golintx/golintx
