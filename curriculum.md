# 📘 Golang TDD Learning Curriculum

> Khóa học Golang từ cơ bản đến nâng cao theo phương pháp **Test-Driven Development**.
> Mỗi topic là một bài tập độc lập: đọc lý thuyết → implement code → chạy test để verify.

---

## Section 01 – Basics

| #   | Topic            | Mô tả                                                       |
| --- | ---------------- | ----------------------------------------------------------- |
| 01  | Variables        | Khai báo biến, zero values, short declaration (`:=`), scope |
| 02  | Constants        | `const`, `iota`, typed vs untyped constants                 |
| 03  | Type Conversions | Chuyển đổi giữa các kiểu dữ liệu, type alias                |
| 04  | Operators        | Arithmetic, comparison, logical, bitwise operators          |

---

## Section 02 – Control Structures

| #   | Topic     | Mô tả                                               |
| --- | --------- | --------------------------------------------------- |
| 01  | If/Else   | Conditional statements, init statement trong `if`   |
| 02  | Switch    | Switch cases, type switch, fallthrough              |
| 03  | For Loops | `for`, `range`, `break`, `continue`, infinite loops |
| 04  | Defer     | `defer` stack, execution order, use cases           |

---

## Section 03 – Collections

| #   | Topic   | Mô tả                                                  |
| --- | ------- | ------------------------------------------------------ |
| 01  | Arrays  | Fixed-size arrays, iteration, multi-dimensional        |
| 02  | Slices  | Slice internals (len, cap), `append`, `copy`, slicing  |
| 03  | Maps    | Map CRUD, iteration order, comma-ok idiom              |
| 04  | Strings | String immutability, rune, `strings` package utilities |

---

## Section 04 – Functions

| #   | Topic              | Mô tả                                                          |
| --- | ------------------ | -------------------------------------------------------------- |
| 01  | Basic Functions    | Parameters, return values, named returns, multiple returns     |
| 02  | Variadic Functions | `...` syntax, passing slices to variadic functions             |
| 03  | Closures           | Anonymous functions, capturing variables, function factories   |
| 04  | Recursion          | Recursive algorithms, tail recursion, stack overflow awareness |

---

## Section 05 – Structs & Pointers

| #   | Topic            | Mô tả                                                 |
| --- | ---------------- | ----------------------------------------------------- |
| 01  | Structs          | Struct declaration, initialization, exported fields   |
| 02  | Methods          | Value receivers vs pointer receivers, method sets     |
| 03  | Pointers         | `&`, `*`, nil pointers, pointer to struct             |
| 04  | Embedded Structs | Composition over inheritance, promoted fields/methods |

---

## Section 06 – Interfaces

| #   | Topic                 | Mô tả                                           |
| --- | --------------------- | ----------------------------------------------- |
| 01  | Basic Interfaces      | Implicit implementation, interface satisfaction |
| 02  | Type Assertions       | Type assertion, type switch, safety checks      |
| 03  | Empty Interface       | `any` / `interface{}`, use cases và hạn chế     |
| 04  | Interface Composition | Embedding interfaces, `io.ReadWriter` pattern   |

---

## Section 07 – Error Handling

| #   | Topic           | Mô tả                                             |
| --- | --------------- | ------------------------------------------------- |
| 01  | Basic Errors    | `error` interface, `errors.New`, `fmt.Errorf`     |
| 02  | Custom Errors   | Tạo error types riêng, `Error()` method           |
| 03  | Error Wrapping  | `%w`, `errors.Is`, `errors.As`, unwrapping chains |
| 04  | Panic & Recover | Khi nào dùng `panic`, `recover` trong `defer`     |

---

## Section 08 – Concurrency

| #   | Topic             | Mô tả                                                 |
| --- | ----------------- | ----------------------------------------------------- |
| 01  | Goroutines        | `go` keyword, goroutine lifecycle, scheduling         |
| 02  | Channels          | Unbuffered channels, send/receive, channel direction  |
| 03  | Buffered Channels | Capacity, blocking behavior, khi nào dùng buffered    |
| 04  | Select            | `select` statement, non-blocking ops, timeout pattern |
| 05  | Mutex             | `sync.Mutex`, `sync.RWMutex`, race conditions         |
| 06  | WaitGroup         | `sync.WaitGroup`, đồng bộ hóa goroutines              |

---

## Section 09 – Context

| #   | Topic           | Mô tả                                                 |
| --- | --------------- | ----------------------------------------------------- |
| 01  | Context Basics  | `context.Background()`, `context.TODO()`, propagation |
| 02  | Context Timeout | `WithTimeout`, `WithDeadline`, cancellation signals   |
| 03  | Context Values  | `WithValue`, truyền request-scoped data               |

---

## Section 10 – Generics

| #   | Topic                   | Mô tả                                          |
| --- | ----------------------- | ---------------------------------------------- |
| 01  | Type Parameters         | Generic functions, type parameter syntax       |
| 02  | Type Constraints        | `comparable`, custom constraints, `~` operator |
| 03  | Generic Data Structures | Generic stack, linked list, result type        |

---

## Section 11 – Reflection

| #   | Topic              | Mô tả                                       |
| --- | ------------------ | ------------------------------------------- |
| 01  | Reflect Basics     | `reflect.TypeOf`, `reflect.ValueOf`, Kind   |
| 02  | Struct Tags        | Parsing struct tags, JSON tag pattern       |
| 03  | Dynamic Invocation | Gọi method bằng reflection, settable values |

---

## Section 12 – Design Patterns

| #   | Topic     | Mô tả                                 |
| --- | --------- | ------------------------------------- |
| 01  | Singleton | `sync.Once`, thread-safe singleton    |
| 02  | Factory   | Factory function pattern trong Go     |
| 03  | Strategy  | Strategy pattern qua interfaces       |
| 04  | Observer  | Event-driven observer với channels    |
| 05  | Decorator | Function wrapping, middleware pattern |

---

## Section 13 – Memory Management

| #   | Topic              | Mô tả                                             |
| --- | ------------------ | ------------------------------------------------- |
| 01  | Stack vs Heap      | Phân biệt stack/heap allocation trong Go          |
| 02  | Escape Analysis    | `go build -gcflags="-m"`, tránh heap allocation   |
| 03  | Garbage Collection | GC internals, tuning `GOGC`, `debug.FreeOSMemory` |

---

## Section 14 – Advanced Testing

| #   | Topic              | Mô tả                                         |
| --- | ------------------ | --------------------------------------------- |
| 01  | Table-Driven Tests | Subtests, `t.Run`, data-driven test pattern   |
| 02  | Test Fixtures      | `TestMain`, setup/teardown, temp files        |
| 03  | Mocking            | Interface-based mocking, dependency injection |
| 04  | Benchmarks         | `b.N`, `b.ResetTimer`, profiling with `pprof` |

---

## 🚀 Cách sử dụng

```bash
# 1. Chọn topic, đọc README.md
# 2. Implement TODO trong solution.go
# 3. Chạy test
go test ./section_08_concurrency/topic_02_channels/ -v

# 4. Tất cả test pass = hoàn thành topic ✅
```
