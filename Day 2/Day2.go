package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var score = 0
var hands [3]string = [3]string{"X", "Y", "Z"}
var winners [3]string = [3]string{"C", "A", "B"}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	x := 0
	for scanner.Scan() {
		outcome(scanner.Text(), x)
	}
	fmt.Println(score)
}

func outcome(matchup string, x int) {
	hands := strings.Split(matchup, " ")
	switch x % 3 {
	//Rock
	case 0:
		if hands[0] == winners[1] {
			score += 6
		} else if hands[0] == winners[2] {
			score += 3
		}
		score += 2
	case 1:
		if hands[0] == winners[0] {
			score += 6
		} else if hands[0] == winners[1] {
			score += 3
		}
		score += 1
	case 2:
		if hands[0] == winners[2] {
			score += 6
		} else if hands[0] == winners[0] {
			score += 3
		}
		score += 3
	}
}
