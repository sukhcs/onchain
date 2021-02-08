package main

import (
	"fmt"
	"onchain/blockchain"
)

func main() {

	onChain := blockchain.NewCryptoBlockChain()
	onChain.AddNewBlock(blockchain.NewCryptoBlock(1, "01/06/2020", []byte("{sender: 'A', recipient: 'B', quantity: 50}")))
	onChain.AddNewBlock(blockchain.NewCryptoBlock(2, "01/07/2020", []byte("{sender: 'B', recipient: 'A', quantity: 100}")))
	onChain.PrintBlockChain()
	valid, err := onChain.CheckChainValidity()
	if !valid && err == nil {
		fmt.Println("blockchain tampered!")
	} else {
		fmt.Println("blockchain valid!")
	}

	/*
	// set up a p2p node
	p2pNode := p2p.NewOnChainNode()

	h2 := p2pNode.GetHost2()
	ctx := p2pNode.GetContext()
	// This connects to public bootstrappers
	for _, addr := range dht.DefaultBootstrapPeers {
		pi, _ := peer.AddrInfoFromP2pAddr(addr)
		// We ignore errors as some bootstrap peers may be down
		// and that is fine.
		fmt.Println("addr:", addr)
		error := h2.Connect(ctx, *pi)

	}

	fmt.Println("New host ID: ", p2pNode.GetHost().ID(), p2pNode.GetHost2().ID())

	 */
}
