package main

import "fmt"

func main() {
	const hourPerDay = 24
	var day, distance int
	fmt.Println("За сколько дней вы хотите долететь до Марса?")
	_, _ = fmt.Scan(&day)
	fmt.Println("Какое расстояние до Марса будет в момент старта?") // км
	_, _ = fmt.Scan(&distance)
	fmt.Printf("Вам необходимо лететь со скоростью %vкм/ч", distance/(day*hourPerDay))
}
