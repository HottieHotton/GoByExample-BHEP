In the GoByExample Explaination: they explain the reason for stateful goroutines as follows:
```text
Another option is to use the built-in synchronization features of goroutines and channels to achieve the same result. This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly one goroutine.
```

However, it also mentions that you can utilize the other explained methods, depending on what is best for your application. It's good to know there's multiple ways to sync access across multiple routines without breaking anything.