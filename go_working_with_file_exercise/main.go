package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Exercise #1
	newFile, err := os.Create("info.txt")

	if err != nil {
		log.Fatal(err)
	}

	newFile.Close()

	// Exercise #2

	oldFileName := "info.txt"
	newFileName := "data.txt"
	_, err = os.Stat(oldFileName)

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}

	err = os.Rename(oldFileName, newFileName)

	if err != nil {
		log.Fatal(err)
	}

	// Exercise #3
	err = os.Remove(newFileName)
	if err != nil {
		log.Fatal(err)
	}

	// Exercise #4
	file, err := os.Open("read.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", data)

	// Exercise #5
	file, err = os.Open("scanned.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Exercise #6
	file, err = os.OpenFile("info.txt", os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	byteSlice := []byte("The Go gopher is an iconic mascot!")
	bytesWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written: %d\n", bytesWritten)

}
