# Serve

[![Go Version](https://img.shields.io/static/v1?style=for-the-badge&label=Go+Version&message=1.19.5&color=blue&logo=go)](https://github.com/golang/go/releases/tag/go1.23.4)
[![Reference](https://img.shields.io/badge/reference-007d9c?style=for-the-badge&logo=go&logoColor=white&&labelColor=5c5c5c)](https://pkg.go.dev/github.com/clavinjune/serve)
[![Go Report Card](https://goreportcard.com/badge/github.com/clavinjune/serve?style=for-the-badge)](https://goreportcard.com/report/github.com/clavinjune/serve)
[![Latest Release](https://img.shields.io/github/tag/clavinjune/serve.svg?style=for-the-badge&logo=github)](https://github.com/clavinjune/serve/releases/latest)
[![Registry](https://img.shields.io/static/v1?style=for-the-badge&label=Registry&message=ghcr.io&color=red&logo=linux-containers)](https://ghcr.io/clavinjune/serve)
[![License](https://img.shields.io/github/license/clavinjune/serve?style=for-the-badge)](https://github.com/clavinjune/serve/blob/main/LICENSE)

A Lightweight HTTP Server Built on Top of Go

## Run

### Via Download

- Download from [here](https://github.com/clavinjune/serve/releases)

| Operating System | Name Format |
| --- | --- |
| MacOS | serve_{version}\_darwin\_{arch} |
| Linux | serve_{version}\_linux\_{arch} |
| Windows | serve_{version}\_windows\_{arch}.exe |

- For MacOS and Linux, make the file executable

```bash
# chmod +x serve_{version}_{os}_{arch}
$ wget https://github.com/clavinjune/serve/releases/download/v1.1.0/serve_1.1.0_darwin_arm64
$ chmod +x serve_1.1.0_darwin_arm64
$ ./serve_1.1.0_darwin_arm64 -v
serve 1.1.0-0d66413211647f62769a52854979cb84af398b62 darwin/arm64 BuildBy=goreleaser BuildDate=2022-04-15T22:09:06Z
```

### Via Go Install

```bash
$ go install github.com/clavinjune/serve@latest
go: downloading github.com/clavinjune/serve v1.1.0
$ serve -v
serve dev-n/a darwin/arm64 BuildBy=dev BuildDate=0001-01-01 00:00:00 +0000 UTC
```

### Via Go Run

```bash
$ go run github.com/clavinjune/serve@latest -v
serve dev-n/a darwin/arm64 BuildBy=dev BuildDate=0001-01-01 00:00:00 +0000 UTC
```

### Via Docker

```bash
$ docker run --rm ghcr.io/clavinjune/serve:latest -v
serve 1.1.0-0d66413211647f62769a52854979cb84af398b62 linux/arm64 BuildBy=goreleaser BuildDate=2022-04-15T22:09:06Z
```

### Via Podman

```bash
$ podman run --rm ghcr.io/clavinjune/serve:latest -v
serve 1.1.0-0d66413211647f62769a52854979cb84af398b62 linux/arm64 BuildBy=goreleaser BuildDate=2022-04-15T22:09:06Z
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
$ podman run --rm \
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
