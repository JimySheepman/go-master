package tutorial02

type FuncMap[T, U any] map[string]func(T) U

func (fm FuncMap[T, U]) Apply(name string, val T) U {
	return fm[name](val)
}

func Compose[T, U, V any](f func(U) T, g func(V) U, v V) T {
	return f(g(v))
}
