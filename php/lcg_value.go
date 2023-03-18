package php

import (
	"math/rand"
	"time"
)

//我们使用time包获取当前时间，并将其作为种子创建一个新的随机数生成器。然后，我们使用rand.Int31()函数生成一个介于0和(2^31-1)之间的随机整数，将其转换为浮点数，并除以(2^31-1)得到0到1之间的随机小数，就像在PHP的lcg_value函数中一样。
//需要注意的是，Go中的伪随机数生成器并不完全相同于PHP中的线性同余生成器算法，因此它们生成的随机数序列可能不完全相同。
//We use the time package to retrieve the current time and use it as a seed to create a new random number generator. Then, we use the rand. Int31() function to generate a random integer between 0 and (2 ^ 31-1), convert it to a floating point number, and divide it by (2 ^ 31-1) to obtain a random decimal between 0 and 1, just like in PHP's lcg_ The value function.
//It should be noted that the pseudo random number generators in Go are not exactly the same as the linear congruence generator algorithm in PHP, so the random number sequences they generate may not be identical.
func LcgValue() float64 {
	// 创建一个新的随机数生成器，以当前时间作为种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成一个介于0和(2^31-1)之间的随机整数
	i := r.Int31()

	// 将随机整数转换为浮点数，并除以(2^31-1)得到0到1之间的随机小数
	return float64(i) / float64(1<<31-1)
}
