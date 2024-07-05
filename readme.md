# Install

Run command
```bash
go install github.com/flamurey/gobench@latest
```

# Ho to run

CMD format:

```bash
gobench [-workers=size_of_worker_pool] [-threads=size_of_thread_pool] [-single]
```

To see more details run:
```bash
gobench -h
```

# Results

Intel i7-3770k

```
gobench -single
Running bench with config {WorkersCount:1 ThreadsCount:1} 
Results of light payload
Total:  38438869         Average:      52ns 
Results of hard payload
Total:      7575         Average: 264.026µs
gobench
Running bench with config {WorkersCount:8 ThreadsCount:8} 
Results of light payload
Total: 126238882         Average:      15ns 
Results of hard payload
Total:     29406         Average:  68.013µs  
```

MacBook Pro/ Apple M1 Pro

```
gobench -single
Running bench with config {WorkersCount:1 ThreadsCount:1} 
Results of light payload
Total:  50533976         Average:      39ns 
Results of hard payload
Total:      6387         Average: 313.136µs 
gobench        
Running bench with config {WorkersCount:10 ThreadsCount:10} 
Results of light payload
Total:  25241130         Average:      79ns 
Results of hard payload
Total:     51191         Average:  39.069µs
```

AMD Ryzen 5 7600

```
gobench -single
Running bench with config {WorkersCount:1 ThreadsCount:1} 
Results of light payload
Total:  66349874         Average:      30ns 
Results of hard payload
Total:     10144         Average:  197.16µs 
gobench        
Running bench with config {WorkersCount:12 ThreadsCount:12} 
Results of light payload
Total:  178901644        Average:      11ns 
Results of hard payload
Total:     60786         Average:  32.902µs
```
