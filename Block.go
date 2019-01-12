package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	PervBlockHash []byte
	Data []byte
	Hash []byte
}

func (block *Block)setHash(){
//	将时间戳转化多维数组
timestamp := strconv.FormatInt(block.Timestamp,10)
fmt.Println("timestamp:",timestamp)

timeBytes := []byte(timestamp)
fmt.Println("timeBytes:",timeBytes)

//  拼接所有属性
blockBytes := bytes.Join([][]byte{timeBytes,block.PervBlockHash,block.Data},[]byte{})
fmt.Println("blockBytes:",blockBytes)
// 	生成hash
hash := sha256.Sum256(blockBytes)
block.Hash = hash[:]

}

func NewBlock (data string,pervBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(),pervBlockHash,[]byte(data),[]byte{}}

	block.setHash()
	return block
	}
