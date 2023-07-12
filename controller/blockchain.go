package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const Difficulty = 1

type Block struct {
	PrevHash   string
	Index      int
	Nonce      string
	Data       int
	Difficulty int
	TimeStamp  string
	Hash       string
}

func GenesisBlock() *Block {
	genesisBlock := &Block{}

	genesisBlock.PrevHash = ""
	genesisBlock.Index = 0
	genesisBlock.Nonce = ""
	genesisBlock.Data = 0
	genesisBlock.Difficulty = Difficulty
	genesisBlock.TimeStamp = time.Now().String()
	genesisBlock.Hash = generateHash(genesisBlock)

	return genesisBlock
}

func NewBlock(prevBlock *Block, data int) *Block {
	newBlock := &Block{}

	newBlock.PrevHash = prevBlock.Hash
	newBlock.Index = prevBlock.Index + 1
	newBlock.Data = data
	newBlock.Difficulty = Difficulty
	newBlock.TimeStamp = time.Now().String()

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !isHashValid(generateHash(newBlock), Difficulty) {
			fmt.Println(generateHash(newBlock), "do more work")
			time.Sleep(time.Second)
		} else {
			fmt.Println(generateHash(newBlock), "🎉 work done")
			newBlock.Hash = generateHash(newBlock)
			break
		}
	}

	return newBlock
}

func generateHash(block *Block) string {
	record := block.PrevHash + strconv.Itoa(block.Index) + block.Nonce + strconv.Itoa(block.Data) + block.TimeStamp
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data int) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(prevBlock, data)

	if isBlockValid(block, prevBlock) {
		bc.Blocks = append(bc.Blocks, block)
	}
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func isBlockValid(block, prevBlock *Block) bool {
	if prevBlock.Index+1 != block.Index {
		return false
	}
	if prevBlock.Hash != block.PrevHash {
		return false
	}
	if generateHash(block) != block.Hash {
		return false
	}

	return true
}
