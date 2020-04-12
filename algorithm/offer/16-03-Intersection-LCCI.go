package offer

import "math"

/**
给定两条线段（表示为起点start = {X1, Y1}和终点end = {X2, Y2}），如果它们有交点，请计算其交点，没有交点则返回空值。

要求浮点型误差不超过10^-6。若有多个交点（线段重叠）则返回 X 值最小的点，X 坐标相同则返回 Y 值最小的点。



示例 1：

输入：
line1 = {0, 0}, {1, 0}
line2 = {1, 1}, {0, -1}
输出： {0.5, 0}
示例 2：

输入：
line1 = {0, 0}, {3, 3}
line2 = {1, 1}, {2, 2}
输出： {1, 1}
示例 3：

输入：
line1 = {0, 0}, {1, 1}
line2 = {1, 0}, {2, 1}
输出： {}，两条线段没有交点


提示：

坐标绝对值不会超过 2^7
输入的坐标均是有效的二维坐标

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	x1, x2, y1, y2 := float64(start1[0]), float64(end1[0]), float64(start1[1]), float64(end1[1])
	x3, x4, y3, y4 := float64(start2[0]), float64(end2[0]), float64(start2[1]), float64(end2[1])

	if (y2-y1)*(x4-x3) == (y4-y3)*(x2-x1) {
		if (y3-y1)*(x2-x1) == (x3-x1)*(y2-y1) {
			if x1 == x2 {
				minY1 := math.Min(y1, y2)
				maxY1 := math.Max(y1, y2)
				minY2 := math.Min(y3, y4)
				maxY2 := math.Max(y3, y4)
				if minY2 <= minY1 && minY1 <= maxY2 {
					return []float64{x1, minY1}
				}
				if minY1 <= minY2 && minY2 <= maxY1 {
					return []float64{x1, minY2}
				}
				return nil
			} else {
				minX1 := math.Min(x1, x2)
				maxX1 := math.Max(x1, x2)
				minX2 := math.Min(x3, x4)
				maxX2 := math.Max(x3, x4)
				k := (y2 - y1) / (x2 - x1)
				b := y1 - k*x1
				if minX2 <= minX1 && minX1 <= maxX2 {
					return []float64{minX1, k*minX1 + b}
				}
				if minX1 <= minX2 && minX2 <= maxX1 {
					return []float64{minX2, k*minX2 + b}
				}
				return nil
			}
		}
		return nil
	} else {
		d := (y2-y1)*(x3-x4) + (x2-x1)*(y4-y3)
		t1 := ((y1-y3)*(x4-x3) + (x3-x1)*(y4-y3)) / d
		t2 := ((y2-y1)*(x3-x1) + (x1-x2)*(y3-y1)) / d
		if t1 < 0 || t1 > 1.0 || t2 < 0 || t2 > 1.0 {
			return nil
		}
		return []float64{x1 + t1*(x2-x1), y1 + t1*(y2-y1)}
	}
}
