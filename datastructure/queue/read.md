# Queue
The queue data structure (DSA), which operates on the
First-In, First-Out (FIFO) principle, has numerous applications in computer science and the real world where tasks or data need to be processed in the order they arrive.


# Types of Queue
- Simple Queue: Simple (Linear) Queues are the most straightforward representation of the Queue Data structure which adheres strictly to FIFO pattern with elements added sequentially at the rear and removed from the front. 

- Circular Queue: The Circular Queue Links the last position of the Queue back to the first; this process is called “Wrap Around”. A Mechanism that allows for the highly efficient reuse of memory slots within the array as elements cycle through, thereby filling up spaces vacated by the dequeued (removed) elements

- Priority Queue(PQ): This is a special variant of Queues for which each element is assigned a priority value, which means the elements are enqueued or dequeued based on their priority value I.e elements with the highest priority are removed first irrespective of the insertion position (yeah, it doesn’t work like a typical queue), however in the case such that they all have equal priority then we adore to the FIFO pattern. PQs are used a lot in implementing heaps (binary heaps to be specific), they are also used in Dijkstra ’s Algorithm (Maths & CS students can relate) and Data Compression.


# Implementation Types
- Linked List: The operations are typically O(1) because you maintain pointers to both the front and the rear of the list, allowing constant-time access to the points of insertion and removal.

- Dynamic Array: Operations are O(1) on average (amortized constant time). In the worst case, if the array needs to be resized, it might take longer, but this happens infrequently.

- Linear Array: The dequeue operation can be O(n) because removing an element from the front would require shifting all subsequent elements to fill the gap to maintain a contiguous array, which scales with the size of the queue. This is inefficient, which is why circular arrays or linked lists are generally preferred for queue implementation.

- Circular Array: Both enqueue and dequeue operations have a worst-case time complexity of O(1), as pointers wrap around the array, avoiding element shifting or costly resizing for a fixed capacity


# Big O Notation (Time Complexity)
 enqueue:                           O(1)

 dequeue :                          O(1)

 peek / front :                     O(1)

 search :                           O(N)


# Real World Applications of Queue
1. Message Queues
2. Email Systems
3. Printer Spooling:
4. Customer Service
5. Media Playlists
6. Interrupt Handling
7. CPU Scheduling
8. Task Scheduling in Operating Systems