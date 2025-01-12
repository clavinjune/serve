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

package internal

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Handle returns http.HandlerFunc depends on the given parameters
func Handle(dir string, isSpa bool) http.HandlerFunc {
	handler := handleNonSPA(dir)

	if isSpa {
		handler = handleSPA(dir)
	}

	return middleware(handler)
}

func handleSPA(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := filepath.Join(dir, r.URL.Path)
		_, err := os.Open(filepath.Clean(p))

		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}

		http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
	}
}

func handleNonSPA(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			p := filepath.Join(dir, r.URL.Path, "index.html")
			_, err := os.Open(filepath.Clean(p))

			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
		}

		http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
	}
}

func middleware(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		n(w, r)
		slog.LogAttrs(r.Context(), slog.LevelInfo, "",
			slog.String("path", r.URL.Path),
			slog.Duration("since", time.Since(start)),
		)
	}
}
