# Go Generics

## Beginner Level – Basic Concepts and Usage

- What are Generics in Go?
  - Purpose and advantages of Generics
  - Support for generics in Go version 1.18
- Basic Generic Usage
  - Generic functions (`func Print[T any](value T)`)
  - Type parameters (type variables like `T, U, V`)
  - `any` keyword (to support all types)
- Writing Simple Generic Functions
  - Generic `Add[T any](a, b T) T` function
  - Generic `Swap[T any](x, y T) (T, T)` function
  - Generic `Compare[T comparable](a, b T) bool` function
- Applications and Exercises
  - Write a generic `PrintSlice[T any](s []T)` function
  - Write a generic `FindMax[T int | float64](s []T) T` function

## Intermediate – Generic Data Structures and Constraints

- Generic Data Structures
  - Using generic `struct`
  - Using generic `map` and `slice`
- Using `Constraints`
  - Using `any` and `comparable`
  - Adding constraints with custom interfaces
  - Allowing specific types with `type constraints`

```go
type Number interface {
    int | float64
}
func Sum[T Number](numbers []T) T { ... }
```

- Generic Interface Usage
  - Defining a custom interface
  - Generic `Reader[T any]` and `Writer[T any]` interfaces
- Performance and Memory Usage
  - Impact of Generics on performance
  - Understanding how Go compiles generics
  - When to use generics, when not to use them?
- Practices and Exercises
  - Write a generic `Stack[T any]` data structure
  - Write a generic `Queue[T any]` data structure
  - Implement the generic `Min[T Number](a, b T) T` function

## Advanced – In-Depth Generics and Real World Usage

- More Complex Constraints
  - Multiple type support with `type sets`
  - Using multiple generic parameters (`func Merge[K comparable, V any](m1, m2 map[K]V) map[K]V`)
  - Using `reflect` and generics
- Generic Methods and Interfaces
  - Structures containing generic methods
  - Structures implementing generic interfaces

```go
type DataStore[T any] interface {
    Save(data T)
    Get(id int) T
}
```

- Real World Scenarios
  - Write a generic cache system
  - Using a generic repository pattern
  - JSON serialization/deserialization using generics
- Practices and Exercises
  - Write a generic LRU Cache implementation
  - Write a generic Observer Pattern example
  - Write a generic Event Dispatcher

## Reference

1. [x] [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
2. [x] [Ardanlabs Generics](https://tour.ardanlabs.com/tour/eng/generics-basics/1)
3. [x] [Go by Example: Generics](https://gobyexample.com/generics)
4. [x] [Know Go: Generics](https://github.com/bitfield/kg-generics2)

5. [ ] <https://www.digitalocean.com/community/tutorials/how-to-use-generics-in-go#prerequisites>
6. [ ] <https://itnext.io/a-comprehensive-guide-to-generics-in-go-5a9dcda5669c>
7. [ ] <https://www.kelche.co/blog/go/golang-generics/>
8. [ ] <https://100go.co/9-generics/>
9. [ ] <https://blog.merovius.de/posts/2024-01-05_constraining_complexity/>
