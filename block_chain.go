package main

import "fmt"

type BlockChain struct {
	Blocks []Block `json:blocks`
}

func (bc *BlockChain) add_block(data string) Block {
	b := new_block(data)
	if len(bc.Blocks) == 0 {
		b = new_gen_block(data)
		bc.Blocks = append(bc.Blocks, b)
		LENGTH++
		return b
	}
	b.PrevHash = bc.Blocks[LENGTH-1].Hash
	bc.Blocks = append(bc.Blocks, b)
	LENGTH++
	return b
}

func (bc *BlockChain) mine(id int) (bool, Block, Block) {
	if TOP >= LENGTH {
		fmt.Println("All Blocks mined, wait for incoming blocks")
		return false, Block{}, Block{}
	}
	if TOP != id {
		return false, Block{}, Block{}
	}
	b := bc.Blocks[TOP]
	pow := new_pow(b)
	nonce, hash := pow.search()
	b.nonce = nonce
	b.Hash = hash
	b.Mined = true
	if TOP+1 < len(bc.Blocks) {
		bc.Blocks[TOP+1].PrevHash = hash
		return true, b, bc.Blocks[TOP+1]
	}
	return true, b, Block{}
}
