package main

import "fmt"


func main() {
	const BaseRate = 5.50
	const TaxRate = 0.12
	const DistanceRate = 2.0
	const FragileFee = 0.2

	fmt.Print("Введите имя:")
	var name string
	fmt.Scan(&name)

	fmt.Print("Вес груза(Кг):")
	var weight float64
	fmt.Scan(&weight)

	fmt.Print("Дистанция(км):")
	var Distance float64
	fmt.Scan(&Distance)

	fmt.Print("Хрупких упаковок:")
	var Fragile float64
	fmt.Scan(&Fragile)

	BaseRate2 := (weight*BaseRate)*(1+FragileFee*Fragile) + (Distance * DistanceRate)
	total := BaseRate2 * TaxRate

	fmt.Println("---Отчет о доставке---")
	fmt.Printf(" Отправитель:%s\n Итоговая стоимость: %.2f\n", name, total)

}
