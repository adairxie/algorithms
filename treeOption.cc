#include "BinaryTree.h"
#include <iostream>


using namespace std;

int main(int argc, char *argv[])
{
    BiTree T = NULL;

    CreateBiTree(&T);

    cout << "pre order:" << endl;
    PreOrderTraverse(T, 0);

    cout << "in order:" << endl;
    InOrderTraverse(T, 0);

    cout << "post order:" << endl;
    PostOrderTraverse(T, 0);

    cout << "level order:" << endl;
    LevelOrderTraverse(T);

    cout << "output by level:" << endl;
    PrintByLevel(T);

    return 0;
}
