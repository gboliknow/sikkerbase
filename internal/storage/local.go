package storage

import (
	"fmt"
	"os"
)

func SaveToLocal(path, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	fmt.Println("File saved successfully:", path)
	return nil
}
