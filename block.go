package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"
)

type Block struct {
	PreviousHash  []byte
	Hash          []byte
	Data          []byte
	UnixTimestamp int64
	Nonce         int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.UnixTimestamp, 10))
	headers := bytes.Join([][]byte{b.PreviousHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, PreviousHash []byte) *Block {
	block := &Block{PreviousHash, []byte{}, []byte(data), time.Now().Unix(), 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		return nil
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		return nil
	}

	return &block
}
