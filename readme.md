# Lynnx

> 🐆 **Lynnx is a powerful, accessible and lightweight URL shortener built with SvelteJS and Golang and backed by a Redis database.**
>\
>\
> **[Contributing Guide](contributing.md) | [Website](lynnx.me) | [Frontend Repo](https://github.com/lukewhrit/lynnx-frontend)**

* Designed for accessability
* Lightweight and Fast
* Available all over the world

## Setup

Setting up the Lynnx backend for local development is easy, it just requires [Git](http://git-scm.org/), [Golang](http://golang.org/) and [Mage](https://magefile.org/).

This project makes use of Golang Modules for dependency management, make sure they're enabled.

```sh
# Clone repository:
$ git clone git@github.com:lukewhrit/lynnx
$ cd lynnx

# Install packages and build to a binary:
$ mage build
```

## Contributors

* [Luke Whrit <me@lukewhrit.xyz>](https://github.com/lukewhrit) — Creator and Maintainer.

## License

Lynnx is available under the MIT license. A copy of this license can be found in the [`license`](license) file.
