package blockchain

import (
	"bytes"
	"fmt"
)

type CryptoBlockChain struct {
	blockchain []*CryptoBlock
	difficulty int
}

func NewCryptoBlockChain() *CryptoBlockChain {
	cbc := new(CryptoBlockChain)
	cbc.blockchain = []*CryptoBlock{cbc.startGenesisBlock()}
	cbc.difficulty = 1
	return cbc
}

func (cbc *CryptoBlockChain) startGenesisBlock() *CryptoBlock {
	genesisBlock := NewCryptoBlock(0, "01/01/2021", []byte("Genesis Block"))
	return genesisBlock
}

func (cbc *CryptoBlockChain) obtainLatestBlock() *CryptoBlock {
	return cbc.blockchain[len(cbc.blockchain) - 1]
}

func (cbc *CryptoBlockChain) AddNewBlock(newBlock *CryptoBlock) {
	newBlock.precedingHash = cbc.obtainLatestBlock().hash
	cbc.proofOfWork(newBlock, cbc.difficulty)
	cbc.blockchain = append(cbc.blockchain, newBlock)
}

func (cbc *CryptoBlockChain) PrintBlockChain() {
	fmt.Println(len(cbc.blockchain))
	for i, v := range cbc.blockchain {
		fmt.Println("Index: ", i, " Value:", v.hash, " Data:", string(v.data))
	}
}

func (cbc *CryptoBlockChain) CheckChainValidity() (bool, error){
	for i := 0; i < len(cbc.blockchain); i++ {
		currentBlock := cbc.blockchain[i]
		//precedingBlock := cbc.blockchain[i-1]

		if currentBlock.hash != currentBlock.computeHash() {
			fmt.Println("BLOCKCHAIN VALIDATION: currentBlockHash: ", currentBlock.hash, " computeHash: ", currentBlock.computeHash())
			return false, nil
		}
		if i > 0 && currentBlock.precedingHash != cbc.blockchain[i-1].hash {
			fmt.Println("BLOCKCHAIN VALIDATION: currentBlock.PrecedingHash: ", currentBlock.precedingHash, " blockchain preceding hash ", cbc.blockchain[i-1].hash)
			return false, nil
		}
	}
	return true, nil
}

func (cbc *CryptoBlockChain) proofOfWork(newBlock *CryptoBlock, difficulty int) {

	var a []byte

	for i:= 0; i < difficulty; i++ {
		a = append(a, 0)
	}
	fmt.Println("start mining ...")
	fmt.Println("COMPUTING HASH: currentHash", newBlock.hash[0:difficulty], " targetHash: ", a, " nonce: ", newBlock.nonce)
	for bytes.Compare(newBlock.hash[0:difficulty], a) != 0 {
		fmt.Println("COMPUTING HASH: currentHash", newBlock.hash[0:difficulty], " targetHash: ", a, " nonce: ", newBlock.nonce)
		newBlock.nonce++
		newBlock.hash = newBlock.computeHash()
	}
}
