Race condition occurs when we try to access a resource, from different functions, at the same time, and try to alter its value.
It Occurs in concurrent programming.

To check for race condition type - 

```
    go run -race .
```

This race condition can be solved if we access the data safely. Alter the value of the data, after the modification is done by other function.

using mutex we can get this funtionality.

#### Points to remember with mutex
1. We should not copy mutex variable. Always pass it by reference.