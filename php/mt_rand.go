package php

import (
	"math/rand"
	"time"
)

//我们使用Go的rand包中的Seed()函数将当前时间作为随机数生成器的种子，并使用Intn()函数生成随机数。
//注意，由于Intn()函数返回的随机数不包括上限，所以我们需要将max参数增加1来确保生成的随机数落在指定的范围之内。
//We use the Seed() function in Go's rand package to seed the current time as a random number generator, and use the Intn() function to generate random numbers.
//Note that since the random number returned by the Intn() function does not include the upper limit, we need to increase the max parameter by 1 to ensure that the generated random number falls within the specified range.
func MtRand(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
