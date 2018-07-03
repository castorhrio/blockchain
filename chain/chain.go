package chain

import "../block"

type Blockchain struct {
	Chains []*block.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Chains[len(bc.Chains)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.Chains = append(bc.Chains, newBlock)
}

func NewGenesisBlock() *block.Block {
	return block.NewBlock("Genesis Block", []byte{})
}

func NewChain() *Blockchain {
	return &Blockchain{[]*block.Block{NewGenesisBlock()}}
}
