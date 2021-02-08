package blockchain

import (
	"crypto/sha256"
	"strconv"
)

type CryptoBlock struct {
	index int
	timestamp string
	data []byte
	precedingHash [32]byte
	hash [32]byte
	nonce int
}

func NewCryptoBlock(index int, timestamp string, data []byte ) *CryptoBlock {
	cb := new(CryptoBlock)
	cb.index = index
	cb.timestamp = timestamp
	cb.data = data
	cb.precedingHash = [32]byte{0}
	cb.hash = cb.computeHash()
	cb.nonce = 0
	return cb
}

func (cb *CryptoBlock) computeHash() [32]byte {
	s := strconv.Itoa(cb.index) + string(cb.precedingHash[:]) + cb.timestamp + string(cb.data) + strconv.Itoa(cb.nonce)
	b := []byte(s)
	return sha256.Sum256(b)
}
