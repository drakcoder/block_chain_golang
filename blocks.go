package main

type Block struct {
	Data     []byte
	PrevHash []byte
	Hash     []byte
	nonce    int
	Mined    bool
}
