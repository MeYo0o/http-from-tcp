package main

import (
	"fmt"
	"io"
	"log"
	"os"
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

	//* infinitely checking
	for {
		//* reading according to the buffer's size
		bytesRead, err := file.Read(buffer)
		if err != nil {
			//* end of file error ? => exit the function
			if err == io.EOF {
				return
			}

			log.Fatal("couldn't read the file:", err)
		}

		//* if we read at least 1 byte, convert it's byte value to string and print it.
		//* notice that we convert only converted bytes to protect our selves from garbled/extra characters in case we read bytes < 8
		if bytesRead > 0 {
			fmt.Printf("read: %s\n", string(buffer[:bytesRead]))
		}
	}

}
