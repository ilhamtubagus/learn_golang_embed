package learn_golang_embed

import (
	_ "embed"
	"fmt"
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
