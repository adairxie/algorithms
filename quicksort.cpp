#include <iostream>
#include <stdlib.h>



int quickSort(int *array, int start, int end);
int partition(int *array, int start, int end);
void swap(int *array, int i, int j);

int quickSort(int *array, int start, int end) {
    if (array == NULL || start < 0 || end < start) {
        return -1;
    }

    int smallIndex = partition(array, start, end);
    if (smallIndex > start) {
        quickSort(array, start, smallIndex-1);
    }

    if (smallIndex < end) {
        quickSort(array, smallIndex+1, end);
    }

    return 0;
}

int partition(int *array, int start, int end) {
    int privot = start + rand()%(end - start + 1);
    swap(array, privot, end);
    int smallIndex = -1;
    for (int i = start; i < end; i++) {
        if (array[i] <= array[end]) {
            smallIndex++;
            if (i > smallIndex) {
                swap(array, smallIndex, i);
            }
        }
    }

    return smallIndex;
}

void swap(int *array, int i, int j) {
    int temp = array[i];
    array[i] = array[j];
    array[j] = temp;
}

int main() {
    int arr[6] = {4, 3, 2, 6, 5, 1};
    quickSort(arr, 0, 5);
    return 0;
}
