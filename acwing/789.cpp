#include <algorithm>
#include <iostream>
// leetcode 34
using namespace std;

const int N = 100001;

int n, q, x;
int a[N];

int main() {
    cin >> n >> q;
    for (int i = 0; i < n; i++) cin >> a[i];

    while (q--) {
        cin >> x;
        int l = 0, r = n - 1;
        while (l < r) {
            int mid = (l + r) >> 1;
            if (a[mid] >= x) r = mid;
            else l = mid + 1;
        }

        if (a[l] != x) cout << "-1 -1" << endl;
        else {
            cout << l << ' ';
            //找左边界，第二个模板
            l = 0, r = n - 1;
            while (l < r) {
                int mid = (l + r + 1) >> 1;
                if (a[mid] <= x) l = mid;
                else r = mid - 1;
            }
            cout << l << endl;
        }
    }

    return 0;
}