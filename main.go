package main

import "fmt"

func main() {
	onChain := NewCryptoBlockChain()
	onChain.addNewBlock(NewCryptoBlock(1, "01/06/2020", []byte("{sender: 'A', recipient: 'B', quantity: 50}")))
	onChain.addNewBlock(NewCryptoBlock(2, "01/07/2020", []byte("{sender: 'B', recipient: 'A', quantity: 100}")))
	onChain.printBlockChain()
	valid, err := onChain.checkChainValidity()
	if !valid && err == nil {
		fmt.Println("blockchain tampered!")
	} else {
		fmt.Println("blockchain valid!")
	}
}
