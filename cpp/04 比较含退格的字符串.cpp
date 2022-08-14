#include <iostream>
#include <stack>

using namespace std;

// 利用两个栈，遍历两个字符串，遇到#就出栈，否则就入栈
// 遍历完后将栈转换为字符串，对比两个转换完后的字符串

// 代码一
/*
bool backspaceCompare(string S, string T) {
    stack<char> st1, st2;
    for (int i = 0; i < S.size(); i++) {
        if (S[i] != '#') {
            st1.push(S[i]);
        } else if (!st1.empty()) {
            st1.pop();
        }
    }
    for (int i = 0; i < T.size(); i++) {
        if (T[i] != '#') {
            st2.push(T[i]);
        } else if (!st2.empty()) {
            st2.pop();
        }
    }
    if (st1.size() != st2.size()) {
        return false;
    } else {
        string str1 = "";
        string str2 = "";
        while (!st1.empty()) {
            str1 += st1.top();
            st1.pop();
        }
        while (!st2.empty()) {
            str2 += st2.top();
            st2.pop();
        }
        if (str1 == str2) {
            return true;
        } else {
            return false;
        }
    }
}
 */

// 代码二（简洁代码）
string build(string str);
bool backspaceCompare(string S, string T) { return build(S) == build(T); }

string build(string str) {
    string ret;
    for (char ch : str) {
        if (ch != '#') {
            ret.push_back(ch);
        } else if (!ret.empty()) {
            ret.pop_back();
        }
    }
    return ret;
}

int main() {
    string S = "abcd";
    string T = "bbcd";
    bool ret = backspaceCompare(S, T);
    cout << ret << endl;
    system("pause");
    return 0;
}