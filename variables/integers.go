package variables

import "fmt"

func ShowIntegers() {
	var intCommon int
	int32Common := int32(10)
	int64Common := int64(100)
	fmt.Println("Int common = ", intCommon)
	fmt.Println("Int 32 = ", int32Common)
	fmt.Println("Int 64 = ", int64Common)
}
