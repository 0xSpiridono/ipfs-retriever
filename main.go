package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/phatngluu/ipfs-retriever/utils"
)

func main() {
	godotenv.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ipfs, err := utils.NewIPFS(ctx)
	if err != nil {
		panic(err)
	}

	cids := os.Args[1:]

	for _, cid := range cids {
		err := ipfs.Retrieve(cid)
		panic(err)
	}
}
