package calculator

type Number interface {
    int | int64 | float32 | float64
}

func Add[T Number](a, b T) T {
    return a + b
}

func Subtract[T Number](a, b T) T {
    return a - b
}

func Multiply[T Number](a, b T) T {
    return a * b
}
