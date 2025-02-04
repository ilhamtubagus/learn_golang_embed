package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
)

// When we build this program,
// it will embed the version.txt
// and sample-image.jpg files into the binary.
// Try removing the image file and see how the program behaves.

//go:embed version.txt
var version string

//go:embed sample-image.jpg
var sampleImage []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("sample-image-new.jpg", sampleImage, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
