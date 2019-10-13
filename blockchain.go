package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Transaction string
	PrevPointer *Block
	PrevHash    []byte
}

func calculate_hash(b *Block) []byte {
	prev_hash_string := fmt.Sprintf("%p", b.PrevHash)
	curr_hash_string := b.Transaction + prev_hash_string
	bytes := sha256.Sum256([]byte(curr_hash_string))
	return bytes[:]
}

func InsertBlock(transaction string, chainHead *Block) *Block {
	if chainHead == nil {
		chainHead = &Block{transaction, nil, nil}
		return chainHead
	} else {
		temp_block := &Block{transaction, chainHead, nil}
		temp_block.PrevHash = calculate_hash(chainHead)
		return temp_block
	}
}

func ListBlocks(chainHead *Block) {
	if chainHead == nil {
		fmt.Printf("Blockchain is Empty!!")
	} else {
		for i := chainHead; i != nil; i = i.PrevPointer {
			fmt.Print("Transaction: ", i.Transaction, "\n")
		}
	}
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	if chainHead == nil {
		fmt.Printf("Blockchain is Empty!!")
	} else {
		for i := chainHead; i != nil; i = chainHead.PrevPointer {
			if chainHead.Transaction == oldTrans {
				chainHead.Transaction = newTrans
			} else {
				chainHead = chainHead.PrevPointer
			}
		}
	}
}

func VerifyChain(chainHead *Block) {
	invalid := 0
	for i := chainHead; i.PrevPointer != nil; i = i.PrevPointer {
		if bytes.Compare(i.PrevHash, calculate_hash(i.PrevPointer)) != 0 {
			fmt.Print("block chain is compromised ", i.PrevPointer.Transaction, " is fake", "\n")
			invalid = 1
		}
	}
	if invalid == 0 {
		fmt.Print("block chain is valid ")
	}
}
