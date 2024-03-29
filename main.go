// Copyright 2021 ClavinJune/serve
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
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/ClavinJune/serve/internal"
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
		internal.LogFatal(err)
	}

	if _, err := os.Stat(root); err != nil {
		internal.LogFatal(err)
	}

	return root
}

func mustGetListener(port int) net.Listener {
	addr := fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		internal.LogFatal(err)
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
	if *versionFlag {
		fmt.Printf("serve %s-%s %s/%s BuildBy=%s BuildDate=%s",
			version, commit,
			runtime.GOOS, runtime.GOARCH,
			builtBy, date)
		return
	}

	internal.LogSetQuite(*quietFlag)

	rootDir := mustGetRootDirectory(*rootFlag)
	listener := mustGetListener(*portFlag)
	server := mustGetServer(rootDir, *spaFlag)

	go func() {
		internal.LogF("listen and serve %s/ at http://0.0.0.0:%d",
			rootDir,
			*portFlag,
		)
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			internal.LogFatal(err)
		} else {
			internal.Log("shutdown gracefully")
		}
	}()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stopSignal

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		internal.LogFatal(err)
	}
}
