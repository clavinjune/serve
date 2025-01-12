// Copyright 2025 clavinjune/serve
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main is the main package where serve is implemented
// Usage of serve:
//
//	-p int
//	  	port on which the server will listen (default 1313)
//	-q	run server quietly
//	-r string
//	  	root document which the server will serve (default ".")
//	-s	serve single page application
//	-v	print current version
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/clavinjune/serve/internal"
)

var (
	defaultReadTimeout = time.Minute
	portFlag           *int
	quietFlag          *bool
	rootFlag           *string
	spaFlag            *bool
	versionFlag        *bool

	version = "dev"
	builtBy = "dev"
	commit  = "n/a"
	date    = "0001-01-01 00:00:00 +0000 UTC"
)

func init() {
	portFlag = flag.Int("p", 1313, "port on which the server will listen")
	quietFlag = flag.Bool("q", false, "run server quietly")
	rootFlag = flag.String("r", ".", "root document which the server will serve")
	spaFlag = flag.Bool("s", false, "serve single page application")
	versionFlag = flag.Bool("v", false, "print current version")
	flag.Parse()
}

func mustGetRootDirectory(dir string) string {
	root, err := filepath.Abs(strings.TrimRight(dir, "/"))
	if err != nil {
		slog.LogAttrs(context.TODO(), slog.LevelError, err.Error())
		os.Exit(1)
	}

	if _, err := os.Stat(root); err != nil {
		slog.LogAttrs(context.TODO(), slog.LevelError, err.Error())
		os.Exit(1)
	}

	return root
}

func mustGetListener(port int) net.Listener {
	addr := fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		slog.LogAttrs(context.TODO(), slog.LevelError, err.Error())
		os.Exit(1)
	}

	return l
}

func mustGetServer(rootDir string, isSpa bool) *http.Server {
	return &http.Server{
		Handler:     internal.Handle(rootDir, isSpa),
		ReadTimeout: defaultReadTimeout,
	}
}

func main() {
	ctx := context.TODO()

	if *versionFlag {
		fmt.Printf("serve %s-%s %s/%s BuildBy=%s BuildDate=%s",
			version, commit,
			runtime.GOOS, runtime.GOARCH,
			builtBy, date)
		return
	}

	level := slog.LevelInfo

	if *quietFlag {
		level = slog.LevelError
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	})))

	rootDir := mustGetRootDirectory(*rootFlag)
	listener := mustGetListener(*portFlag)
	server := mustGetServer(rootDir, *spaFlag)

	go func() {
		slog.LogAttrs(ctx, slog.LevelInfo,
			"listen and serve",
			slog.String("root", rootDir),
			slog.Int("port", *portFlag),
		)
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			slog.LogAttrs(ctx, slog.LevelError, err.Error())
			os.Exit(1)
		}
		slog.LogAttrs(ctx, slog.LevelInfo, "shutdown gracefully")
	}()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stopSignal

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.LogAttrs(ctx, slog.LevelError, err.Error())
		os.Exit(1)
	}
}
