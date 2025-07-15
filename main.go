package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func allZero(buffer []byte) bool {
	for _, v := range buffer {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	var blockSize int
	var output string
	var input string
	var debug bool
	var progress bool
	var inFile *os.File

	flag.IntVar(&blockSize, "block-size", 65536, "Block size")
	flag.StringVar(&input, "input", "-", "Input file name")
	flag.StringVar(&output, "output", "output", "Output file name")
	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.BoolVar(&progress, "progress", false, "Show progress")

	flag.Parse()

	file, err := os.Create(output)

	if err != nil {
		fmt.Println("Error creating file")
		os.Exit(1)
	}

	buffer := make([]byte, blockSize)

	crtSize := 0

	sparseSize := 0

	inSize := -1

	if input == "-" {
		inFile = os.Stdin
	} else {
		inFile, err = os.Open(input)

		if err != nil {
			fmt.Println("Error opening file")
			os.Exit(1)
		}

		fileInfo, err := inFile.Stat()

		if err != nil {
			fmt.Println("Error getting file info")
			os.Exit(1)
		}

		inSize = int(fileInfo.Size())
	}

	lastProgress := time.Now().Unix() - 10

	for {
		nRead, err := inFile.Read(buffer)

		if nRead == 0 || err != nil {
			break
		}

		crtSize += nRead

		isZero := allZero(buffer)

		if isZero {
			file.Seek(int64(crtSize), 0)
			sparseSize += nRead
		} else {
			file.Write(buffer[:nRead])
		}

		if debug {
			fmt.Printf("Read block %d, isZero: %t, totalSize: %d\n", nRead, isZero, crtSize)
		}

		now := time.Now().Unix()

		if progress && (lastProgress+1 < now) {
			if inSize < 0 {
				fmt.Printf("Progress: %d, sparseSize=%d(%%%f)\r", crtSize, sparseSize, float64(sparseSize)/float64(crtSize)*100.0)
			} else {
				fmt.Printf("Progress: %d/%d(%%%f), sparseSize=%d(%%%f)\r",
					crtSize, inSize,
					float64(crtSize)/float64(inSize)*100.0,
					sparseSize, float64(sparseSize)/float64(crtSize)*100.0)
			}

			lastProgress = now
		}
	}

	file.Close()
}
