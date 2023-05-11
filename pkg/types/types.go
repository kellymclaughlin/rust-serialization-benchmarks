package types

import "bytes"

//go:generate codecgen -o types.generated.go types.go

type LogMsg struct {
	SvcID    string `codec:"svc_id" json:"svc_id"`
	Endpoint string `codec:"endpoint" json:"endpoint"`
	Msg      []byte `codec:"msg" json:"msg"`
}

type IngestData struct {
	Type      string `codec:"type" json:"type"`
	Source    string `codec:"src" json:"src"`
	Timestamp string `code:"timestamp" json:"timestamp"`
	Msg       LogMsg `codec:"msg" json:"msg"`
}

func NewIngestData() *IngestData {
	logMsg := LogMsg{
		SvcID:    "deadbeef",
		Endpoint: "sometestendpoint",
		Msg:      bytes.Repeat([]byte("a"), 131072),
	}

	return &IngestData{
		Type:      "log",
		Source:    "xqd",
		Timestamp: "Sun Apr 23 01:48:34 PM MDT 2023",
		Msg:       logMsg,
	}
}
