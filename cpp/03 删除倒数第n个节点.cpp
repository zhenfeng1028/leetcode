#include <iostream>
using namespace std;

// 定义一个链表节点
struct ListNode {
    int value;
    ListNode* next;
};

// 插入一个新节点到链表中(放在链表头部)
void CreateList(ListNode*& head, int data) {
    // 创建新节点
    ListNode* p = (ListNode*)malloc(sizeof(ListNode));
    p->value = data;
    p->next = NULL;

    if (head == NULL) {
        head = p;  // 如果传进来的是一个空表，那么让head指针指向第一个节点
        return;
    }
    p->next = head;  // 让新节点指向上一个节点（head指向的是上一个节点）
    head = p;        // 再让head重新指向新节点（头插）
}

void printList(ListNode* head) {
    ListNode* p = head;
    while (p != NULL) {
        cout << p->value << " ";
        p = p->next;
    }
    cout << endl;
}

ListNode* removeNthFromEnd(ListNode* head, int n) {
    if (head == NULL || head->next == NULL) return NULL;
    ListNode* fast = head;
    ListNode* slow = head;
    for (int i = 0; i < n; i++) {
        fast = fast->next;  // 指针后移
    }
    if (fast == NULL) return head->next;
    while (fast->next != NULL) {
        fast = fast->next;  // 指针后移
        slow = slow->next;  // 指针后移
    }
    slow->next = slow->next->next;
    return head;
}

int main() {
    ListNode* head = NULL;
    for (int i = 9; i >= 0; i--) CreateList(head, i);
    printList(head);
    head = removeNthFromEnd(head, 3);
    printList(head);
    return 0;
}