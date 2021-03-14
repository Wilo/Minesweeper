package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
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
		return nil, errors.New("only numbers between 1 and 26")
	}
	for row := 0; row < size; row++ {
		rows = append(rows, string(rune('A'+row)))
	}
	return
}

func GenerateColumns(size int) (columns []string, err error) {
	if size < 1 || size > 26 {
		return nil, errors.New("only numbers between 1 and 26")
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

func index(target string, array []string) (index int) {
	for index, value := range array {
		if value == target {
			return index
		}
	}
	return
}

func GetCoordinates(size int, coordinates string) (col int, row int, err error) {
	re, err := regexp.Compile("([A-Z]+)([0-9]+)")
	if err != nil {
		return
	}
	letters, _ := GenerateRows(size)
	found := re.MatchString(coordinates)
	if found {
		yxArray := re.FindStringSubmatch(coordinates)[1:]
		y, x := yxArray[0], yxArray[1]
		col, _ := strconv.Atoi(x)
		row = index(y, letters)
		return row, col, nil
	}
	return
}

func PutMines(size int, board [][]string, mines int, minesList []string) [][]string {
	for _, coordinate := range minesList {
		col, row, _ := GetCoordinates(size, coordinate)
		if board[col][row] != "*  " {
			board[col][row] = "*  "
		}
	}
	return board
}

func PutSignals(row int, column int, board [][]string) [][]string {
	var map_coordinates []int = []int{-1, 0, 1}
	for y := 0; y < column; y++ {
		for x := 0; x < row; x++ {
			if board[y][x] == "*  " {
				for _, i := range map_coordinates {
					for _, j := range map_coordinates {
						if 0 <= y+i && y+i <= row-1 && 0 <= x+j && x+j <= column-1 && board[y+i][x+j] != "*  " {
							index, _ := strconv.Atoi(board[y+i][x+j])
							tmpvalue, _ := strconv.Atoi(board[y+i][x+j])
							value := strconv.Itoa(tmpvalue + 1)
							board[y+i][x+j] = SetSpaces(index+1, value)
						}
						//fmt.Println(board[y+i][x+j])
						//fmt.Println(y, i, x, j)
					}
				}
			}
		}
	}
	return board
}

func main() {
	row, column := 5, 5
	//board := MakeBoard(row, column, ".")
	board := MakeBoard(row, column, "0")
	//DrawBoard(row, column, board)
	//fmt.Println(MapOfMines(10, 10))
	//fmt.Println(GetCoordinates(10, "A1"))
	var mines_list []string = []string{"B3", "D4"}
	fmt.Println(PutMines(10, board, 2, mines_list))
	fmt.Println(PutSignals(row, column, board))
}
