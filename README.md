# Introduction

`aliasnimby` checks a package for [alias declarations](https://tip.golang.org/ref/spec#Alias_declarations), and if it
finds one, returns a non zero exit status - just like everyone else does when aliases are discussed. It's not an
overreaction at all.

# Usage

```
go get -u github.com/bradleyfalzon/aliasnimby/...
aliasnimby [package]
```

# Example

```
$ cat testdata/testdata.go
package testdata

type T1 struct{}

type T3 = T1
$ aliasnimby ./testdata/
/home/bradleyf/go/src/github.com/bradleyfalzon/aliasnimby/testdata/main.go:5:6: T3 is an alias
Argh fark, it's in my backyard!
```

