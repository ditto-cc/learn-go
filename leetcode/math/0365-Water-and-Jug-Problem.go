package math

/**
You are given two jugs with capacities x and y litres. There is an infinite amount of water supply available. You need to determine whether it is possible to measure exactly z litres using these two jugs.

If z liters of water is measurable, you must have z liters of water contained within one or both buckets by the end.

Operations allowed:

Fill any of the jugs completely with water.
Empty any of the jugs.
Pour water from one jug into another till the other jug is completely full or the first jug itself is empty.
Example 1: (From the famous "Die Hard" example)

Input: x = 3, y = 5, z = 4
Output: True
Example 2:

Input: x = 2, y = 6, z = 5
Output: False
*/

func gcd(a, b int) int {
	if a == b {
		return a
	}
	r1, r2 := a&1, b&1
	switch {
	case r1 == 0 && r2 == 0:
		return gcd(a>>1, b>>1) * 2
	case r1 == 0 && r2 == 1:
		return gcd(a>>1, b)
	case r1 == 1 && r2 == 0:
		return gcd(a, b>>1)
	default:
		if a > b {
			return gcd(a-b, b)
		} else {
			return gcd(b-a, a)
		}
	}
}

func canMeasureWater(x int, y int, z int) bool {
	if z <= 0 {
		return true
	}

	if x == 0 {
		return z == y
	}

	if y == 0 {
		return z == x
	}

	d := gcd(x, y)
	if z <= x+y && d != 0 && z%d == 0 {
		return true
	}
	return false
}
