package main

import "fmt"

func main() {

	var task1 = "Watch GoLang course!"
	var task2 = "Terraform"
	var task3 = "Cloud Computing assignment"
	var task4 = "Laptop sleeve bag"
	task5 := "Buy Scuderia Ferrari Merch"

	var taskItems = []string{task1, task2, task3, task4, task5}

	fmt.Println("###### Welcome to our ToDo list app ######")

	fmt.Println("List of all my To-Dos :")

	for index, task := range taskItems {
		//fmt.Println(index+1, ".", task)
		fmt.Printf("%d: %s\n", index+1, task)
	}

}
