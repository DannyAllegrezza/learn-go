# Context

This sample program explores the `context` package, which was created to make it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all goroutines involved in handling a request. 

> ## Resources
* https://go.dev/blog/context
* https://www.youtube.com/watch?v=h2RdcrMLQAo

## Things to consider using context for

* Limiting the total execution time of a request
* Passing data between API boundaries (ex: a HTTP requests unique ID)