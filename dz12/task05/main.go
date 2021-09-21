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

	//f, err := os.Open("Log.txt")
	//if err != nil {
	//	fmt.Println("Not open file", err)
	//	return
	//}
	//defer f.Close()
	//
	//fi, err := f.Stat()
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//size := fi.Size()
	//buf := make([]byte, size)
	//_, err = f.Read(buf)
	//if err != nil {
	//	fmt.Println("Not red file", err)
	//	return
	//}
	//fmt.Println(string(buf))
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
