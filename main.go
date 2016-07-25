// Copyright © 2016 Robert Coleman <github@robert.net.nz>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
)

func main() {
	for _, key := range os.Args[1:] {
		s := strings.Split(key, "_")
		prefix := s[0]

		value := os.Getenv(key)

		u, err := url.Parse(value)
		if err != nil {
			os.Exit(1)
		}

		if u.Scheme == "" {
			os.Exit(1)
		}

		if u.Scheme != "" {
			export(prefix, "SCHEME", u.Scheme)
		}

		if u.User.Username() != "" {
			export(prefix, "USERNAME", u.User.Username())
		}

		password, _ := u.User.Password()
		if password != "" {
			export(prefix, "PASSWORD", password)
		}

		host, port, _ := net.SplitHostPort(u.Host)
		if host != "" {
			export(prefix, "HOST", host)
		}
		if port != "" {
			export(prefix, "PORT", port)
		}

		if u.Path != "" {
			export(prefix, "PATH", u.Path)
		}

		if u.Fragment != "" {
			export(prefix, "FRAGMENT", u.Fragment)
		}

		if u.RawQuery != "" {
			export(prefix, "QUERYSTRING", u.RawQuery)
		}
	}
}

func export(prefix, key, value string) {
	fmt.Printf("export %s_%s=%s\n", prefix, key, value)
}
