/*
    给你两个二进制字符串，返回它们的和（用二进制表示）。

    输入为非空字符串且只包含数字 1 和 0。
*/

/*
    我们可以借鉴「列竖式」的方法，末尾对齐，逐位相加。在十进制的计算中「逢十进一」，二进制中我们需要「逢二进一」。

    具体的，我们可以取 n=max{|a|,|b|}，循环 n 次，从最低位开始遍历。我们使用一个变量 carry 表示上一个位置的进位，初始值为 0。
    记当前位置对其的两个位为 a(i) 和 b(i)，则每一位的答案为 (carry+a(i)+b(i)) mod 2，下一位的进位为 (carry+a(i)+b(i)) / 2。
    重复上述步骤，直到字符串 a 和 b 的每一位计算完毕。最后如果 carry 的最高位不为 0，则将最高位添加到计算结果的末尾。
*/

class Solution {
public:
    string addBinary(string a, string b) {
        string ans;
        reverse(a.begin(), a.end());
        reverse(b.begin(), b.end());

        int n = max(a.size(), b.size()), carry = 0;
        for (size_t i = 0; i < n; ++i) {
            if (i < a.size()) {
                carry += int(a.at(i) - '0');
            }
            if (i < b.size()) {
                carry += int(b.at(i) - '0');
            }
            ans.push_back((carry % 2) ? '1' : '0');
            carry /= 2;
        }

        if (carry) {
            ans.push_back('1');
        }
        reverse(ans.begin(), ans.end());

        return ans;
    }
};
