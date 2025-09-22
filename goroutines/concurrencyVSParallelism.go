package main

/*
   Concurrency
   - Definition: Dealing with multiple tasks at once by interleaving their execution, making progress on each task during its "time slice".
   - Goal: To improve the responsiveness and interactivity of an application.
   - Requirement: Can be achieved on a single-core processor through time-slicing or context switching.
   - Analogy: A single cashier serving multiple customers by quickly switching between them, creating the illusion of simultaneous service.
   - Examples: A web browser handling multiple tabs, a web server processing multiple requests, or a chat application responding to different users.

   Paralleslism
    - Definition: Executing multiple tasks simultaneously, with each task running at the same exact moment.
    - Goal: To improve the speed and efficiency of the application by performing work at the same time.
    - Requirement: Needs multiple processing units, such as multiple CPU cores or processors.
    - Analogy: Multiple cashiers each serving a different customer at the exact same time.
    - Examples: A video editing application processing different parts of a video in parallel, or a scientific simulation distributing calculations across multiple cores.

    Key Differences:
    - Execution: Concurrency is about managing tasks over time; parallelism is about executing tasks at the same time.
    - Hardware: Concurrency can happen with one processor; parallelism requires multiple processors.
    - Purpose: Concurrency aims for responsiveness; parallelism aims for speed.
    - Relationship: Parallelism is a specific type of concurrency where tasks are truly simultaneous; concurrency is a more general concept that can include parallelism.
*/
