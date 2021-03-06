// Copyright © 2019 Martin Tournoij – This file is part of GoatCounter and
// published under the terms of a slightly modified EUPL v1.2 license, which can
// be found in the LICENSE file or at https://license.goatcounter.com

package main

import (
	"runtime"
	"testing"
)

func TestServe(t *testing.T) {
	// I don't know why, but this doesn't work in Windows; I think it may be
	// related to permission issues for binding to a port or some such?
	if runtime.GOOS == "windows" {
		t.Skip()
	}

	_, dbc, clean := tmpdb(t)
	defer clean()

	run(t, 0, []string{"serve", "-test-hook-do-not-use=1",
		"-debug", "all",
		"-listen", "localhost:31874",
		"-tls", "none",
		"-db", dbc})
}
