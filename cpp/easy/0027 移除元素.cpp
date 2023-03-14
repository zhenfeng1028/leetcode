/*
    给你一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，并返回移除后数组的新长度。

    元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
 */

/*
    方法：双指针
    如果要移除的元素恰好在数组的开头，例如序列 [1, 2, 3, 4, 5]，当 val 为 1 时，我们需要把每一个元素都左移一位。
    注意到题目中说：「元素的顺序可以改变」。实际上我们可以直接将最后一个元素 5 移动到序列开头，取代元素 1，
    得到序列 [5, 2, 3, 4]，同样满足题目要求。这个优化在序列中 val 元素的数量较少时非常有效。

    如果左指针 left 指向的元素等于 val，此时将右指针 right 指向的元素复制到左指针 left 的位置，
    然后右指针 right 左移一位。如果赋值过来的元素恰好也等于 val，可以继续把右指针 right 指向的元素的值赋值过来
    （左指针 left 指向的等于 val 的元素的位置继续被覆盖），直到左指针指向的元素的值不等于 val 为止。

    当左指针 left 和右指针 right 重合的时候，左右指针遍历完数组中所有的元素。
 */

class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int left = 0, right = nums.size();
        while (left < right) {
            if (nums[left] == val) {
                nums[left] = nums[right - 1];
                right--;
            } else {
                left++;
            }
        }
        return left;
    }
};
