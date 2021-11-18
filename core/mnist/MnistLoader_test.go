package mnist

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMnistLoader(t *testing.T) {
	Convey("Given : Mnistのデータを取得", t, func() {
		os.Mkdir("data", 0777)
		train, test, err := LoadData("data")

		Convey("Then : エラーが発生しないこと", func() {
			So(err, ShouldBeNil)
		})

		Convey(