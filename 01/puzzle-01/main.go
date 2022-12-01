package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileData := readFile("../input.txt")
	maxCalories := countCalories(string(fileData))

	fmt.Println(maxCalories)

}

func readFile(filePath string) string {
	fileData, err := os.ReadFile(filePath)
	checkError(err)

	return string(fileData)
}

func checkError(e error) {
	if e != nil {
		fmt.Println("hit error")
		panic(e)
	}
}

func countCalories(data string) int {
	elves := getAllElvesCalories(data)
	elvesTotalCalories := sumElvesCalories(elves)
	maxCalories := getHighestCalorieElf(elvesTotalCalories)

	return maxCalories

}

func getAllElvesCalories(data string) map[int][]int {
	var elves = make(map[int][]int)

	values := strings.Split(data, "\n")

	index := 0

	for _, value := range values {
		if value == "" {
			index++
			continue
		}

		valueInt, err := strconv.Atoi(value)
		checkError(err)

		elves[index] = append(elves[index], valueInt)
	}

	return elves
}

func sumElvesCalories(elves map[int][]int) map[int]int {
	var summedElves = make(map[int]int)

	for elf, values := range elves {
		totalCalories := 0
		for _, calories := range values {
			totalCalories += calories
		}

		summedElves[elf] = totalCalories
	}

	return summedElves
}

func getHighestCalorieElf(elves map[int]int) int {
	maxCalories := 0

	for _, calories := range elves {
		if calories > maxCalories {
			maxCalories = calories
		}
	}

	return maxCalories
}
