package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bontusss/igboapi-go"
)

func main() {
	key := ""
	search(key)
	find(key)
}	

func search(key string) {
	fmt.Println("Searcing...")
	client := igboapi.NewClient(key)
	t, err := client.TermService.Search(context.Background(), &igboapi.TermParams{Keyword: "bia"})
	// fmt.Println(client)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(len(*t)) // returns 3
}

func find(key string) {
	fmt.Println("Finding ...")
	client := igboapi.NewClient(key)
	t, err := client.TermService.Find(context.Background(), "5f90c35e49f7e863e92b7269")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(t.Word) // returns bị̀à
}
