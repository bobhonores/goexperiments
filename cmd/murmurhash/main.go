// Using Murmur Hash version 3
// https://github.com/aappleby/smhasher/wiki/MurmurHash3

package main

import (
	"fmt"
	"strconv"

	murmur3 "github.com/yihleego/murmurhash3"
)

func main() {
	murmur := murmur3.New128()
	r := murmur.HashString("A" + "property:sample")
	fmt.Printf("%s hash as string\n", r.String())
	fmt.Printf("%s hash as hex\n", r.AsHex())
	fmt.Printf("%s hash as int32\n", strconv.Itoa(r.AsInt()))
	fmt.Printf("%s hash as int64\n", strconv.FormatInt(r.AsInt64(), 10))
}
