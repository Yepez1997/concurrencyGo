# Concurrency Go Notes
Concurrency Go Notes 

## Chapter 1 Why Concurrency is Hard ? 
   Program behavior is non determinstic 
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
    1. `Deadlock`
      - All concurrent proccesses are waiting on each other to access a resource 
      - Program will never recover without outside intervention
      - In other words go routines infinitly wait on eacho other 
      - Few conditions must be met to prevent, detect, and correct a deadlock - Edgar Coffman
       #### Coffman Conditions 
        i. Mutual Exclusion
        ii. Wait for Condition
        iii. No Preemption
        iv. Circular Wait
    2. `Livelock` 
      - Programs that are actively running concurrent operations; however, the operations do notn move the state of the 
      program
      - Two or more proccesses are attemping to prevent a deadlock without coordination 
      - Livelocks more difficult to spot than deadlocks 
      - Livelocks are part of a larger set of problems known as starvation 
      - All concurrent proccess are starved equally, or no work is accomplished 
    3. `Starvation`
     - Any situation where concurrent proccesses cannot get all the resources it needs to perfomr work 
     - In livelocks the resource that was starved was a shared lock
     - One or more greedy proccesses that are unfairly preventing one or more concurrent processes from accomplishing work as     efficiently as possible
     - Greedy worker holds on to a shared lock for its entirety whereas polite worker only hold one at a time - or perhaps only when it needs to 
     - Keeping track of metrics is critical to discovering starvation 

