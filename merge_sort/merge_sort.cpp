#include<iostream>

using namespace std;
int n, q[100], tmp[100];

void merge_sort(int q[], int l, int r) {
    if (l >= r)return; // left >= right 直接退出
    int mid = (l + r) >> 1; // (left + right) /2
    merge_sort(q, l, mid), merge_sort(q, mid + 1, r); // 分成两块,分别排序
    // 归并，tmp 暂时存放结果
    int k = 0, i = l, j = mid + 1; // i 指向左半边数组的起点，j指向右半边的起点
    while (i <= mid && j <= r) // 当左右两半边都没循环完时
        if (q[i] <= q[j]) tmp[k++] = q[i++]; // 把较小者放入tmp
        else tmp[k++] = q[j++];
    // 两边如果不等长就会出现有一边还未循环完，直接接到 tmp 中
    while (i <= mid) tmp[k++] = q[i++];
    while (j <= r) tmp[k++] = q[j++];
    // tmp 写入 q 完成排序
    for (i = l, j = 0; i <= r; i++, j++) q[i] = tmp[j];
}


int main() {

    scanf("%d", &n);

    for (int i = 0; i < n; ++i) {
        scanf("%d", &q[i]);
    }

    merge_sort(q, 0, n - 1); // left 第一个元素，right 最后一个元素

    for (int i = 0; i < n; ++i) {
        printf("%d ", q[i]);
    }
    return 0;
}