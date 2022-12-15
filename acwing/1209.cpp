#include <iostream>
#include <vector>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

const int N = 9;
bool state[N];
int data1[N];
int n;
int cnt = 0;

int calc(int l, int r) {
    int res = 0;
    for (int i = l; i <= r; ++i) {
        res = res * 10 + data1[i];
    }

    return res;
}

void dfs(int u) {
    if (u == 9) {
        for (int i = 0; i < 7; ++i) {
            for (int j = i + 1; j < 8; ++j) {
                int a = calc(0, i);
                int b = calc(i+1, j);
                int c = calc(j+1, 8);
                if (a * c + b == c * n) {
                    cnt++;
                }
            }
        }
        return;
    }


    for (int i = 0; i < 9; ++i) {
        if (!state[i]) {
            state[i] = true;
            data1[u] = i + 1;
            dfs(u + 1);
            state[i] = false;
        }
    }

}

int main() {
    cin >> n;
    dfs(0);
    cout << cnt;
}