package main

import (
	"fmt"
	"strings"
)

func generate(open, closing int, queue string, result *[]string) {
	if open == 0 {
		*result = append(*result, queue+strings.Repeat(")", closing))
		return
	}
	if closing > open {
		generate(open, closing-1, queue+")", result)
	}
	generate(open-1, closing, queue+"(", result)
}

func main() {
	result := make([]string, 0)
	n := 4
	generate(n, n, "", &result)
	fmt.Println(result)
	//var b bytes.Buffer
	//var text, varText, s, s1, variant1 string
	//var number int
	//element1, element2 := "(", ")"
	//text = "Введите число для которого необходимо сгенерировать все возможные комбинации круглых скобок: \n"
	//fmt.Print(text)
	//b.WriteString(text)
	//fmt.Scan(&number)
	//b.WriteString(fmt.Sprintf("%v\n", number))
	//iterazii := int(math.Pow(2, float64(number-1)))
	//numberVar1, numberVar2 := number, number
	//
	//for i := 0; i < iterazii; i++ {
	//	for numberVar1 >= 0 {
	//
	//		if numberVar1 == 0 {
	//			variant1 += strings.Repeat(element2, numberVar2)
	//			break
	//
	//		}
	//		if numberVar2 != 0 {
	//			variant1 += element2
	//			numberVar2--
	//		}
	//	}
	//
	//	//variant2 = variant1
	//	varText += variant1 + ","
	//	if len(variant1)/2-i >= 0 {
	//		s = variant1[:len(variant1)/2-i]
	//		s1 = variant1[len(variant1)/2+i:]
	//		fmt.Println(s)
	//		fmt.Println(s1)
	//		numberVar2 = i
	//		numberVar1 = i
	//
	//		variant1 = s1 + s
	//	}
	//
	//}
	//text = varText + "\n"
	//fmt.Print(text)
	//b.WriteString(text)
	//fileName := "Comb.txt"
	//if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
	//	panic(err)
	//}
	//file, err := os.Open(fileName)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//factBytes, err := ioutil.ReadAll(file)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("=============")
	//fmt.Println(string(factBytes))
}
