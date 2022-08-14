/*
    给定一个已排序的链表的头 head，删除所有重复的元素，使每个元素只出现一次。返回已排序的链表。
 */

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */

// 快慢指针
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if (head == nullptr) {
            return head;
        }
        ListNode *slow = head, *fast;
        if (head->next != nullptr) {
            fast = head->next;
        }
        while (fast != nullptr) {
            if (slow->val != fast->val) {
                slow->next = fast;
                slow = slow->next;
            }
            fast = fast->next;
        }
        slow->next = nullptr;
        return head;
    }
};