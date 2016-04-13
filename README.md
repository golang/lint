[![wercker status](https://app.wercker.com/status/c390b2a2d61a54339208ab6534382534/m "wercker status")](https://app.wercker.com/project/bykey/c390b2a2d61a54339208ab6534382534)

Original: https://github.com/golang/lint

Golintx is a linter for Go source code.

# Differences From Original golint

* Support per-directory config files
    * if a config file on lint-target directory does not exists, search files in ancestor diretories recursively.
* exit(1) if any problem exists
* Support multi diretories on command line
    * ex: golintx $(glide novendor)


# Installation

Golintx requires Go 1.6 or later.

    go get -u github.com/haruyama/golintx/golintx

# Config File (.golintx.hcl)

```
exclude {
        // array of categories which golintx do not output
        categories = ["comments"]
}
// array of initialisms
initialisms = [
        "API",
        "ASCII",
        "CPU",
        "CSS",
        "DNS",
        "EOF",
        "GUID",
        "HTML",
        "HTTP",
        "HTTPS",
        // "ID",
        "IP",
        "JSON",
        "LHS",
        "QPS",
        "RAM",
        "RHS",
        "RPC",
        "SLA",
        "SMTP",
        "SQL",
        "SSH",
        "TCP",
        "TLS",
        "TTL",
        "UDP",
        "UI",
        "UID",
        "UUID",
        // "URI",
        // "URL",
        "UTF8",
        "VM",
        "XML",
        "XSRF",
        "XSS",
]
```
