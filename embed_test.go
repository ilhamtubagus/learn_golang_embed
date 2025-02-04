package learn_golang_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string // must be placed globally outside function

//go:embed version.txt
var version2 string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed sample-image.jpg
var img []byte

func TestByteArray(t *testing.T) {
	err := os.WriteFile("sample-image-new.jpg", img, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/*
var files embed.FS

// or you can specify a specific file to embed
// //go:embed files/a.txt

func TestEmbedMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}
