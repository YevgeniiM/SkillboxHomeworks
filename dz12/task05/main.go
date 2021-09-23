package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	/*
		переписать задачи из урока 2 и 3 на пакет ioutil
	*/
	fileName := "Log.txt"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	factBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(factBytes))
}
