package math

import "math"

/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sqrtx
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func mySqrt(x int) int {
	//if x == 0 {
	//	return 0
	//}
	//temp := float64(x)
	//var e1, e2 float64 = temp, 0.0
	//for {
	//	e2 = (e1 + temp/e1) * 0.5
	//	if math.Abs(e1-e2) < 1e-7 {
	//		break
	//	}
	//	e1 = e2
	//}
	//return int(e1)

	res := int(math.Exp(0.5*math.Log(float64(x)))) + 1
	if res*res <= x {
		return res
	}
	return res - 1
}
