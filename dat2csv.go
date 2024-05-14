package main

import (
	"os"
)

func readInput() []byte {
	data := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			break
		}
		data = append(data, buffer[:n]...)
	}
	return data
}

func decodeData(data []byte) []byte {
	decodedData := make([]byte, len(data))
	for i, char := range data {
		decodedData[i] = char ^ 0xFF
	}
	return decodedData
}

func main() {
	data := readInput()
	decodedData := decodeData(data)
	os.Stdout.Write(decodedData)
}
