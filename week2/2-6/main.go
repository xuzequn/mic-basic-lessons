package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// io reader wirter
func main() {

	reader := strings.NewReader("Talk is Cheap, Show me Code")
	p := make([]byte, 8)

	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}

	meats := []string{
		"chicken",
		"|beaf",
		"|pork",
		"|mutoon",
	}
	var writer bytes.Buffer
	for _, p := range meats {
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println(writer.String())
}
