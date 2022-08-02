#include <iostream>
#include <vector>

using namespace std;
// C = A - B, 满足A >= B, A >= 0, B >= 0
vector<int> sub(vector<int> &A, vector<int> &B)
{
    vector<int> C;
    for (int i = 0, t = 0; i < A.size(); i ++ ) // 在大的数里面遍历
    {
        t = A[i] - t; // 先计算被减数的位的值
        if (i < B.size()) t -= B[i]; // 位数足够的话，就要再减去减数位的值
        C.push_back((t + 10) % 10); // push 进 模10后的数，注意先 +10 是如果 t 为负数的话，+10 就可以变正数，本来就是整数就不影响
        if (t < 0) t = 1; // 小于 0 的话，说明有借位，前一位需要减去 1
        else t = 0; // 大于等于0 的话，说明本来就够，没有借位情况
    }

    while (C.size() > 1 && C.back() == 0) C.pop_back(); // 如果出现了 007 这样的情况就要把 多余的 0 去掉，使得结果为 7
    return C;
}
