# Word frequency service
Small in-memory store of words to frequencies, this service can be queried for the top-k frequent words,
the least frequent word and the median frequent word.

### High level overview:
- Internal data structures for maintaining the above statistics are hash-map and heaps. the hash map is used to store all word-frequency key value pairs.
  for top-K we maintain a min-heap in order to keep k frequent words fetch-able in O(1) unsorted and O(nlogn) sorted. for the median and least frequent words
  we maintain min-max heap and a min-heap, the max aspect of the min-max heap is used with the other min-heap to maintain the median
  fetch-able in O(1) time, and the min aspect of the min-max is used in order to query the least frequent in O(logn) time,
  this was a tradeoff between maintaining another separate min-heap in order to fetch in O(1) but using more memory.


- The service process insertions in an event-loop style behaviour. it processes every insertion sequentially in order to keep its internal datastructures invariants.
  when an insertion request is made it is simply added to the insertion channel and the main thread is responsible for processing it.
  it is also worth mentioning that when an insertion event is being handled, a request for the statistics is blocked and vice-versa,
  again to ensure correctness of the internals.

### Future thoughts:

- top-k frequent words can be optimized using [count-min-sketch](https://en.wikipedia.org/wiki/Count%E2%80%93min_sketch)
, a fixed size probabilistic data structure which can return efficiently an estimation of the frequency of some word with a small error.


- parallelising insertions by distributing the load among shards using consistent hashing on the words, a coordinator will 
  decide for each word its shard and send it to it. when querying, a typical scatter-gather approach can be used in order to aggregate
  all results from all the shards. the coordinator can be horizontally scaled as well as a stateful set and kept behind a load balancer to distribute the load.

### How to run:
Configuration:
````yaml
# listening address for incoming HTTP requests
listen_address: "0.0.0.0:5000"
store:
  # top-k word frequencies to return
  k: 5

  # insertion channel buffer capacity
  capacity: 1000
````
1. `git clone` this repository
2. `cd` to the repository
3. `make build-docker` (make sure docker is running)
4. `make run-docker` (make sure the port exposed in the docker run command in the Makefile is the same as the port in the configuration file)
5. service is available at `localhost:5000/api/v1`

- `make test` to run tests

### API endpoints:

- `POST /api/v1/insert_words` 

    Payload: 
   ```json
   {
    "words": "table, mouse, banana"
   }
   ````

- `GET /api/v1/get_stats` for the desired statistics