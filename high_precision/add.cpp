#include <iostream>
#include <vector>

using namespace std;
//高精度加法
// C = A + B, A >= 0, B >= 0
// 注意 保存顺序是 个十百千万 位对应，小的放前面，这样要进位的话，就可以直接在末尾追加，如果反过来就要最前面插入，后续元素后移
vector<int> add(vector<int> &A, vector<int> &B) {
    if (A.size() < B.size()) return add(B, A); // 让第一个入参始终是更大的 size，也就是数值更大

    vector<int> C;
    int t = 0;
    for (int i = 0; i < A.size(); i++) {
        t += A[i]; // 先加上 被加数的位
        if (i < B.size()) t += B[i]; // 如果加数位数足够的话，就加 加数的位
        C.push_back(t % 10); // 新的数组对应位是 模10 后的值
        t /= 10; // t 就是 对应的多出的数，同时进入下一位的运算中
    }

    if (t) C.push_back(t); // 算完 A 后，发现 t != 0 ，说明要进位，直接 push 到最末尾
    return C;
}
