package main

import "fmt"

func mayorMenor(vec [5]int) (int, int) {
	var may = vec[0]
	var men = vec[0]
	for f := 1; f < len(vec); f++ {
		if vec[f] > may {
			may = vec[f]
		} else {
			if vec[f] < men {
				men = vec[f]
			}
		}
	}
	return may, men
}

func main() {
	vector := [5]int{12, 2, 30, 1, 5}
	may, men := mayorMenor(vector)
	fmt.Println("El elemento mayor del vector es:", may)
	fmt.Println("El elemento menor del vector es:", men)
}
