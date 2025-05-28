// SPDX-FileCopyrightText: Copyright The Authors
// SPDX-License-Identifier: The Unlicense

package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/alexandear-org/golangci-lint-eol-windows/internal"
)

func main() {
	spew.Dump("a")
	fmt.Println(internal.Exported())
}
