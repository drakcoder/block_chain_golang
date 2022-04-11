package main

type Block struct {
	BlockId    int    `json:blockid`
	StringData string `json:stringdata`
	Data       []byte `json:data`
	PrevHash   []byte `json:prevhash`
	Hash       []byte `json:hash`
	nonce      int    `json:nonce`
	Mined      bool   `json:mined`
}
