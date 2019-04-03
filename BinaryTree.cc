#include <iostream>
#include <queue>

#include "BinaryTree.h"

using namespace std;


void
CreateBiTree(BiTree *T)
{
    char ch;
    cin >> ch;

    if (ch == '#') {
        *T = NULL;
    } else {
        *T = new BiTNode();
        (*T)->element = ch;
        CreateBiTree(&((*T)->left));
        CreateBiTree(&((*T)->right));
    }
    
}


void
PreOrderTraverse(BiTree T, int level)
{
    if (T == NULL)
        return;

    cout << T->element << " " << level << endl;
    PreOrderTraverse(T->left, level+1);
    PreOrderTraverse(T->right, level+1);
}


void
InOrderTraverse(BiTree T, int level)
{
    if (T == NULL)
        return;

    InOrderTraverse(T->left, level+1);
    cout << T->element << " " << level << endl;
    InOrderTraverse(T->right, level+1);
}


void
PostOrderTraverse(BiTree T, int level)
{
    if (T == NULL)
        return;

    PostOrderTraverse(T->left, level+1);
    PostOrderTraverse(T->right, level+1);
    cout << T->element << " " << level << endl;
}


void
LevelOrderTraverse(BiTree t)
{
    if (t == NULL)
        return;

    queue<BiTree> q;

    q.push(t);

    while (!q.empty()) {
        BiTree tmp = q.front();
        q.pop();
        cout << tmp->element << endl;

        if (tmp->left != NULL) {
            q.push(tmp->left);
        }

        if (tmp->right != NULL) {
            q.push(tmp->right);
        }
    }
}

void
PrintByLevel(BiTree t)
{
    if (t == NULL)
        return;

    queue<BiTree> q;

    int currentLevel = 1, nextLevel = 0;

    q.push(t);

    while (!q.empty()) {
        BiTree tmp = q.front();
        cout << tmp->element << " ";
        currentLevel--;

        if (tmp->left != NULL) {
            nextLevel++;
            q.push(tmp->left);
        }

        if (tmp->right != NULL) {
            nextLevel++;
            q.push(tmp->right);
        }

        q.pop();

        if (currentLevel == 0) {
            cout << endl;
            currentLevel = nextLevel;
            nextLevel = 0;
        }
    }
}
