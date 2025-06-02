/*
   罗马数字包含以下七种字符: I，V，X，L，C，D 和 M。

   字符          数值
   I             1
   V             5
   X             10
   L             50
   C             100
   D             500
   M             1000
   例如， 罗马数字 2 写做 II，即为两个并列的 1。12 写做 XII，即为 X + II。 27 写做 XXVII，即为 XX + V + II。

   通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5
   的左边，所表示的数等于大数 5 减小数 1 得到的数值 4。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

   I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
   X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
   C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
   给定一个罗马数字，将其转换成整数。
*/

package easy

func romanToInt(s string) int {
	hashmap := make(map[string]int)
	hashmap["I"] = 1
	hashmap["V"] = 5
	hashmap["X"] = 10
	hashmap["L"] = 50
	hashmap["C"] = 100
	hashmap["D"] = 500
	hashmap["M"] = 1000

	value := 0
	if len(s) == 1 {
		return hashmap[s]
	}
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && hashmap[string(s[i])] < hashmap[string(s[i+1])] {
			value -= hashmap[string(s[i])]
			value += hashmap[string(s[i+1])]
			i++
		} else {
			value += hashmap[string(s[i])]
		}
	}
	return value
}
