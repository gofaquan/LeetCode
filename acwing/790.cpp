#include <algorithm>
#include <iostream>

using namespace std;

const int N = 100001;

double n;


int main() {
    cin >> n;
    double l = -10000, r = 10000;
    while (r - l > 1e-8) {
        double mid = (l + r) / 2;
        if (mid * mid * mid <= n) {
            l = mid;
        } else {
            r = mid;
        }
    }

    printf("%.6f", l);
    return 0;
}