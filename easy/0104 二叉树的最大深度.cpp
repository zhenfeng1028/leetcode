/*
    给定一个二叉树，找出其最大深度。

    二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 */

/*
    方法：深度优先搜索
    如果我们知道了左子树和右子树的最大深度 l 和 r，那么该二叉树的最大深度即为

                                max(l,r)+1

    而左子树和右子树的最大深度又可以以同样的方式进行计算。
 */
class Solution {
public:
    int maxDepth(TreeNode* root) {
        if (root == nullptr) return 0;
        return max(maxDepth(root->left), maxDepth(root->right)) + 1;
    }
};
