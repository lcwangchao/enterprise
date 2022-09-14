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
	"go.uber.org/zap"
)

type connHandler struct{}

func (h *connHandler) CreateConnEventListener() extensions.ConnEventListener {
	return h
}

func (h *connHandler) OnConnEvent(event *extensions.ConnEvent) {
	var name string
	switch event.Tp {
	case extensions.ConnEstablished:
		name = "ConnEstablished"
	case extensions.ConnAuthenticated:
		name = "ConnAuthenticated"
	case extensions.ConnRejected:
		name = "ConnRejected"
	case extensions.ConnReset:
		name = "ConnReset"
	case extensions.ConnDisconnect:
		name = "ConnDisconnect"
	default:
		name = fmt.Sprintf("%v", event.Tp)
	}

	log.Info(name,
		zap.Uint64("ConnectionID", event.ConnectionID),
		zap.String("ConnectionType", event.ConnectionType),
		zap.String("ClientIP", event.ClientIP),
		zap.String("ClientPort", event.ClientPort),
		zap.Int("ServerID", event.ServerID),
		zap.Int("ServerPort", event.ServerPort),
		zap.String("User", event.User),
		zap.String("ClientVersion", event.ClientVersion),
		zap.String("SSLVersion", event.SSLVersion),
		zap.String("DB", event.DB),
	)
}

func (h *connHandler) CreateStmtEventListener() extensions.StmtEventListener {
	return h
}

func (h *connHandler) OnStmtEvent(event *extensions.StmtEvent) {
	var name string
	switch event.Tp {
	case extensions.StmtParserError:
		name = "StmtParserError"
	case extensions.StmtStart:
		name = "StmtStart"
	case extensions.StmtEnd:
		if event.StmtContext.Error() != nil {
			name = "StmtError"
		} else {
			name = "StmtSuccess"
		}
	default:
		name = fmt.Sprintf("%v", event.Tp)
	}

	normalized, _ := event.StmtDigest()
	log.Info(name,
		zap.Uint64("ConnectionID", event.Conn.ConnectionID),
		zap.String("OriginalSQL", event.OriginalSQL()),
		zap.Strings("ExecuteArgs", event.StmtArguments()),
		zap.String("NormalizedSQL", normalized),
	)
}
