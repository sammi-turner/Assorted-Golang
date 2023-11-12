package main

import "fmt"

func DNAStrand(dna string) string {
	s := ""
	l := len(dna)
	
	i := 0
	for i < l {
		switch dna[i] {
    	case 'A':
        	s += "T"
    	case 'T':
        	s += "A"
    	case 'C':
        	s += "G"
		case 'G':
			s += "C"
    	}
		i++
	}

	return s
}

func main() {
	fmt.Printf("GATTACA has a compliment of %s\n", DNAStrand("GATTACA"))
}
