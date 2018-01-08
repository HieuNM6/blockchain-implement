package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func main() {
	bc, err := NewBlockchain()
	if err != nil {
		log.Fatalf("Failed to create or open blochain by error: %v \n", err)
	}

	defer func() {
		if err := bc.db.Close(); err != nil {
			log.Println(err)
		}
	}()

	cli := CLI{bc}
	cli.Run()
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
