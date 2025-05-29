package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	task := flag.String("task", "Brush teeth", "Task to do")
	add := flag.Bool("add", false, "Add Flag")
	remove := flag.Bool("remove", false, "Remove Flag")
	update := flag.Bool("update", false, "Update Flag")
	read := flag.Bool("read", false, "Read Flag")
	updatedTask := flag.String("updatedTask", "Brush teeth", "Task to be updated")

	flag.Parse()

	if *add {
		addTask(*task)
	} else if *remove {
		removeTask(*task)
	} else if *update {
		updateTask(*task, *updatedTask)
	} else if *read {
		readTasks()
	}
}

func addTask(task string) {
	// read from file
	val, err := os.ReadFile("tasks.txt")
	if err != nil {
		log.Fatal(err)
	}
	toAdd := string(val)
	if strings.Contains(toAdd, task+"\n") == true {
		fmt.Println("Already Exists!")
		fmt.Println(toAdd)
		return
	}
	toAdd += task + "\n"
	writeInFile(toAdd)
}
func removeTask(task string) {
	// read from file
	val, err := os.ReadFile("tasks.txt")
	if err != nil {
		log.Fatal(err)
	}
	toAdd := string(val)
	if strings.Contains(toAdd, task+"\n") == false {
		fmt.Println("Not Exists!")
		fmt.Println(toAdd)
		return
	}
	toAdd = strings.ReplaceAll(toAdd, task+"\n", "")
	writeInFile(toAdd)
}
func updateTask(task string, updatedTask string) {
	// read from file
	val, err := os.ReadFile("tasks.txt")
	if err != nil {
		log.Fatal(err)
	}
	toAdd := string(val)
	if strings.Contains(toAdd, task+"\n") == false {
		fmt.Println("Not Exists!")
		fmt.Println(toAdd)
		return
	}
	toAdd = strings.Replace(toAdd, task+"\n", updatedTask+"\n", 1)
	writeInFile(toAdd)
}

func readTasks() {
	// read from file
	val, err := os.ReadFile("tasks.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(val))
}

func writeInFile(toAdd string) {
	err1 := os.WriteFile("tasks.txt", []byte(toAdd), 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(toAdd)
}
