#include <algorithm>
#include <iostream>

using namespace std;

const int N = 100010;

int n, max_v = 1;
int h[N];

bool check(int e) {
    for (int i = 1; i <= n; i++) {
        e = e * 2 - h[i];
        if (e >= 1e5) return true;
        if (e < 0) return false;
    }
    return true;
}

int main() {
    scanf("%d", &n);
    for (int i = 1; i <= n; i++) {
        scanf("%d", &h[i]);
        if (h[i] > max_v) {
            max_v = h[i];
        }
    }


    int l = 0, r = max_v;
    while (l < r) {
        int mid = l + r >> 1;
        if (check(mid)) r = mid;
        else l = mid + 1;
    }

    printf("%d\n", r);

    return 0;
}