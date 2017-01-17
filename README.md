# go-shm

[![GoDoc](https://godoc.org/github.com/tmthrgd/go-shm?status.svg)](https://godoc.org/github.com/tmthrgd/go-shm)

go-shm provides functions to open and unlink shared memory without relying on cgo (on Linux).
It provides an implementation of [`shm_open`](https://linux.die.net/man/3/shm_open) and
[`shm_unlink`](https://linux.die.net/man/3/shm_unlink) from `sys/mman.h`.

## Download

```
go get github.com/tmthrgd/go-shm
```

## License

Unless otherwise noted, the go-shm source files are distributed under the Modified BSD License
found in the LICENSE file.
