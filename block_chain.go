package main

import "fmt"

type BlockChain struct {
	Blocks []*Block `json:blocks`
}

func (bc *BlockChain) add_block(data string) {
	b := new_block(data)
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, new_gen_block(data))
		LENGTH++
		return
	}
	b.PrevHash = bc.Blocks[LENGTH-1].Hash
	bc.Blocks = append(bc.Blocks, b)
	LENGTH++
}

func (bc *BlockChain) mine(id int) bool {
	if TOP >= LENGTH {
		fmt.Println("All Blocks mined, wait for incoming blocks")
		return false
	}
	if TOP != id {
		return false
	}
	b := bc.Blocks[TOP]
	pow := new_pow(b)
	nonce, hash := pow.search()
	b.nonce = nonce
	b.Hash = hash
	b.Mined = true
	if TOP+1 < len(bc.Blocks) {
		bc.Blocks[TOP+1].PrevHash = hash
	}
	TOP++
	return true
}
