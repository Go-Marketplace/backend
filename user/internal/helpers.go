package internal

func Unwrap[T any](v *T) T {
    if v != nil {
        return *v
    }
    return *new(T) // zero value of T
}
