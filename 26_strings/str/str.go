package str

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//ReadOneFile - reads the contents of the passed file
func ReadOneFile(name string) (s string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer CloseFile(file)

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Can't get stat from file:", err)
		return
	}

	if fileInfo.Size() == 0 {
		fmt.Printf("File %s is empty\n", name)
		return
	}

	res, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Can't read file:", err)
		return
	}
	s = fmt.Sprintf("%s\n", res)
	return
}

//ReadTwoFiles - reads the contents of 2 passed files and join it as string
func ReadTwoFiles(name1, name2 string) (res string) {
	var s []string
	s = append(s, ReadOneFile(name1))
	s = append(s, ReadOneFile(name2))
	res = strings.Join(s, "")
	return
}

//ConcatTwoFiles - reads the contents of 2 passed files, join it as string ang write to third passed file
func ConcatTwoFiles(name1, name2, name3 string) {
	res := ReadTwoFiles(name1, name2)
	file, err := os.OpenFile(name3, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Can't open file:", err)
		return
	}
	defer CloseFile(file)

	writer := bufio.NewWriter(file)
	writer.WriteString(res)
	writer.Flush()
}

//CloseFile - closes passed file
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
