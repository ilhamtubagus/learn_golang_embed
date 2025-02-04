package learn_golang_embed

import (
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
