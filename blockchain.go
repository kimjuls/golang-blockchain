package main

type Blockchain struct {
	blocks []*Block
}

func NewGenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}

func (blockchain *Blockchain) CreateBlock(data string) {
	previousBlock := blockchain.blocks[len(blockchain.blocks)-1]
	created := NewBlock(data, previousBlock.Hash)
	blockchain.blocks = append(blockchain.blocks, created)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
