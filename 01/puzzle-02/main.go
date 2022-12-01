package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileData := ReadFile("../input.txt")
	maxCalories := countCalories(string(fileData))

	fmt.Println(maxCalories)

}

func ReadFile(filePath string) string {
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
	sortedElvesCalories := getSortedElvesCalories(elvesTotalCalories)
	topThreeElvesCalories := getTopElvesCalories(sortedElvesCalories, 3)
	summedTopThreeElvesCalories := sumCalories(topThreeElvesCalories)

	return summedTopThreeElvesCalories

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
		summedElves[elf] = sumCalories(values)
	}

	return summedElves
}

func sumCalories(calories []int) int {
	totalCalories := 0

	for _, calorie := range calories {
		totalCalories += calorie
	}

	return totalCalories
}

func getSortedElvesCalories(elves map[int]int) []int {
	var elvesCalories []int

	for _, calories := range elves {
		elvesCalories = append(elvesCalories, calories)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elvesCalories)))

	return elvesCalories
}

func getTopElvesCalories(elvesCalories []int, topNum int) []int {
	return elvesCalories[:topNum]
}
