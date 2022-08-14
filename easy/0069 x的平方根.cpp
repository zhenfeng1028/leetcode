/*
    给你一个非负整数 x ，计算并返回 x 的算术平方根 。

    由于返回类型是整数，结果只保留整数部分，小数部分将被舍去 。

    注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
 */

/*
    方法：二分法
    由于 x 平方根的整数部分是满足 k^2 ≤ x 的最大 k 值，因此我们可以对 k 进行二分查找，从而得到答案。
 */

class Solution {
public:
    int mySqrt(int x) {
        int l = 0, r = x, ans = -1;
        while (l <= r) {
            int mid = l + (r - l) / 2;
            if ((long long)mid * mid <= x) {
                ans = mid;
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }
        return ans;
    }
};
