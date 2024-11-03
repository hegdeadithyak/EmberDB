package main

import (
	"fmt"
	"os"
)

// The  disadvantages in this code, is what if the file is off in between the write operation what happens the write syscall status.
// No concurrency
// entriley erases the filee if found.
func savedata1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer fp.Close()
	_, err = fp.Write(data)

	return err
}

func savedata2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}

func randomInt() int {
	return 42 // Replace with actual random number generation logic if needed
}

func main() {
	path := "sample.txt"
	data := []byte("Adithya hegde is learning golanguage for being a pro in it")
	err := savedata2(path, data)

	if err != nil {
		fmt.Println("Error in saving data.")
	} else {
		fmt.Println("Data saved Sucessfully.")

	}
}
