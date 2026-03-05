# 🚀 Golang TDD Learning

Khóa học Golang từ **cơ bản đến nâng cao** theo phương pháp **Test-Driven Development**.

## 📋 Cách học

1. Chọn section/topic theo thứ tự trong [`curriculum.md`](curriculum.md)
2. Đọc `README.md` trong folder topic để nắm lý thuyết
3. Mở `solution.go` — implement các function có `// TODO`
4. Chạy test để verify:

```bash
go test ./section_01_basics/topic_01_variables/ -v
```

✅ **Mục tiêu**: tất cả tests PASS!

## 📂 Cấu trúc

```
section_XX_<chủ_đề>/
  └── topic_XX_<kiến_thức>/
        ├── README.md          ← Lý thuyết
        ├── solution.go        ← Code template (TODO)
        └── solution_test.go   ← Unit tests
```

## 📚 Sections

| #   | Section            | Topics                                                            |
| --- | ------------------ | ----------------------------------------------------------------- |
| 01  | Basics             | variables, constants, type_conversions, operators                 |
| 02  | Control Structures | if_else, switch, for_loops, defer                                 |
| 03  | Collections        | arrays, slices, maps, strings                                     |
| 04  | Functions          | basic_functions, variadic, closures, recursion                    |
| 05  | Structs & Pointers | structs, methods, pointers, embedded_structs                      |
| 06  | Interfaces         | basic_interfaces, type_assertions, empty_interface, composition   |
| 07  | Error Handling     | basic_errors, custom_errors, error_wrapping, panic_recover        |
| 08  | Concurrency        | goroutines, channels, buffered_channels, select, mutex, waitgroup |
| 09  | Context            | basics, timeout, values                                           |
| 10  | Generics           | type_parameters, type_constraints, generic_data_structures        |
| 11  | Reflection         | reflect_basics, struct_tags, dynamic_invocation                   |
| 12  | Design Patterns    | singleton, factory, strategy, observer, decorator                 |
| 13  | Memory Management  | stack_vs_heap, escape_analysis, garbage_collection                |
| 14  | Advanced Testing   | table_driven_tests, test_fixtures, mocking, benchmarks            |

## 🛠 Commands

```bash
# Test 1 topic
go test ./section_01_basics/topic_01_variables/ -v

# Test toàn bộ
go test ./... -v

# Chạy benchmarks
go test ./section_14_advanced_testing/topic_04_benchmarks/ -bench=. -benchmem

# Check escape analysis
go build -gcflags="-m" ./section_13_memory_management/topic_02_escape_analysis/
```

## ⚙️ Yêu cầu

- Go 1.22+
