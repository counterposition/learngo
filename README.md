# Dependency Injection with samber/do

[samber/do](https://github.com/samber/do) is a minimal dependency injection framework for Go that supports generics and does not use reflection.

The code in this branch sets up a simple HTTP server that exposes a single endpoint, `/`.
The request handler keeps track of the cumulative number of requests received.
It uses an [in-memory redis](https://github.com/alicebob/miniredis) to store state.

The HTTP server depends on a logger (log/slog from the standard library) and an HTTP handler.
The HTTP handler itself depends on a logger and a redis client.

All of these elements are automatically wired together by `samber/do` using its `Provide` and `Invoke` functions. 

---
*Click the "branches" button up top [to see recent commits](https://github.com/counterposition/learngo/branches/active)*
