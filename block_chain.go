package main

import "fmt"

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) add_block(data string) {
	b := new_block(data)
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, new_gen_block())
		LENGTH++
	}
	b.PrevHash = bc.Blocks[LENGTH-1].Hash
	bc.Blocks = append(bc.Blocks, b)
	LENGTH++
}

func (bc *BlockChain) mine() {
	if TOP >= LENGTH-1 {
		fmt.Println("All Blocks mined, wait for incoming blocks")
		return
	}
	b := bc.Blocks[TOP+1]
	pow := new_pow(b)
	nonce, hash := pow.search()
	b.nonce = nonce
	b.Hash = hash
	b.Mined = true
	if TOP+2 < len(bc.Blocks) {
		bc.Blocks[TOP+2].PrevHash = hash
	}
	TOP++
}

func (bc *BlockChain) display_chain() {
	for i := 1; i < LENGTH; i++ {
		fmt.Println(bc.Blocks[i].Data, bc.Blocks[i].Mined, bc.Blocks[i].Hash, bc.Blocks[i].PrevHash)
	}
}
