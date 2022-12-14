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

package enterprise

import (
	"github.com/pingcap/tidb/extensions"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/terror"
	"github.com/pingcap/tidb/sessionctx/variable"
)

const ExtensionName = "demo"

const CustomPriv = "CUSTOM_PRIV"
const CustomVar = "tidb_custom_sys_var"

func Register() {
	terror.MustNil(extensions.Register(CreateExtension))
}

func CreateExtension() (*extensions.ExtensionManifest, error) {
	listener := &connListener{}
	return extensions.NewExtension(
		ExtensionName,
		extensions.WithNewDynamicPrivileges([]string{
			CustomPriv,
		}),
		extensions.WithNewSysVariables([]*variable.SysVar{
			{
				Name:  CustomVar,
				Value: "ON",
				Scope: variable.ScopeGlobal | variable.ScopeSession,
				Type:  variable.TypeBool,
			},
		}),
		extensions.WithHandleCommand(func(n ast.ExtensionCmdNode) (extensions.ExtensionCmdHandler, error) {
			switch stmt := n.(type) {
			case *ast.AuditCmdStmt:
				return NewAuditCmdHandler(stmt)
			default:
				return nil, nil
			}
		}),
		extensions.WithHandleConnect(listener.HandleConn),
	), nil
}
