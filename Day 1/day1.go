package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var inventoryTotals []int
	highestElf := 0

	counter := 0
	inventoryTotals = append(inventoryTotals, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if inventoryTotals[counter] > highestElf {
				highestElf = inventoryTotals[counter]
			}
			counter++
			inventoryTotals = append(inventoryTotals, 0)
		} else {
			toAdd, _ := strconv.Atoi(scanner.Text())
			inventoryTotals[counter] += toAdd
		}
	}

	invSlice := inventoryTotals[:]
	sort.Sort(sort.Reverse(sort.IntSlice(invSlice)))

	highest3 := inventoryTotals[0] + inventoryTotals[1] + inventoryTotals[2]
	fmt.Println(highest3)
}
