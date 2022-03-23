package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type ProofOfWork struct {
	Target *big.Int
	Block  *Block
}

func (pow *ProofOfWork) prepare_hash(nonce int64) []byte {
	ph := bytes.Join([][]byte{pow.Block.Data, pow.Block.PrevHash, []byte(strconv.FormatInt(nonce, 16))}, []byte{})
	return ph
}

func (pow *ProofOfWork) search() (int, []byte) {
	nonce := 0
	max := math.MaxInt64
	var hashInt big.Int
	for nonce <= max {
		tbh := pow.prepare_hash(int64(nonce))
		hash := sha256.Sum256(tbh)
		hashInt.SetBytes(hash[:])
		fmt.Println(hashInt)
		if hashInt.Cmp(pow.Target) == -1 {
			fmt.Println("found", nonce)
			return nonce, hash[:]
		}
		nonce++
	}
	return 0, []byte{0}
}
