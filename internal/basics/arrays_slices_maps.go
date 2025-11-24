package basics

import "fmt"

// ArraysSlicesMaps demonstrates core collection types.
func ArraysSlicesMaps() {
	var arr [3]int
	arr[0], arr[1], arr[2] = 1, 2, 3

	slice := []string{"go", "is", "fun"}
	slice = append(slice, "!")

	m := map[string]int{
		"alice": 30,
		"bob":   25,
	}

	fmt.Println("array:", arr)
	fmt.Println("slice:", slice, "len:", len(slice), "cap:", cap(slice))
	fmt.Println("map:", m)

	if age, ok := m["alice"]; ok {
		fmt.Println("alice exists with age:", age)
	}
}
