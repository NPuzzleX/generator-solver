package main

import (
	"Sudoku"
	"fmt"
	"log"
	"time"
)

func main() {
	testSudoku()
}

func testSudoku() {
	const n = 4

	for count := 0; count < 2; count++ {
		dt, err := Sudoku.Generate(n, time.Now().UnixMicro(), true, 150, 6)

		if err != nil {
			log.Fatal(err)
		}

		for i, e := range dt {
			if (i != 0) && (i%n == 0) {
				fmt.Println()
			}
			for j, e2 := range e {
				if (j != 0) && (j%n == 0) {
					fmt.Print("| ")
				}
				if n > 3 {
					if e2 < 10 {
						fmt.Print(" ")
					}
				}
				fmt.Print(e2)
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println("------------")
		fmt.Println()
	}

	/*
		seed := [][]int{
			{0, 0, 5, 9, 2, 0, 0, 0, 6},
			{8, 1, 0, 0, 0, 0, 7, 0, 0},
			{0, 2, 0, 7, 3, 0, 9, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 8, 0, 0},
			{0, 6, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 7, 0, 0, 0, 5, 0, 0},
			{6, 0, 4, 0, 0, 0, 0, 0, 7},
			{0, 0, 0, 5, 0, 0, 6, 9, 0},
		}

		sol, _ := sudoku.Solver(seed)
		for i, e := range sol {
			if (i != 0) && (i%n == 0) {
				fmt.Println()
			}
			for j, e2 := range e {
				if (j != 0) && (j%n == 0) {
					fmt.Print("| ")
				}
				if n > 3 {
					if e2 < 10 {
						fmt.Print(" ")
					}
				}
				fmt.Print(e2)
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	*/
}
