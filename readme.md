# 🐯 Lynnx

> **Lynnx is a powerful, accessible and lightweight URL shortener built with HTML/CSS, and Golang that's backed by a Redis database.**

* Simple, but powerful
* Lightweight and fast
* Designed for accessibility

## Setup

Setting up the Lynnx server for local development is easy, it just requires [Git](http://git-scm.org/), [Golang](http://golang.org/), and [Mage](https://magefile.org/).

```command
# Clone repository:
$ git clone git@github.com:lukewhrit/lynnx
$ cd lynnx

# Install packages and build to a binary:
$ mage build
```

## API

### `GET /v1/:short`

* **Response:**

  ```json
  {
  	"long": "https://github.com/",
    "success": true
  }
  ```

* **URL Parameters:**

  ```json
  {
  	"short": ""
  }
  ```

### `POST /v1/`

* **Request:**

  ```json
  {
  	"long": "https://github.com/"
  }
  ```

* **Response:**

  ```json
  {
  	"short": "On2f2e3v",
  	"success": true
  }
  ```

# Contributors

* [Luke Whrit <me@lukewhrit.xyz>](https://github.com/lukewhrit) — Creator and Maintainer.

## License

Lynnx is available under the MIT license. A copy of this license can be found in the [`license`](license) file.
