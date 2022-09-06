// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ignore
// +build ignore

package main

import (
	"os"

	"github.com/pingcap/errors"
	"github.com/pingcap/tidb/parser/terror"
)

const (
	bootGoFile        = "../enterprise/boot.go"
	bootGoFileContent = `
package enterprise

import _ "github.com/pingcap/tidb/test-enterprise"
`
)

func main() {
	doClear := false
	if len(os.Args) > 1 {
		action := os.Args[1]
		switch action {
		case "genfile":
			break
		case "clear":
			doClear = true
		default:
			terror.MustNil(errors.New("Invalid action: " + action))
		}
	}

	if doClear {
		err := os.Remove(bootGoFile)
		if !os.IsNotExist(err) {
			terror.MustNil(err)
		}
	} else {
		terror.MustNil(os.WriteFile(bootGoFile, []byte(bootGoFileContent), 0644))
	}
}
