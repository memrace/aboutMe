package commands

import (
	"fmt"
	"io"
	"os"
)

func getTextFromReadMe() string {
	file, err := os.Open("README.md")

	if err != nil {
		fmt.Println(err)
		return "error"
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	data := make([]byte, 64)
	var text string
	for {
		n, err := file.Read(data)

		if err == io.EOF {
			break
		}

		text += string(data[:n])
	}
	return text
}
