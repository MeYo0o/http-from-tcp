package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		//* defer closing the file and channel when done
		defer f.Close()
		defer close(ch)

		//* initializing the buffer to read 8 bytes at a time
		buffer := make([]byte, 8)

		//* string holder: will hold the current incomplete line being built
		var currentLine string

		//* infinitely checking
		for {
			//* reading according to the buffer's size
			bytesRead, err := f.Read(buffer)
			if err != nil {
				//* end of file error ? => send any remaining line and exit
				if err == io.EOF {
					if currentLine != "" {
						ch <- currentLine
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
					//* send the current line + this part (which completes a line)
					currentLine += parts[i]
					ch <- currentLine
					//* reset current line for the next sentence
					currentLine = ""
				}

				//* add the last part to current line (it doesn't end with \n, so it's incomplete)
				currentLine += parts[len(parts)-1]
			}
		}
	}()

	return ch
}

func main() {
	//* set up a TCP listener on port :42069
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("couldn't create listener:", err)
	}

	//* close the listener when the program exits
	defer listener.Close()

	//* infinite loop to accept connections
	for {
		//* wait for and accept a connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("couldn't accept connection:", err)
		}

		//* print that connection was accepted
		fmt.Println("connection accepted")

		//* get the lines channel from our reusable function
		lines := getLinesChannel(conn)

		//* range over the channel and print each line
		for line := range lines {
			fmt.Println(line)
		}

		//* print that the connection has been closed
		fmt.Println("connection closed")
	}
}
