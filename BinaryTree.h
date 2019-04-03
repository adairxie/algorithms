#ifndef _BINARYTREE_H_INCLUDE_
#define _BINARYTREE_H_INCLUDE_


#include <string>

typedef struct BiTNode{
    char element;
    struct BiTNode *left, *right;
}BiTNode, *BiTree;


void CreateBiTree(BiTree *T);
void PreOrderTraverse(BiTree T, int level);
void InOrderTraverse(BiTree T, int level);
void PostOrderTraverse(BiTree T, int level);
void LevelOrderTraverse(BiTree T);
void PrintByLevel(BiTree T);

#endif /* _BINARYTREE_H_INCLUDE_ */
