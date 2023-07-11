package _io

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func bufIoReaderStdin() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
func bufIoReaderFile() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(inputFile)

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was:%s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
func osReaderFile() {
	file, err := os.ReadFile("input.dat")
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "File Error:%s\n", err)
		if err != nil {
			return
		}
	}
	fmt.Printf("%s\n", string(file))
	err = os.WriteFile("output.dat", file, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func fscanlnFile() {
	file, err := os.Open("input.dat")
	if err != nil {
		panic(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err.Error())
		}
	}(file)

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

type Book struct {
	title  string
	price  float64
	number int
}

func readCSV() {
	file, err := os.Open("csv.txt")
	if err != nil {
		panic(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err.Error())
		}
	}(file)

	reader := bufio.NewReader(file)
	books := make([]Book, 1)
	for {
		line, readErr := reader.ReadString('\n')
		if readErr == io.EOF {
			fmt.Println("end done")
			break
		}
		if readErr != nil {
			panic(readErr.Error())
		}

		fmt.Println("line:", line, len(line))
		if line == " " {
			fmt.Println("")
		}

		line = line[:len(line)-2]

		split := strings.Split(line, ";")

		book := new(Book)
		book.title = split[0]
		book.price, err = strconv.ParseFloat(split[1], 32)
		book.price = math.Round(book.price*100) / 100
		if err != nil {
			panic(err.Error())
		}
		book.number, err = strconv.Atoi(split[2])
		if err != nil {
			panic(err.Error())
		}
		if books[0].title == "" {
			books[0] = *book
		} else {
			books = append(books, *book)
		}

	}
	fmt.Println(books)

}

func readZipFile(zipPath string) error {
	zipFile, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	for _, file := range zipFile.File {
		fmt.Println("Reading file:", file.Name)
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		scanner := bufio.NewScanner(rc)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}
func writeFile() {
	outputFile, outputErr := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputErr != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(outputFile)

	writer := bufio.NewWriter(outputFile)
	outputString := "Hello world!\n"
	for i := 0; i < 10; i++ {
		writeString, err := writer.WriteString(outputString)
		if err != nil {
			fmt.Println(writeString)
			return
		}
	}
	err := writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
}
