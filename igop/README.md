### build igop for Go version >= 1.23

- go.mod use `golang.org/x/tools v0.29.0` for support range func

- build use `-ldflags="-checklinkname=0"` for support linkname

### Install igop for Go version >= 1.23
`go install -ldflags="-checklinkname=0" github.com/goplus/igop/igop@v0.23.9`


### Local development environment

go.work
```
go 1.23

use (
	.
	./..
)
```


