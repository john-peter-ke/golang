# Linked Lists
A linked list is a fundamental data structure in computer science, where each element, known as a node, contains data and a reference (or link) to the next node in the sequence. 

Unlike arrays, linked lists do not store elements contiguously in memory; instead, each node points to the next, allowing for efficient insertion and deletion operations without reorganizing the entire data structure.


# Types of Linked Lists
- Singly Linked Lists (SLLs): Each node contains data and a reference to the next node.

- Doubly Linked Lists (DLLs): Each node contains data, a reference to the next node, and a reference to the previous node; thereby setting a bi-directional structure which gives a lot of advantages over the SLLs.

- Circular Linked Lists (CLLs): The last node points back to the first node, forming a circle, thereby forming a continuous loop, a Circular Linked Lists could occur in two forms; the circular singly-linked list & the the circular doubly linked list.


# Big O Notation (Time Complexity)
 Search :                           O(N)

 Insertion (at Head) :              O(1)
 Deletion (at Head) :               O(1)

 Insertion/Deletion (at Tail) :     O(N)
 Insertion/Deletion (at Tail) :     O(N)
 Insertion/Deletion (in Middle) :   O(N)


# Key Trade-offs
The primary drawback of all linked lists, whether sorted or unsorted, is the lack of random access (i.e., you cannot jump to the 50th element in O(1) time).

This makes binary search impossible and slows down algorithms that rely on indexing. Searching always requires linear time, O(n), even in a sorted linked list, because you have to traverse the nodes sequentially from the head


# Real World Applications of Linked Lists
1. Implementing Stacks and Queues
2. Web Browsers (History Tracking)
3. Undo/Redo Functionality
4. Music Playlist Management
5. Memory Management in Operating Systems