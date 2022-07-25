package main

import (
	"fmt"
	"math/rand"
)

const creditsToGraduate = 121

func main(){
	credits := 3
	major := ""
	degree := ""

	fmt.Println("			Major     		  CoursesToTakePerYear	YearsToFinish  	Degree type PricePerYear")
	fmt.Println("==================================================================================================================")

	for count := 0; count <10; count++ {
		switch rand.Intn(5) {
		case 0:
			major = "Information Technology Management/Computer Science"
		case 1:
			major = "Information & Communication Technology"
		case 2:
			major = "Law"
		case 3: 
			major = "Business Administration"
		case 4: 
			major = "Tourism and Hospitality Management"
		}
	
	courses := rand.Intn(11) + 	4
	duration := creditsToGraduate / courses / credits
	price := courses * 600

	if rand.Intn(2) == 1 {
		degree = "Dual"
		price = price + 300
	} else {
		degree ="Single"
	}
	fmt.Printf("%-56v %5v %15v %18v       $%3v \n", major, courses,duration, degree, price)
	}
}