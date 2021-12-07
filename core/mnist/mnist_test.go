package mnist

import (
	"fmt"
	"image/color"
	"os"
	"testing"

	"github.com/petar/GoMNIST"
)

func TestMnistLoad(t *testing.T) {
	os.Mkdir("data", 0777)
	train, _, err := GoMNIST.Load("data")
	if err != nil {
		t.Fatalf("can't load mnist data! error detail i