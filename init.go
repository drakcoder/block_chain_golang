package main

import (
	"math/big"
)

var TOP int = 0
var LENGTH int = 0

func new_pow(b Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-22)
	return &ProofOfWork{Target: target, Block: b}
}

func new_block(data string) Block {
	return Block{StringData: data, Data: []byte(data), Mined: false}
}

func new_gen_block(data string) Block {
	return Block{StringData: data, Data: []byte(data), PrevHash: []byte{}, Mined: false}
}

func new_block_chain() *BlockChain {
	return &BlockChain{Blocks: []Block{}}
}
