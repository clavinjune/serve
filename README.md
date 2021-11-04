# Serve

[![Go Version](https://img.shields.io/static/v1?style=for-the-badge&label=Go+Version&message=1.17.2&color=blue&logo=go)](https://github.com/golang/go/releases/tag/go1.17.2)
[![Latest Release](https://img.shields.io/github/tag/ClavinJune/serve.svg?style=for-the-badge&logo=github)](https://github.com/ClavinJune/serve/releases/latest)
[![License](https://img.shields.io/github/license/ClavinJune/serve?style=for-the-badge)](https://github.com/ClavinJune/serve/blob/main/LICENSE)

A Simple HTTP Server Built on Top of Go

## Run

### Via Download

- Download from [here](https://github.com/ClavinJune/serve/releases)

| Operating System | Name Format |
| --- | --- |
| MacOS | serve_{version}\_darwin\_{arch} |
| Linux | serve_{version}\_linux\_{arch} |
| Windows | serve_{version}\_windows\_{arch}.exe |

- For MacOS and Linux, make the file executable

```bash
# chmod +x serve_{version}_{os}_{arch}
$ chmod +x serve_0.0.1_linux_amd64
$ ./serve_0.0.1_linux_amd64 -h
```

### Via Golang Install

```bash
$ go install github.com/ClavinJune/serve@latest
$ serve -h
```

### Via Go Run

```bash
$ go run github.com/ClavinJune/serve@latest -h
```

### Via Docker

```bash
$ docker run -it --rm ghcr.io/clavinjune/serve:latest -h
```

## Usage

```bash
$ serve -h
Usage of serve:
  -p int
        port on which the server will listen (default 1313)
  -q    run server quietly
  -r string
        root document which the server will serve (default ".")
  -s    serve single page application
```