package errors

import (
	"errors"
	"fmt"
)

// Sentinel error.
var ErrNotFound = errors.New("not found")

// find returns either a result or a sentinel error.
func find(key string, data map[string]string) (string, error) {
	v, ok := data[key]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

// ErrorsDemo shows creating, wrapping, and testing errors.
func ErrorsDemo() {
	data := map[string]string{"go": "fast", "rust": "safe"}

	v, err := find("go", data)
	fmt.Println("find go:", v, "err:", err)

	_, err = find("python", data)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("python is not in the map (ErrNotFound)")
	}

	// Wrapping an error with extra context.
	_, err = find("java", data)
	if err != nil {
		wrapped := fmt.Errorf("looking up key %q: %w", "java", err)
		fmt.Println("wrapped error:", wrapped)
	}
}
