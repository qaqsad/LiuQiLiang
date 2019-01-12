package main

import (
	"fmt"
	"liuqiliang.com/BLC"
)

func main() {
 block := BLC.NewBlock("Genenis Block",[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
 fmt.Println(block)
 fmt.Println(string(block.Data))
 fmt.Printf("%x\n",block.PervBlockHash)
 fmt.Printf("%x",block.Hash)
}
