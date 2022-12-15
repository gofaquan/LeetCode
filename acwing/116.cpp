#include <iostream>
#include <vector>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;
const int N = 4;
char g[N][N], backup[N][N];
typedef pair<int, int> PII;

void change(int x, int y) {
    for (int i = 0; i < 4; ++i) {
        g[i][y] == '+' ? g[i][y] = '-' : g[i][y] = '+';
        g[x][i] == '+' ? g[x][i] = '-' : g[x][i] = '+';
    }

    g[x][y] == '+' ? g[x][y] = '-' : g[x][y] = '+';
}


int main() {
    vector<PII> res;
    for (int i = 0; i < 4; ++i) {
        cin >> g[i];
    }


    for (int op = 0; op < 1 << 16; ++op) {
        memcpy(backup, g, sizeof g);
        vector<PII> temp;


        for (int i = 0; i < 4; ++i) {
            for (int j = 0; j < 4; ++j) {
                if (op >> (i * 4 + j) & 1) {
                    temp.push_back({i, j});
                    change(i, j);
                }
            }
        }

        bool closed = false;
        for (int i = 0; i < 4; ++i) {
            for (int j = 0; j < 4; ++j) {
                if (g[i][j] == '+') {
                    closed = true;
                }
            }
        }

        if (!closed) {
            if (res.empty() || res.size() > temp.size()) res = temp;
        }
        memcpy(g, backup, sizeof g);
    }

    cout << res.size() << endl;
    for (int i = 0; i < res.size(); ++i) {
        cout << res[i].first + 1 << ' ' << res[i].second + 1 << endl;
    }
}