package _io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestBufIoReader(t *testing.T) {
	bufIoReaderStdin()
}
func TestBufIoReaderFile(t *testing.T) {
	bufIoReaderFile()
}
func TestOsReaderFile(t *testing.T) {
	osReaderFile()
}
func TestFscanlnFile(t *testing.T) {
	fscanlnFile()
}
func TestReadCSV(t *testing.T) {
	readCSV()
}
func TestReadZipFile(t *testing.T) {
	err := readZipFile("csv.txt.zip")
	fmt.Println(err)
}
func TestWriteFile(t *testing.T) {
	writeFile()
}
func TestFlushFile(t *testing.T) {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Conversion done")
}
