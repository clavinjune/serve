# Serve

[![Go Version](https://img.shields.io/static/v1?style=for-the-badge&label=Go+Version&message=1.17.6&color=blue&logo=go)](https://github.com/golang/go/releases/tag/go1.17.6)
[![Reference](https://img.shields.io/badge/reference-007d9c?style=for-the-badge&logo=go&logoColor=white&&labelColor=5c5c5c)](https://pkg.go.dev/github.com/ClavinJune/serve)
[![Latest Release](https://img.shields.io/github/tag/ClavinJune/serve.svg?style=for-the-badge&logo=github)](https://github.com/ClavinJune/serve/releases/latest)
[![Registry](https://img.shields.io/static/v1?style=for-the-badge&label=Registry&message=ghcr.io&color=red&logo=linux-containers)](https://ghcr.io/clavinjune/serve)
[![License](https://img.shields.io/github/license/ClavinJune/serve?style=for-the-badge)](https://github.com/ClavinJune/serve/blob/main/LICENSE)

A Lightweight HTTP Server Built on Top of Go

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
$ chmod +x serve_0.3.8_linux_amd64
$ ./serve_0.3.8_linux_amd64 -v
serve v0.3.8
```

### Via Go Install

```bash
$ go install github.com/ClavinJune/serve@latest
go: downloading github.com/ClavinJune/serve v0.3.8
$ serve -v
serve v0.3.8
```

### Via Go Run

```bash
$ go run github.com/ClavinJune/serve@latest -v
serve v0.3.8
```

### Via Docker

```bash
$ docker run -it --rm ghcr.io/clavinjune/serve:latest -v
serve v0.3.8
```

### Via Podman

```bash
$ podman run -it --rm ghcr.io/clavinjune/serve:latest -v
serve v0.3.8
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
  -v    print current version
```

## Example

```bash
$ podman run -it --rm \
> -v /tmp/foo:/app/src \
> -p 1313:1313 \
> ghcr.io/clavinjune/serve:latest
2021/11/04 06:32:19 listen and serve /app/src/ at http://0.0.0.0:1313
2021/11/04 06:32:23 / 46.107µs
2021/11/04 06:32:23 /favicon.ico 78.868µs
2021/11/04 06:32:40 /index.html 6.642µs
2021/11/04 06:32:40 / 31.058µs
2021/11/04 06:32:44 /foo/ 60.574µs
```
