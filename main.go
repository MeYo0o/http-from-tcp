package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const fileName string = "messages.txt"

func main() {
	//* open the file.
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("couldn't open the file:", err)
	}

	//* at the end of the function, close file the reading process.
	defer file.Close()

	//* initializing the buffer to read 8 bytes at a time
	buffer := make([]byte, 8)

	//* string holder: will hold the current incomplete line being built
	var currentLine string

	//* infinitely checking
	for {
		//* reading according to the buffer's size
		bytesRead, err := file.Read(buffer)
		if err != nil {
			//* end of file error ? => exit the function and print any remaining line
			if err == io.EOF {
				if currentLine != "" {
					fmt.Printf("read: %s\n", currentLine)
				}
				return
			}

			log.Fatal("couldn't read the file:", err)
		}

		//* if we read at least 1 byte, process it
		if bytesRead > 0 {
			//* convert read bytes to string
			data := string(buffer[:bytesRead])

			//* split on newlines to find complete lines
			parts := strings.Split(data, "\n")

			//* process all parts except the last one (they represent complete lines)
			for i := 0; i < len(parts)-1; i++ {
				//* print the current line + this part (which completes a line)
				currentLine += parts[i]
				fmt.Printf("read: %s\n", currentLine)
				//* reset current line for the next sentence
				currentLine = ""
			}

			//* add the last part to current line (it doesn't end with \n, so it's incomplete)
			currentLine += parts[len(parts)-1]
		}
	}

}
