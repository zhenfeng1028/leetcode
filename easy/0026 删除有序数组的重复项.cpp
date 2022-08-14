/*
    给你一个升序排列的数组 nums，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
    元素的相对顺序应该保持一致。

    由于在某些语言中不能改变数组的长度，所以必须将结果放在数组 nums 的第一部分。
    更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
 */

/*
    方法一：双指针
    如果数组 nums 的长度为 0，则数组不包含任何元素，因此返回 0。

    当数组 nums 的长度大于 0 时，数组中至少包含一个元素，
    在删除重复元素之后也至少剩下一个元素，因此 nums[0] 保持原状即可，从下标 1 开始删除重复元素。

    定义两个指针 fast 和 slow 分别为快指针和慢指针，快指针表示遍历数组到达的下标位置，
    慢指针表示下一个不同元素要填入的下标位置，初始时两个指针都指向下标 1。

    假设数组 nums 的长度为 n。将快指针 fast 依次遍历从 1 到 n-1 的每个位置，
    对于每个位置，如果 nums[fast] ≠ nums[fast-1]， 说明 nums[fast] 和之前的元素都不同，
    因此将 nums[fast] 的值复制到 nums[slow]，然后将 slow 的值加 1，即指向下一个位置。

    遍历结束之后，从 nums[0] 到 nums[slow-1] 的每个元素都不相同且包含原数组中的每个不同的元素，
    因此新的长度即为 slow，返回 slow 即可。
 */

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int n = nums.size();
        if (n == 0) {
            return 0;
        }
        int fast = 1, slow = 1;
        while (fast < n) {
            if (nums[fast] != nums[fast - 1]) {
                nums[slow] = nums[fast];
                ++slow;
            }
            ++fast;
        }
        return slow;
    }
};
