/**
* Created by soeren 13.4.2017
* License see the "LICENSE" document
*/

#include <stdio.h>
#include <stdlib.h>

struct node {
    int data;
    struct node *leftchild;
    struct node *rightchild;
};

int anzahl=0;
struct node *head=NULL;

void fillarray(int* array){
    printf("\ninsert the numbers for sorting and seperate them with backspackes ...\n");
    for (int i=0;i<anzahl;i++){
	printf("Zahl %2i: ",i+1);
	scanf("%d",&array[i]);
    }
    printf("\n");
}

void filltree(int insert,struct node *cursor){
    struct node *bfrc=NULL;

    struct node* newkey(int insert){
        struct node *new = malloc (sizeof(struct node));
        new->data=insert;
	if (head==NULL){
            head=new;
        }

	return new;
    }

    if (cursor==NULL && head==NULL){
	head=newkey(insert);
    } else if (insert <= cursor->data) {
	if (cursor->leftchild==NULL){
	    cursor->leftchild=newkey(insert);
	} else {
	    filltree(insert,cursor->leftchild);
	}
    } else {
	if (cursor->rightchild==NULL){
	    cursor->rightchild=newkey(insert);
	} else {
	    filltree(insert,cursor->rightchild);
	}
    }
}

void printInorder(struct node *root){
    if (root != NULL){
        printf("(");
        printInorder(root->leftchild);
        printf(" ");
        printf("%d", root->data); printf(" ");
        printInorder(root->rightchild);
        printf(")");
    }
}

void freeInorder(struct node *root){
    if (root != NULL){
        freeInorder(root->leftchild);
        free(root->leftchild);
        freeInorder(root->rightchild);
        free(root->rightchild);
    }
}

int main() {
	/*start*/
    printf("**************************** Binary Tree Sort ****************************");
    printf("\n\nhow many numbers do you want to sort? ");
    scanf("%d",&anzahl);

	/*creating the unsorted array*/
    int unsorted[anzahl];
    fillarray(unsorted);

	/*repeating the array*/
    int absatz=0;
    printf("*** array ***\n");
    for (int i=0;i<anzahl;i++) {
        printf("%12d",unsorted[i]);
        absatz++;
        if (absatz>8) {
            printf("\n");
            absatz = 0;
        }
    }
    printf("\n******************************************************\n\n");
	
	/*filling of the tree*/
    for (int i=0;i<anzahl;i++) {
        filltree(unsorted[i],head);
    }

	/*printing the tree inorder*/
    printf("\n\nSTART INORDER PRINT OF TREE\n");
    printInorder(head);
    printf("\n\nEND INORDER PRINT OF TREE\n");

    /*freeing the allocated memory*/
    freeInorder(head);
    free(head);

    return 0;
}
