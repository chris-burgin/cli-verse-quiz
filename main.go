package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/chrisburgin95/cli-verse-search/esv"
	"github.com/rapidloop/skv"
)

func newStore() skv.KVStore {
	store, err := skv.Open("./verse.db")
	if err != nil {
		panic(err)
	}

	return *store
}

func setup(store skv.KVStore) {
	fmt.Print("Verse to setup: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	store.Delete("verse")
	store.Put("verse", text)
}

func getVerse(store skv.KVStore) string {

	var verse string
	err := store.Get("verse", &verse)
	if err != nil {
		panic(err)
	}

	res, err := esv.GetVerse(url.QueryEscape(verse))
	if err != nil {
		panic(err)
	}

	return res
}

func view(store skv.KVStore) {
	println(getVerse(store))
}

func practice(store skv.KVStore) {
	println("to do")
}

func main() {
	store := newStore()
	defer store.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What action would you like to take? (setup, view, practice):")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	if text == "setup\n" {
		setup(store)
	} else if text == "practice\n" {
		practice(store)
	} else if text == "view\n" {
		view(store)
	} else {
		println("Unkown Command. Please try 'verse', 'view', or 'practice'.")
	}

}
