#include <iostream>
#include <limits>
#include <string>

using namespace std;

// 不考虑溢出
// 利用整型数字和字符串的相互转换以及字符串拼接的方法

int reverse(int x) {
    if (x > 0) {
        string str = to_string(x);
        int length = str.size();
        string str1;
        for (int i = length - 1; i >= 0; i--) {
            str1 += str[i];
        }
        int ret = atoi(str1.c_str());
        return ret;
    } else {
        int y = -x;
        string str = to_string(y);
        int length = str.size();
        string str1;
        for (int i = length - 1; i >= 0; i--) {
            str1 += str[i];
        }
        int ret = atoi(str1.c_str());
        return ret;
    }
}

void test01() {
    cout << "请输入一个整数：";
    int x = 0;
    cin >> x;
    cout << "反转后为：" << reverse(x) << endl;
}

int main() {
    test01();
    return 0;
}