package tests

import (
	"math/rand"
	"phpfunc/php"
	"testing"
	"time"
)

// 测试 mt_rand 函数的输出结果是否在指定范围内
func TestMtRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// 生成 1000 个随机数，并检查它们是否在 [0, 100] 范围内
	for i := 0; i < 1000; i++ {
		num := php.MtRand(0, 100)
		if num < 0 || num > 100 {
			t.Errorf("TestMtRand generated number %d is not within the expected range [0, 100]", num)
		}
	}

	// 生成 1000 个随机数，并检查它们是否在 [-50, 50] 范围内
	for i := 0; i < 1000; i++ {
		num := php.MtRand(-50, 50)
		if num < -50 || num > 50 {
			t.Errorf("TestMtRand generated number %d is not within the expected range [-50, 50]", num)
		}
	}
}
