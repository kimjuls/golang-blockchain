package main

type Block struct {
	PreviousHash  []byte
	Hash          []byte
	Data          []byte
	UnixTimestamp int64
}
