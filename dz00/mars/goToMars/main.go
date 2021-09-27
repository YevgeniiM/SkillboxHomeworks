package main

import "fmt"

func main() {
	//minDistance := 56e6
	//maxDistance := 401e6

	//lightSpeed := big.NewInt(299792)
	//secondsPerDay := big.NewInt(86400)
	//
	//distance := new(big.Int)
	//distance.SetString("24000000000000", 10)
	//fmt.Println("Расстояние до галактики Андромеды составляет", distance, "км.") // Выводит: Расстояние до галактики Андромеды составляет 24000000000000000000 км.
	//
	//seconds := new(big.Int)
	//seconds.Div(distance, lightSpeed)
	//fmt.Println(seconds,distance)
	//
	//days := new(big.Int)
	//days.Div(seconds, secondsPerDay)
	//
	//
	//fmt.Println("Это", days, "дня полета на скорости света.") // Выводит: Это 926568346 дня полета на скорости света.

	const distance = 236000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365

	years := distance / lightSpeed / secondsPerDay / daysPerYear

	fmt.Println("Расстояние в световых годах до Карликовой галактики в Большом Псе составляет:", years)
}
