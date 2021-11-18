package mnist

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMnistLoader(t *testing.T) {
	Convey("Given : Mnistのデータを取得", t, func() {
		os.Mkdir("data",