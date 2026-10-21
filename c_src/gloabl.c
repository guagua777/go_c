#include <stdio.h>
#include "gloabl.h"

int g_count = 0;          // 定义并初始化
double g_value;           // 定义未初始化（默认为 0）

void increment() {
    g_count++;            // 函数内直接访问
    printf("g_count = %d\n", g_count);
}

void getCount() {
    printf("g_count = %d\n", g_count);
}

int main() {
    printf("%d\n", g_count);  // 0
    increment();
    printf("%d\n", g_count);  // 1
    getCount();
    return 0;
}