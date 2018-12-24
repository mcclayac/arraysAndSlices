package main

import "fmt"

func printer1a(w [4]string) {
	fmt.Printf("1a : ")
	for _, word := range w {
		fmt.Printf("\t%s", word)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

func printer1b(w [8]string) {
	fmt.Printf("1b : ")
	for _, word := range w {
		fmt.Printf("\t%s", word)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

func printer2(w []string) { // passing in a slice
	// no length in the slice
	fmt.Print("================================\n")
	fmt.Printf("printer2 : \t")
	for _, word := range w {
		fmt.Printf("\t%s", word)
	}
	fmt.Printf("\n")
	//w[2] = "green"
}

func main() {

	words1a := [4]string{"the", "quick", "brown", "fox"} // fixed size array declaration
	words1b := [8]string{"These", "are", "the", "times",
		"that", "try", "men's", "souls"} // fixed size array declaration
	// Passed by value .. Copying

	words2 := [...]string{"the", "quick", "brown", "fox", "jumps",
		"over", "the", "lazy", "dog"} //  variable array declaration

	words3 := []string{"the", "quick", "brown", "fox", "jumps",
		"over", "the", "lazy", "dog"} // slice

	fmt.Printf("words1a: %s\n", words1a[2])
	fmt.Printf("words1b: %s\n", words1b[2])

	// Passing array as a copy
	printer1a(words1a)
	printer1b(words1b)
	//printer2(words2)

	fmt.Printf("words1a: %s\n", words1a[2])
	fmt.Printf("words1b: %s\n", words1b[2])

	fmt.Print("================================")
	fmt.Print("\n\n\n\n")

	fmt.Printf("Len of slice words2 : %d\n", len(words2))
	fmt.Printf("Len of slice words3 : %d\n", len(words3))
	//printer2(words2)  //  this is an error
	printer2(words3)
	printer2(words3[2:5]) // a portion of the slice
	printer2(words3[2:])  // a portion of the slice
	printer2(words3[:3])  // a portion of the slice
	printer2(words3[:2])  // a portion of the slice

	words4 := make([]string, 4) // make an empty slice with a size
	words4[0] = "the"
	words4[1] = "quick"
	words4[2] = "brown"
	words4[3] = "fox"

	printer2(words4)

	fmt.Print("================================")
	fmt.Print("\n\n\n\n")
	fmt.Printf("words5\n\n")

	words5 := make([]string, 0, 4) // initialize with 0 up to 4 items
	fmt.Printf("len :%d\tcapacity: %d\n\n", len(words5), cap(words5))
	words5 = append(words5, "the")
	words5 = append(words5, "quick")
	words5 = append(words5, "brown")
	words5 = append(words5, "fox")
	fmt.Printf("len :%d\tcapacity: %d\n\n", len(words5), cap(words5))

	printer2(words5)
	words5 = append(words5, "jumps")
	fmt.Printf("len :%d\tcapacity: %d\n", len(words5), cap(words5))
	//words5[6] = "over"  // run-time error - Over the slice amount

	fmt.Print("================================")
	fmt.Print("\n\n\n\n")

	fmt.Printf("Make a copy of the slice\n")
	fmt.Printf("words5:\n")
	printer2(words5)

	fmt.Printf("Copying Slice > copy(words6,words5 )\n\n")
	words6 := make([]string, 4) // 4 element slice
	copy(words6, words5)
	printer2(words6)
	words6[2] = "blue"
	printer2(words6)
	printer2(words5)

}
