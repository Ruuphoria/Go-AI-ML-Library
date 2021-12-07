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