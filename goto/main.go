package main

import "fmt"

func demoGOTO() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				goto end
			}
			fmt.Println("i:", i, "j:", j)
		}
	}
end:
	fmt.Println("Exited all loops")
}

func main() {
	demoGOTO()
}