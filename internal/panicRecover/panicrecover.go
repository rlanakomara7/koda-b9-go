package panicrecover

import (
	"fmt"
	"io"
	"os"
)

// PANIC & RECOVER MINITASK 6

func ProcessFile(filepath string) {

	// recover
	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Println("Panic ditangani:", panicValue)
			fmt.Println("Continue..")
		}
	}()

	//buka file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()

	//baca file
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Isi file:")
	fmt.Println(string(content))

}
