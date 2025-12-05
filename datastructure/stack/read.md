# Heap
Stack is a linear data structure that follows LIFO (Last In First Out) Principle, the last element inserted is the first to be popped out. It means both insertion and deletion operations happen at one end only.

# LIFO(Last In First Out) Principle
- New elements are always pushed on top.
- Removal (pop) also happens only from the top.
- This ensures a strict order: last in → first out.

# Types of Stack:
- Fixed Size Stack 
- Dynamic Size Stack 

# Common Operations on Stack:
- push() to insert an element into the stack.
- pop() to remove an element from the stack.
- top() Returns the top element of the stack.
- isEmpty() returns true if stack is empty else false.
- size() returns the size of the stack.

# Benefits of stack with LinkedList
- The main advantage of using a linked list over arrays is that it is possible to implement a stack that can shrink or grow as much as needed.
- Using an array will put a restriction on the maximum capacity of the array which can lead to stack overflow. Here each new node will be dynamically allocated. so overflow is not possible.
- If we use built in dynamic sized arrays like vector in C++, list in Python or ArrayList in Java, we get automatically growing stack, but the worst case time complexity is not O(1) for push() and pop() as there might be a resizing step once in a while. With Linked List, we get worst case O(1).

# Application
- Undo/Redo operation
- Recursion
- Memory management