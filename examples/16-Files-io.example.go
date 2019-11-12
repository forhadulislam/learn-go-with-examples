package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Write to a file using ioutil.WriteFile
	d1 := []byte("This is a simple text. \n Do not take it seriously. \n")
	err := ioutil.WriteFile("/tmp/output1", d1, 0644)
	fmt.Errorf("Error: %v",err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("/tmp/output2")
	fmt.Errorf("Error: %v",err)
	// Defer a Close immediately after opening a file.
	defer f.Close()

	// You can Write byte slices as you’d expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	fmt.Errorf("Error: %v",err)
	fmt.Printf("wrote %d bytes\n", n2)

	//A WriteString is also available.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied to the underlying writer.
	w.Flush()
}
