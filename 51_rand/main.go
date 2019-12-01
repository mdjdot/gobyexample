package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100), "\n")

	fmt.Println(rand.Float64()) // Float64 returns, as a float64, a pseudo-random number in [0.0,1.0) from the default Source.

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64()*5)+5, "\n")

	s1 := rand.NewSource(time.Now().UnixNano()) // seed建议用时间
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100), "\n")

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100), "\n")

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))

}

/*
需要注意的是，如果你出于加密目的，需要使用随机数的话，请使用 crypto/rand 包，此方法不够安全。
如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列。
*/
