#include <iostream>

class Node {
public:
    int data;
    Node* next;

    Node(int data) {
        this->data = data;
    };
};

void reverseLinkedList(Node** head) {
    if (head == NULL) {
        return;
    }

    Node* prev = *head;
    Node* current = (*head)->next;
    Node* next = NULL;

    while (current != NULL) {
        next = current->next;
        current->next = prev;
        prev = current;
        current = next;
    }

    (*head)->next = NULL;
    *head = prev;
}

int main(int argc, char* argv[]) {
    
    Node* head = new Node(1);
    head->next = new Node(2);
    Node* temp = head->next;
    temp->next = new Node(3);
    temp = temp->next;
    temp->next = new Node(4);

    temp = head;
    while (temp != NULL) {
        std::cout << temp->data << std::endl;
        temp = temp->next;
    }

    reverseLinkedList(&head);

    std::cout << "reversed!" << std::endl;
    temp = head;
    Node* current;
    while (temp != NULL) {
        std::cout << temp->data << std::endl;
        current = temp;
        temp = temp->next;
        delete current;
    }

    
    return 0;
}
