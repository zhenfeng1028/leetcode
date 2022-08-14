#include <iostream>
using namespace std;

// 题目描述
// 给出两个非空的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的
// 并且它们的每个节点只能存储一位数字。如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

// 定义一个链表节点
struct ListNode {
    int value;
    ListNode* next;
    ListNode(int x) : value(x), next(nullptr) {}
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

ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    ListNode* head = nullptr;
    ListNode* tail = nullptr;
    int carry = 0;
    while (l1 || l2) {
        int n1 = l1 ? l1->value : 0;  // 如果不为空，取节点中的值；如果为空，取0
        int n2 = l2 ? l2->value : 0;
        int sum = n1 + n2 + carry;
        if (!head) {
            head = tail = new ListNode(sum % 10);  // 一开始让head和tail指向新创建的节点
        } else {
            tail->next = new ListNode(sum % 10);  // 此后让tail一直指向新创建的节点
            tail = tail->next;                    // tail节点前移
        }
        carry = sum / 10;
        if (l1) l1 = l1->next;
        if (l2) l2 = l2->next;
    }
    if (carry > 0) {
        tail->next = new ListNode(carry);
    }
    return head;
}

int main() {
    ListNode* l1 = NULL;
    CreateList(l1, 2);
    CreateList(l1, 4);
    CreateList(l1, 3);

    ListNode* l2 = NULL;
    CreateList(l2, 5);
    CreateList(l2, 6);
    CreateList(l2, 4);

    ListNode* head = addTwoNumbers(l1, l2);
    printList(head);
    return 0;
}