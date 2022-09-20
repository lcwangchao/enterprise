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
	"fmt"

	"github.com/pingcap/log"
	"github.com/pingcap/tidb/extensions"
	"github.com/pingcap/tidb/sessionctx/variable"
	"go.uber.org/zap"
)

type connHandler struct{}

func NewConnHandler() (extensions.ConnHandler, error) {
	return &connHandler{}, nil
}

func (h *connHandler) ConnEventListener() extensions.ConnEventListener {
	return h
}

func (h *connHandler) OnConnEvent(tp extensions.ConnEventTp, conn *variable.ConnectionInfo) {
	var name string
	switch tp {
	case extensions.Connected:
		name = "Connected"
	case extensions.ConnAuthenticated:
		name = "ConnAuthenticated"
	case extensions.ConnRejected:
		name = "ConnRejected"
	case extensions.ConnReset:
		name = "ConnReset"
	case extensions.ConnDisconnect:
		name = "ConnDisconnect"
	default:
		name = fmt.Sprintf("%v", tp)
	}

	log.Info(name,
		zap.Uint64("ConnectionID", conn.ConnectionID),
		zap.String("ConnectionType", conn.ConnectionType),
		zap.String("ClientIP", conn.ClientIP),
		zap.String("ClientPort", conn.ClientPort),
		zap.Int("ServerID", conn.ServerID),
		zap.Int("ServerPort", conn.ServerPort),
		zap.String("User", conn.User),
		zap.String("ClientVersion", conn.ClientVersion),
		zap.String("SSLVersion", conn.SSLVersion),
		zap.String("DB", conn.DB),
	)
}

func (h *connHandler) StmtEventListener() extensions.StmtEventListener {
	return h
}

func (h *connHandler) OnStmtEvent(tp extensions.StmtEventTp, stmt extensions.StmtEventContext) {
	var name string
	switch tp {
	case extensions.StmtParserError:
		name = "StmtParserError"
	case extensions.StmtStart:
		name = "StmtStart"
	case extensions.StmtEnd:
		if stmt.Error() != nil {
			name = "StmtError"
		} else {
			name = "StmtSuccess"
		}
	default:
		name = fmt.Sprintf("%v", tp)
	}

	normalized, _ := stmt.StmtDigest()
	log.Info(name,
		zap.Uint64("ConnectionID", stmt.GetConnectionInfo().ConnectionID),
		zap.String("OriginalSQL", stmt.OriginalSQL()),
		zap.Strings("ExecuteArgs", stmt.StmtArguments()),
		zap.String("NormalizedSQL", normalized),
	)
}
