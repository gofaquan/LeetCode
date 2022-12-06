#include <iostream>
#include <vector>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

struct Interval {
    int left, right;
    Interval(int l, int r) : left(l), right(r) {}
};


bool cmp(Interval a, Interval b) {
    return a.left < b.left;
}

int main() {
    int n;
    cin >> n;
    int l, r;
    vector<Interval> interval;
    for (int i = 0; i < n; ++i) {
        cin >> l >> r;
        interval.push_back(Interval(l, r));
    }
    sort(interval.begin(), interval.end(), cmp);

    int cnt = 1;
    int right = interval[0].right;
    for (int i = 1; i < n; ++i) {
        if (interval[i].left <= right) {
            right = max(interval[i].right, right);
        } else {
            cnt++;
            right = interval[i].right;
        }
    }

    cout << cnt << endl;
    return 0;
}