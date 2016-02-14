# rands

[![GoDoc](https://godoc.org/github.com/zbiljic/rands?status.svg)](https://godoc.org/github.com/zbiljic/rands)

Package rands provides functions for generating random strings in Go.

It combines [math/rand][] with [crypto/rand][] for acceptable ratio of performance/security.

## Installation

```bash
go get -u github.com/zbiljic/rands
```

## Example:

```go
import "github.com/zbiljic/rands"

rands.AlphanumericString(64) // returns random string with 64 alpha-numeric characters
rands.AlphabeticString(32)   // returns random string with 32 alphabetic characters
rands.NumericString(16)      // returns random string with 16 numeric characters
```

See the [reference][] for more info.

[math/rand]: https://golang.org/pkg/math/rand
[crypto/rand]: https://golang.org/pkg/crypto/rand
[reference]: http://godoc.org/github.com/zbiljic/rands

---

Copyright © 2016 Nemanja Zbiljić
