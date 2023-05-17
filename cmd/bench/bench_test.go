package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/kellymclaughlin/rust-serialization-benchmarks/pkg/types"
	"github.com/ugorji/go/codec"
)

func BenchmarkMarshalJSON(b *testing.B) {
	msg := types.NewIngestData()

	buf, err := json.Marshal(&msg)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	buffer := bytes.NewBuffer(make([]byte, 1<<20))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buffer.Reset()
		encoder := json.NewEncoder(buffer)
		err := encoder.Encode(&msg)
		if err != nil {
			b.Fatalf("Marshal: %v", err)
		}
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	msg := types.NewIngestData()
	buf, err := json.Marshal(&msg)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	var tmp types.IngestData
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(buf, &tmp)
		if err != nil {
			b.Fatalf("Unmarshal: %v", err)
		}
	}
}

func BenchmarkMarshalMsgPack(b *testing.B) {
	h := new(codec.MsgpackHandle)
	msg := types.NewIngestData()
	var buf []byte
	enc := codec.NewEncoderBytes(&buf, h)
	err := enc.Encode(msg)
	if err != nil {
		b.Fatalf("Encode: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		err := enc.Encode(msg)
		if err != nil {
			b.Fatalf("Encode: %v", err)
		}
	}
}

func BenchmarkUnmarshalMsgPack(b *testing.B) {
	h := new(codec.MsgpackHandle)
	msg := types.NewIngestData()
	var buf []byte
	enc := codec.NewEncoderBytes(&buf, h)
	err := enc.Encode(msg)
	if err != nil {
		b.Fatalf("Encode: %v", err)
	}

	b.SetBytes(int64(len(buf)))

	b.ResetTimer()

	var tmp types.IngestData
	for i := 0; i < b.N; i++ {
		err := codec.NewDecoderBytes(buf, h).Decode(&tmp)
		if err != nil {
			b.Fatalf("Decode: %v", err)
		}

	}
}
