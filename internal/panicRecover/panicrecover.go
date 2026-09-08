package panicrecover

import (
	"fmt"
	"io"
	"os"
)

// PANIC & RECOVER MINITASK 6

func ReadFile(filepath string) ([]byte, error) {

	//buka file
	file, err := os.Open(filepath)
	if err != nil {

		return nil, fmt.Errorf("gagal membuka file: %w", err)
	}

	// recover
	defer func() {
		file.Close()
		fmt.Println("File ditutup.")
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("error membaca file : %v", err))

	}
	return content, nil
}
