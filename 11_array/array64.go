package main

import "fmt"

func main() {
	var na, nb, nc, nd int
	fmt.Print("Na = ")
	fmt.Scan(&na)

	a := make([]int, na, na)
	for i := 0; i < na; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Print("Nb = ")
	fmt.Scan(&nb)
	b := make([]int, nb, nb)
	for i := 0; i < nb; i++ {
		fmt.Scan(&b[i])
	}

	fmt.Print("Nc = ")
	fmt.Scan(&nc)
	c := make([]int, nc, nc)
	for i := 0; i < nc; i++ {
		fmt.Scan(&c[i])
	}

	nd = na + nb + nc
	var ia, ib, ic, id int
	d := make([]int, nd, nd)
	var max int
	var index *int
	var ar *[]int
	var isSetMax bool

	for id < nd {
		if !isSetMax {
			isSetMax = true
			switch {
			case ia < na:
				max = a[ia]
				index = &ia
				ar = &a
			case ib < nb:
				max = b[ib]
				index = &ib
				ar = &b
			case ic < nc:
				max = c[ic]
				index = &ic
				ar = &c
			default:
				break
			}
		}

		if ia < na && max < a[ia] {
			max = a[ia]
			index = &ia
			ar = &a
		}
		if ib < nb && max < b[ib] {
			max = b[ib]
			index = &ib
			ar = &b
		}
		if ic < nc && max < c[ic] {
			max = c[ic]
			index = &ic
			ar = &c
		}
		d[id] = max
		id++
		*index++
		if *index == len(*ar) {
			isSetMax = false
		} else {
			max = (*ar)[*index]
		}
	}

	fmt.Println()
	for i := 0; i < nd; i++ {
		fmt.Print(d[i], " ")
	}
	fmt.Println()
}
