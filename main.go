package main

import "fmt"

func printer(w [4]string) {
	for _, word := range w {
		fmt.Printf("\t%s", word)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

func printer2(w []string) {   // passing in a slice
	for _, word := range w {
		fmt.Printf("\t%s", word)
	}
	fmt.Printf("\n")
	//w[2] = "green"
}

func main() {

	words := [4]string{"the", "quick", "brown", "fox"} // fixed sizearray declaration

	words2 := [...]string{"the", "quick", "brown", "fox", "jumps",
		"over", "the", "lazy", "dog"} //  variable array declaration

	words3 := []string{"the", "quick", "brown", "fox", "jumps",
		"over", "the", "lazy", "dog"} // slice

	fmt.Printf("%s\n", words[2])
	fmt.Printf("%s\n", words2[2])

	printer(words)
	printer(words)
	//printer2(words2)

	fmt.Printf("Len of slice : %d\n", len(words2))
	printer2(words3)
	printer2(words3)
	printer2(words3[2:5])  // a portion of the slice
	printer2(words3[2:])  // a portion of the slice
	printer2(words3[:3])  // a portion of the slice
	printer2(words3[:2])  // a portion of the slice

	words4 := make([]string, 4)  // make an empty slice
	words4[0] = "the"
	words4[1] = "quick"
	words4[2] = "brown"
	words4[3] = "fox"

	printer2(words4)

	words5 := make([]string, 0, 4)  // initialize with 0 up to 4 items
	fmt.Printf("len :%d\tcapacity: %d\n", len(words5), cap(words5))
	words5 = append(words5, "the")
	words5 = append(words5, "quick")
	words5 = append(words5, "brown")
	words5 = append(words5, "fox")
	fmt.Printf("len :%d\tcapacity: %d\n", len(words5), cap(words5))

	printer2(words5)
	words5 = append(words5, "jumps")
	fmt.Printf("len :%d\tcapacity: %d\n", len(words5), cap(words5))
	//words5[6] = "over"  // run-time error - Over the slice amount

	fmt.Printf("Make a copy of the slice\n")
	fmt.Printf("words5:\n")
	printer2(words5)

	fmt.Printf("Copying Slice > copy(words6,words5 )\n\n")
	words6 := make([]string, 4)  // 4 element slice
	copy(words6,words5 )
	printer2(words6)
	words6[2] = "blue"
	printer2(words6)
	printer2(words5)






}
