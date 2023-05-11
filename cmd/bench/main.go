package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kellymclaughlin/rust-serialization-benchmarks/pkg/types"
	"github.com/ugorji/go/codec"
)

func main() {
	h := new(codec.MsgpackHandle)
	msg := types.NewIngestData()
	var buf []byte
	enc := codec.NewEncoderBytes(&buf, h)
	err := enc.Encode(msg)
	if err != nil {
		log.Fatalf("Encode: %v", err)
	}

	fmt.Printf("buf size: %d\n", len(buf))

	start := time.Now()
	for i := 0; i < 1_000_000; i++ {
		var tmp types.IngestData
		err := codec.NewDecoderBytes(buf, h).Decode(&tmp)
		if err != nil {
			log.Fatalf("Decode: %v", err)
		}

	}
	duration := time.Since(start).Nanoseconds()
	fmt.Printf("Duration: %v\n", duration)
}
