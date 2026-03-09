# Hash Table
A hash table is a data structure that uses a hash function to map keys to specific locations (buckets) in the underlying array. This allows fast access to the associated values using the keys.

Primary strength lies in providing fast lookups, insertions, and deletions (average O(1) time complexity)


# Big O Notation (Time Complexity)
Average Case
    Search :                           O(1)
    Insertion :                        O(1)
    Deletion :                         O(1)

Worst Case
    Search :                           O(n) / O(log n)
    Insertion :                        O(n) / O(log n)
    Deletion :                         O(n) / O(log n)

The worst-case scenario occurs due to hash collisions, where many keys map to the same bucket. 


# Key Trade-offs
HashMaps generally use more memory compared to arrays or packed data structures because of the underlying hash table structure.

While average performance is, worst-case scenarios (high collision rates) can degrade performance, although modern implementations often mitigate this.


# Real World Applications of Queue
1. Database Indexing
2. Caching Systems 
3. Network Routing Tables
4. Load Balancing
5. Compilers