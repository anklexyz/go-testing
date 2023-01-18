package benchmark

import (
	"github.com/NebulousLabs/fastrand"
	"math/rand"
)

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

func Select() int {
	//Intn作为int返回一个非负的伪随机数，
	//该伪随机数位于默认Source的半开放区间[0,n)。如果n <= 0，它会出现恐慌。
	//未设定种子数而产生的随机数是固定数
	return ServerIndex[rand.Intn(10)]
}

func FastSelect() int {
	return ServerIndex[fastrand.Intn(10)]
}
