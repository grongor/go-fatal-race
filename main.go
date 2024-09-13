package main

func main() {
	m := make(map[int]int)

	go func() {
		for {
			m[0] = 0
		}
	}()

	for {
		m[0] = 0
	}
}
