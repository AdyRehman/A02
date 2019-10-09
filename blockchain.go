package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	transaction string
	prevPointer *Block
	prevHash    []byte
}

func calculate_hash(b *Block) []byte {
	prev_hash_string := fmt.Sprintf("%p", b.prevHash)
	curr_hash_string :=  b.transaction + prev_hash_string
	bytes := sha256.Sum256([]byte(curr_hash_string))
	return bytes[:]
}

func InsertBlock(transaction string, chainHead *Block) *Block {
	if chainHead == nil {
		chainHead = &Block{transaction, nil, nil}
		return chainHead
	} else {
		temp_block := &Block{transaction, chainHead, nil}
		temp_block.prevHash = calculate_hash(chainHead)
		return temp_block
	}
}

func ListBlocks(chainHead *Block) {
	if chainHead == nil {
		fmt.Printf("Blockchain is Empty!!")
	} else {
		for i := chainHead; i != nil; i = i.prevPointer {
			fmt.Print("Transaction: ", i.transaction, "\n")
		}
	}
}

func  ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	if chainHead == nil {
		fmt.Printf("Blockchain is Empty!!")
	} else {
		for i := chainHead; i != nil; i = chainHead.prevPointer {
			if chainHead.transaction == oldTrans {
				chainHead.transaction = newTrans
			} else {
				chainHead = chainHead.prevPointer
			}
		}
	}
}

func VerifyChain(chainHead *Block) {
	invalid := 0
	for i := chainHead; i.prevPointer != nil; i = i.prevPointer {
		if bytes.Compare(i.prevHash, calculate_hash(i.prevPointer)) != 0 {
			fmt.Print("block chain is compromised ", i.prevPointer.transaction, " is fake", "\n")
			invalid = 1
		}
	}
	if invalid == 0 {
		fmt.Print("block chain is valid ")
	}
}
