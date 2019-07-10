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
     - Constrain memory access synchronization to critical sections 
     - Starvation can cause a program to behave inefficiently or incorrectly 
     - Can happen with outside resources 
   * `Concurrency Safety`
      - concurency is a difficult area in computer science 
      - with go, can more cleary and safely write concurrent algorithms 
      - go's concurrent, low-latency, garbage collector
      - no need for a thread pool 
      - go's concurrency primitives make it easier to compose larger concurrent programs - channels, selectors

## Chapter 2  Modeling Your Code: Communicating Sequential Processes?
   * `Concurrency vs Parallelism`
      - Concurrency is the property of the code, parrallism is the propertyof running the program
      - Code is not parallel, code is concurrent - in hopes that it will be parallel
      - Parrallism is a property of runtime 
      - Go primitives include channels, and go routines 
    * `Communicating Sequential Processes`
      - Tony Hoare paper 1978 
      - model input and output between processes correctly
      - go uses the guarded command - invented by Dijkstra 
    
## Chapter 3  Go's Concurrency Building Blocks 
   * `Goroutines`
   * `Sync Package `
   * `Wait Groups`
   * `Cond`
   * `Once`
   * `Pool`
   * `Channels`
   * `Queues`
   * `Select`
   * `GOMAXPROCSLEVER`
      
## Chapter 4 Concurrency Patterns in Go
   * `Confinement`
      - Lexical Confinement 
      - Ad hoc Confinement
   * `For Select`
   * `Preventing Go Routine Leaks`
   * `Or Channel`
    - Combine one or more done channels wiht a single channel if any of the n channels passed in closes 
   * `Error Handling`
    - Should always seperate concerns with logic
   * `Pipelines`
      - Can be thought of as higher order functions or monads
      - Batch Proccesing -> opperate on chunks of data at once
      - Stream Proccesing -> 
        opperate on a stage that receives one element at time  
   * `Fan In Fan Out`
    - Proces of starting multiple goroutines to handle input from a pipeline 
    - Fanning in means multiplexing streams of data into a single stream 
    - For fan in pattern order should be unimportant 
   * `Or Done Channel`
    - Checks if a valid channel was returned and not cancelled (revisit this)
   * `Tee Channel`
    - Used when the values returned from a channel might be sent to different areas of the code base etc etc.
   * `Bridge Channel`
   - Used when want to retreive values from a sequence of channels
   * `Queuing`
      - Queing Theory
      - Little's Law- Predict the throughout of pipeline
   * `Context Package`
      - often neccesary to preempty operations because of time outs,
      cancellationm or failure of another part of the system. 
      - may be useful to communicate useful information alongside the simple notification to cancel 
      - very common to wrap in systems of any size 
      - useful to know why the system was cancelled or whether function has a deadline to complete 
      - context type will flow through the system 
      - deadline should return the time when the work done on behalf of the context should be cancelled 
        - returns ok == false when no deadline is set 
      - done returns a channel thats closed when work done on behalf of this context should be cancelled 
        - done may return nil if this context could never be cancelled 
      - err returns a non nil error value after done is closed 
      - value returns the value associated with this context for key 
      - done returns a channel thats closed 
      - deadline to indicate if a go routine was cancelled 
      - primary uses of go routine was to service requests 
      - purpose of the value function is so that request specifix information needs to be passed along in addition to information about preemption 
      - 'main uses of context package' 
        - to provide an api for canceling branches of your call graph 
        - to provide a data bag for transporting scoped data through your call graph 
      - cancelation 
        - has three aspects 
          - a go routines parent may want to cancel it 
          - a go routine may want to cancel its children 
          - any blocking operations within a goroutine need to be pre emptable so that it may be canceled 
          - childen cannot cancel context stack 
      - withCancel 
      - withDeadline 
      - withTimeout 
      - succesive layers can create a context that adheres to their needs 
        - composeable and elegant to manage braches 
        - call graph is asychronous 
      - at each stack frame  a function can affect the entirety of the call below it 
      - use context to store data when it transmites processes and api boundaries 
      1. data should tarnsmite process or api boundaries 
      2. data should be immutable 
      3. data should tend toward simple tyoes 
      4. data should be data - not types with methods 
      5. data should help decorate operations - not drive them 
  ## Chapter 5 Concurrency at Scale
      

      
      
      
