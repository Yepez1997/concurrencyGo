# concurrency-go-notes
Concurrency Go Notes 

## Chapter 1 Why Concurrency is Hard ? 
  * `Race Conditions` 
    - Two or more opertations trying to access the same resource. In other words, two or more operations must execute in the correct order; however, the program does not guarantee
  * `Atomicity`
    - When something is defined to be atomic, it is indivisible, or uninteruptable - in the context it is in 
    - Critcial to define the scopoe something is atomic in.
    - Will happen independently in the context that it is in
    - Something is atomic -> safe in its concurrrent context
  * `Memory Access Synchronization` 
    - When two or more programs are attemping to access the same are of memory, the way memory is accessed is not atomic.
    - The section of a program that needs exclusive access is called the `critical section`
    - Typically add synchronization to critical sections 
    - Synchronizing memory has performance ramifications 
  * `Deadlocks, Livelocks, Starvations` 
    - Program can stop working all together if none of these properties are satisfied
    1. Deadlock
    2. Livelock 
    3. Starvation

