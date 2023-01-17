package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Error struct {
	message string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s", e.message)
}

var (
	dt               = time.Now()
	err              = errors.New("ERROR - ")
	ErrorWrongFormat = fmt.Errorf("%w wrong format; Error time - %s ", err, dt.Format("01-02-2006 15:04"))
)

func div(a, b float64) (float64, error) {
	defer fmt.Println(ErrorWrongFormat)
	return a / b, nil
}

func CreateFile(name string) string {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "File not created"
	}
	defer func() {
		file.Close()
	}()
	return "File created"
}

func main() {
	res, err := div(1, 0) //return panic error
	if err != nil {
		fmt.Println("result - ", res)
	}
	fmt.Println(CreateFile("Test")) // Returns result of the file creation function
}
