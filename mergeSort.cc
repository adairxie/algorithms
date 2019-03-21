#include <iostream>
#include <vector>

using namespace std;


void mergeSort(int *data, int start, int end);
void merge(int *data, int start, int mid, int end);

void
mergeSort(int *data, int start, int end)
{
    if (data == NULL || start >= end || start < 0 || end < 0)
        return;

    int mid = (start + end) / 2;
    mergeSort(data, start, mid);
    mergeSort(data, mid+1, end);
    merge(data, start, mid, end);
}

void
merge(int *data, int start, int mid, int end)
{
    vector<int>  tmp;
    int i = start, j = mid + 1;

    while (i <= mid && j <= end) {
        if (data[i] < data[j]) {
            tmp.push_back(data[i++]);
        } else {
            tmp.push_back(data[j++]);
        }
    }

    while (i <= mid) {
        tmp.push_back(data[i++]);
    }

    while (j <= end) {
        tmp.push_back(data[j++]);
    }

    for (int i = 0; i < tmp.size(); i++) {
        data[start + i] = tmp[i];
    }
}

int main() {
    int data[] = {3, 2, 1, 8, 10, 12, 5, 3};
    mergeSort(data, 0, sizeof(data)/sizeof(int) - 1);
    for (int i = 0; i < sizeof(data)/sizeof(int); i++) {
        cout << data[i] << " ";
    }

    cout << endl;

    return 0;
}
