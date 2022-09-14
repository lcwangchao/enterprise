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
	"context"

	"github.com/pingcap/tidb/extensions"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/mysql"
	"github.com/pingcap/tidb/util/chunk"
)

func handleCommand(stmt ast.ExtensionCmdNode) (extensions.ExtensionCmdHandler, error) {
	switch stmt.(type) {
	case *ast.AuditCmdStmt:
		return &auditCmdHandler{}, nil
	default:
		return nil, nil
	}
}

type auditCmdHandler struct {
}

func (h *auditCmdHandler) OutputColumnsNum() int {
	return 2
}

func (h *auditCmdHandler) BuildOutputSchema(addColumn func(tableName string, name string, tp byte, size int)) {
	addColumn("", "COL1", mysql.TypeLonglong, 4)
	addColumn("", "COL2", mysql.TypeVarchar, 64)
}

func (h *auditCmdHandler) ExecuteCmd(_ context.Context, chk *chunk.Chunk) error {
	chk.AppendInt64(0, 1)
	chk.AppendString(1, "row1")

	chk.AppendInt64(0, 2)
	chk.AppendString(1, "row2")
	return nil
}
