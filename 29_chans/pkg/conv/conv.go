package conv

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var StopWord = errors.New("stop")

//ReadInt - reads integers from input using infinite loop
func ReadInt() (int, error) {
	var n int
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		err := fmt.Errorf("invalid input")
		return n, err
	}
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, "\r")
	if strings.ToLower(str) == "стоп" || strings.ToLower(str) == "stop" {
		return n, StopWord
	}
	n, err = strconv.Atoi(str)
	if err != nil {
		err := fmt.Errorf("invalid input, try again please")
		return n, err
	}
	return n, nil
}

//PassToGoroutine - writes integers into channel
func PassToGoroutine(n int, wg *sync.WaitGroup) <-chan int {
	intChan := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		intChan <- n
		close(intChan)
		log.Printf("Goroutine that puts the number %d into channel has finished its work\n", n)
	}()
	return intChan
}

//GetSquare - calculates square(n) and translates it into the next calculating goroutine
func GetSquare(intChan <-chan int, wg *sync.WaitGroup) <-chan int {
	sqrChan := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range intChan {
			square := n * n
			fmt.Printf("Square of %d: %d\n", n, square)
			sqrChan <- square
		}
		close(sqrChan)
		time.After(2 * time.Second)
		log.Println("Squares goroutine has finished its work")
	}()
	return sqrChan
}

//GetMultiple - calculates multiplication of square(n) by 2
func GetMultiple(sqrChan <-chan int, wg *sync.WaitGroup) <-chan int {
	multChan := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for square := range sqrChan {
			mult := square * 2
			fmt.Printf("Multiplication of %d by 2: %d\n", square, mult)
			multChan <- mult
		}
		close(multChan)
		time.After(2 * time.Second)
		log.Println("Multiple goroutine has finished its work")
	}()
	return multChan
}
