package main

import (
	"fmt"
	"net/http"
)

var task1 = "Watch GoLang course!"
var task2 = "Terraform"
var task3 = "Cloud Computing assignment"
var task4 = "Laptop sleeve bag"
var task5 = "Buy Scuderia Ferrari Merch"
var taskItems = []string{task1, task2, task3, task4, task5}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "###### Welcome to our To-Do list app built using Go ######"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func printTasks(taskItems []string) {
	fmt.Println("List of all my To-Dos :")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
