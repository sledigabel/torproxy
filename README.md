# torproxy
A http proxy to TOR using go

## What does it do?

It accepts HTTP connections and runs them through the TOR socket.

## Why does this exist?

It's to access TOR from docker or the cli easily.
Plus it's a fun little project.

## What does it use?

- [Bine](https://github.com/cretz/bine): the go library for TOR
- [Goproxy](https://github.com/elazarl/goproxy): a library providing some easy proxying

