package binary2

import (
	"fmt"
	"go-structure/other"
	"go-structure/subone"
)

func main() {
	otherRes := other.OtherOne()
	fmt.Println(otherRes)
	variousRes := subone.SubOne()
	fmt.Println(variousRes)
}
