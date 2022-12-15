#include <algorithm>
#include <iostream>

using namespace std;

const int N = 1000010;

int n, m;
int datas[N];
int sum[N];

// scanf 比 cin 高效挺多
// 比起 cin 和 scanf 的区别，还是for 优化空间大
int main() {
    scanf("%d %d", &n, &m);
    for (int i = 1; i <= n; ++i) {
        scanf("%d", &datas[i]);
    }

    // 比起 cin 和 scanf 的区别，还是for 优化空间大
    for (int i = 1; i <= n; ++i) {
        sum[i] = sum[i - 1] + datas[i];
    }
// 这个比上面慢几百ms
//    for (int i = 0; i < n; ++i) {
//        sum[i + 1] = sum[i] + datas[i + 1];
//    }
    for (int i = 0; i < m; ++i) {
        int l, r;
        scanf("%d %d", &l, &r);
        printf("%d\n", sum[r] - sum[l - 1]);
    }
}