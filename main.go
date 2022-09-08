package main

import (
	"log"

	"github.com/junkeWu/post/downloader"
)

func main() {
	log.Println("start get posts...")
	err := downloader.GetPostsAndWriteFile("./data/post.json")
	if err != nil {
		log.Printf("GetPostsAndWriteFile Failed,err=%v", err)
		return
	}
	log.Println("get posts success...")
}
