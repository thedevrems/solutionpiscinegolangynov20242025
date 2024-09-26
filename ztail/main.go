package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	var i int
	var positif bool = true
	var countsignefirst int = 0

	for positioncaractere, c := range s {
		positioncaractere++
		if (!(c >= '0' && c <= '9')) && !(c == '+' || c == '-') || countsignefirst == 2 || (positioncaractere > 3 && (c == '+' || c == '-')) {
			return 0
		}
		if c == '+' || c == '-' {
			countsignefirst++
			if countsignefirst > 1 {
				return 0
			}
			if c == '-' {
				positif = false
			}
			continue
		}
		i = i*10 + int(c-'0')
	}

	if positif {
		return i
	} else {
		return -i
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . -c <num> <file1> [<file2>...]")
		os.Exit(1)
	}
	if os.Args[1] != "-c" {
		fmt.Println("Error: Missing -c option")
		os.Exit(1)
	}
	n, err := Atoi(os.Args[2])
	if err != nil || n <= 0 {
		fmt.Println("Error: Invalid number of bytes")
		os.Exit(1)
	}
	exitCode := 0
	for i, fileName := range os.Args[3:] {
		if len(os.Args[3:]) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", fileName)
		}
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("open %s: %v\n", fileName, err)
			exitCode = 1
			continue
		}
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Printf("stat %s: %v\n", fileName, err)
			file.Close()
			exitCode = 1
			continue
		}
		fileSize := fileInfo.Size()
		startPos := fileSize - int64(n)
		if startPos < 0 {
			startPos = 0
		}
		_, err = file.Seek(startPos, 0)
		if err != nil {
			fmt.Printf("seek %s: %v\n", fileName, err)
			file.Close()
			exitCode = 1
			continue
		}
		buffer := make([]byte, fileSize-startPos)
		_, err = file.Read(buffer)
		if err != nil {
			fmt.Printf("read %s: %v\n", fileName, err)
			file.Close()
			exitCode = 1
			continue
		}
		fmt.Print(string(buffer))
		file.Close()
	}
	os.Exit(exitCode)
}
