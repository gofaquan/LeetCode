#include<iostream>

using namespace std;

void quick_sort(int q[], int l, int r) {
    if (l >= r)return; // left >= right 直接退出

    int x = q[l], i = l-1, j = r + 1; // x = queue[left], i = left -1, j = right + 1
    while (i < j) { // i < j 时
        do i++; while (q[i] < x); // 左指针右移，直到当前指向数不再小于x
        do j--; while (q[j] > x); // 右指针左移，直到当前指向数不再大于x
        if (i < j) swap(q[i], q[j]); // 如果此时满足 i<j ，就 swap 交换数组对应下标值
    }
    quick_sort(q, l, j); // 递归调用，新的 right 为 j
    quick_sort(q,j+1,r); // 递归调用，新的 left 为 j+1
}


int main() {
    int n, q[100];
    scanf("%d", &n);

    for (int i = 0; i < n; ++i) {
        scanf("%d", &q[i]);
    }

    quick_sort(q,0, n-1); // left 第一个元素，right 最后一个元素

    for (int i = 0; i < n; ++i) {
        printf("%d ",q[i]);
    }
    return 0;
}