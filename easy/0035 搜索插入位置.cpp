/*
    给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

    请必须使用时间复杂度为 O(log n) 的算法。
 */

/*
    方法：二分查找
    考虑这个插入的位置 pos，它成立的条件为：nums[pos−1] < target ≤ nums[pos]
    其中 nums 代表排序数组。由于如果存在这个目标值，我们返回的索引也是 pos，
    因此我们可以将两个条件合并得出最后的目标：「在一个有序数组中找第一个大于等于 target 的下标」。
 */

class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int n = nums.size();
        int l = 0, r = n - 1;
        while (l <= r) {
            int mid = l + (r - l) / 2;  // 不使用 (l+r)/2 的原因是防止溢出
            if (nums[mid] < target)
                l = mid + 1;
            else
                r = mid - 1;
        }
        return l;
    }
};

// [ 1, 3, 5, 7, 9 ] size = 5 val = 2
// l = 0, r = 4, mid = 0 + (4 - 0) / 2 = 2, nums[mid] = nums[2] = 5 > 2, r = mid - 1 = 1
// l = 0, r = 1, mid = 0 + (1 - 0) / 2 = 0, nums[mid] = nums[0] = 1 < 2, l = mid + 1 = 1
// l = 1, r = 1, mid = 1 + (1 - 1) / 2 = 1, nums[mid] = nums[1] = 3 > 2, r = mid - 1 = 0
// l = 1, r = 0, break while, return l