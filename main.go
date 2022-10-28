package main

import (
	"fmt"
	"github.com/akrylysov/pogreb"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	files, err := filepath.Glob("../4bytes/signatures/*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(files)) // contains a list of all files in the current directory
	db, err := pogreb.Open("4bytes.db", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	for i := range files {
		if i%10000 == 0 {
			fmt.Println(i)
		}
		content, err := ioutil.ReadFile(files[i])
		tmp := strings.Split(files[i], "/")
		key := tmp[len(tmp)-1]
		if err != nil {
			log.Fatal(err)
		}
		err = db.Put([]byte(key), []byte(content))
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("finish")

}
