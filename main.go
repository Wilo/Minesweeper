package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func SetSpaces(index int, value string) string {
	if len(fmt.Sprintf("%d", index)) > 1 && isNumeric(value) {
		return fmt.Sprintf("%s ", value)
	}
	return fmt.Sprintf("%s  ", value)
}

// We generate a list with the letters of the alphabet
// Max length = 26
// Min length = 1

// example:
// >>> generate_rows(3)
// ['A', 'B', 'C']
// >>> generate_rows(0)
// []
func GenerateRows(size int) (rows []string, err error) {
	if size < 1 || size > 26 {
		return nil, errors.New("Only numbers between 1 and 26!")
	}
	for row := 0; row < size; row++ {
		rows = append(rows, string(rune('A'+row)))
	}
	return
}

func GenerateColumns(size int) (columns []string, err error) {
	if size < 1 || size > 26 {
		return nil, errors.New("Only numbers between 1 and 26!")
	}
	columns = append(columns, "   ")
	for column := 0; column < size; column++ {
		columns = append(columns, SetSpaces(column+1, fmt.Sprintf("%d", column+1)))
	}
	return
}

/**
Check OS: If your OS is Ms Windows the command for clear screen is cls, but if your OS is
          Apple OSX or any Linux Kernel Distro then the command is clear
*/
func ClearScreen() {
	var command string
	if runtime.GOOS == "windows" {
		command = "cls"
	} else {
		command = "clear"
	}
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DifficultyCalculator(row int, column int, difficulty string) (percentage float32) {
	switch difficulty {
	case "F":
		percentage = 0.10 // 10% Easy
	case "M":
		percentage = 0.15 // 15% Medium
	case "D":
		percentage = 0.20 // 20% Hard
	default:
		percentage = 0.25 // 25% Extreme Obviously is X
	}
	return float32(row*column) * percentage
}

// https://codereview.stackexchange.com/questions/60074/in-array-in-go
func in(target string, array []string) (result bool) {
	for index := range array {
		if result = array[index] == target; result {
			return
		}
	}
	return
}

func MapOfMines(size int, numberOfMines int) (totalList []string) {
	minValue := 0
	letters, _ := GenerateRows(size)
	var (
		number int
		letter string
		code   string
	)
	rand.Seed(time.Now().UnixNano())
	for {
		number = rand.Intn((size - 1) - minValue)
		letter = letters[rand.Intn(len(letters))]
		code = fmt.Sprintf("%s%d", letter, number)
		if len(totalList) < numberOfMines {
			//Validation if list is not empty
			if len(totalList) > 0 {
				// Don't repeat random codes
				if !in(code, totalList) {
					totalList = append(totalList, code)
				}
			} else {
				totalList = append(totalList, code)
			}
		} else {
			break
		}
	}
	// Prepend size of string into total Slice
	totalList = append([]string{fmt.Sprintf("%d", size)}, totalList...)
	return
}

func MakeBoard(row int, column int, character string) (board [][]string) {
	for y := 0; y < row; y++ {
		board = append(board, []string{})
		for x := 0; x < column; x++ {
			board[y] = append(board[y], SetSpaces(x, character))
		}
	}
	return
}

func DrawBoard(row int, column int, board [][]string) {
	rows, _ := GenerateRows(row)
	columns, _ := GenerateColumns(column)
	fmt.Println(strings.Join(columns, ""))
	for index, rowBoard := range board {
		fmt.Println(fmt.Sprintf("%s  %s", rows[index], strings.Join(rowBoard, "")))
	}
}

func main() {
	//row, column := 8, 8
	//board := MakeBoard(row, column, ".")
	//DrawBoard(row, column, board)
	fmt.Println(MapOfMines(10, 10))
}
