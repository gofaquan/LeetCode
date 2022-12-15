#include <iostream>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;
const int N = 110;
char g[N], aim[N];

void change(int x) {
    g[x] == '*' ? g[x] = 'o' : g[x] = '*';
}


int main() {
    cin >> g >> aim;

    int cnt = 0;
    for (int i = 0; i < strlen(g); ++i) {
        if (g[i] != aim[i]){
            change(i);
            change(i + 1);
            cnt++;
        }
    }

    cout << cnt << endl;

}