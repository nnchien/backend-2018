<p align="center">
    <a href="https://caddyserver.com"><img src="https://user-images.githubusercontent.com/1128849/36338535-05fb646a-136f-11e8-987b-e6901e717d5a.png" alt="Caddy" width="450"></a>
</p>
<h3 align="center">Every Site on HTTPS <!-- Serve Confidently --></h3>
<p align="center">Caddy is a general-purpose HTTP/2 web server that serves HTTPS by default.</p>

---

Caddy is a **production-ready** open-source web server that is fast, easy to use, and makes you more productive.

Available for Windows, Mac, Linux, BSD, Solaris, and [Android](https://github.com/mholt/caddy/wiki/Running-Caddy-on-Android).

## Menu

- [Features](#features)
- [Build](#build)
- [Quick Start](#quick-start)

## Features

- **Easy configuration** with the Caddyfile
- **Automatic HTTPS** on by default (via [Let's Encrypt](https://letsencrypt.org))
- **HTTP/2** by default
- **Virtual hosting** so multiple sites just work
- Experimental **QUIC support** for cutting-edge transmissions
- TLS session ticket **key rotation** for more secure connections
- **Extensible with plugins** because a convenient web server is a helpful one
- **Runs anywhere** with **no external dependencies** (not even libc)

[See a more complete list of features built into Caddy.](https://caddyserver.com/features) On top of all those, Caddy does even more with plugins: choose which plugins you want at [download](https://caddyserver.com/download).

Altogether, Caddy can do things other web servers simply cannot do. Its features and plugins save you time and mistakes, and will cheer you up. Your Caddy instance takes care of the details for you!

## Build

To build from source you need **[Git](https://git-scm.com/downloads)** and **[Go](https://golang.org/doc/install)** (1.9 or newer). Follow these instruction for fast building:

- Get the source with `go get github.com/mholt/caddy/caddy` and then run `go get github.com/caddyserver/builds`
- Now `cd $GOPATH/src/github.com/mholt/caddy/caddy` and run `go run build.go`

Then make sure the `caddy` binary is in your PATH.

To build for other platforms, use build.go with the `--goos` and `--goarch` flags.


## Quick Start

Expose a `https` static web page, open `websocket`, manage `logs`

1. create an `index.html` file
2. create `Caddyfile` for configuration
3. run `caddy`

There will be an investment on Caddy with more complex servers
