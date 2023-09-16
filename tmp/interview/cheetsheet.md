# Interview Cheet Sheet

### Data Structure

- Hash map implementation
- tree (BST) => in-order pre-order post-order
- Heap

-

### Algo

- Sorting(heap, merge, quick) time and space complexity
- sliding windows
- greedy
- DFS(Stack) for edge solution
  BFS(Queue) for shortest path
    O(V + E) when Adjacency List is used
    O(V^2) when Adjacency Matrix is used

  Djikstra's Algorithm?
-
- minimum spanning trees
- interval scheduling
- uniform pricing
- integer knapsack
- interval pricing

  Uncommon: Topological sort, Dijkstra’s algorithm
Rare: Bellman-Ford algorithm, Floyd-Warshall algorithm, Prim’s algorithm, and Kruskal’s algorithm

- two pointers
- sequence
- list (reverse, detect cycle, merge, remote kth elem)

### Algo common pitfall

- Validate input: Negative value or Nil value, Range of value
- Alog constraints: time and space
- Size
- Duplication
- Index start point
- Look over again for spotting bugs and restructure
- corner case(overflow)

-
-

### Golang

- string concate, substring
- slice internal
- channel, go routine (how to stop)
- context, sync.once
- bit operation

    num & 1 == 0 => even
    Test kth bit is set: num & (1 << k) != 0
    Set kth bit: num |= (1 << k)
    Turn off kth bit: num &= ~(1 << k)
    Toggle the kth bit: num ^= (1 << k)
    To check if a number is a power of 2: num & num - 1 == 0.

- num, err := strconv.Atoi()

### Java

### Javascript

### SOLID

- Single Responsibility
- Open-Close
- Liskov substitution
- Interface Segregation
- Dependency Inversion

### Design Pattern

### RDB

ACID

- Atomic
- Consistency
- Isolation
  - RU
  - RC
  - RR
  - Serialization
- Durability

### NOSQL

CAP

- Consistency
- Availability
- Partition tolerance

##### 4 Type

- Column
- Document
- Graph
- Key-Value

### AWS

EC2(security group, acl), ECS(task definition, ), IAM
SSM(parameter store, secret manager)
Cloudformation
SQS, SNS, Pinpoint, CloudWatch, ACM, KMS,
RDBMS, ElasticSearch, Lambda
CloudFront, S3, EBS, VPC, Kinesis, Auto-Scaling, Dynamo DB, EFS
Step Funcitons
