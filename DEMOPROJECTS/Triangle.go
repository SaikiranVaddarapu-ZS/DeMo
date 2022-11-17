package main

func Triangle(a, b, c int) string {
	if a == 0 || b == 0 || c == 0 {
		return "Sides cannot be zero"
	}
	if (a+b) < c || (b+c) < a || (c+a) < b {
		return "Cannot form a Triangle"
	}
	if a == b && b == c {
		return "Equilateral"
	}
	if a == b || b == c || c == a {
		return "Isosceles"
	}
	return "Scalene"
}

// func main() {
//	fmt.Println(Triangle(2, 3, 0))
//}
