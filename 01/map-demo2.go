package main

import (
	"fmt"
	"math"
	"sort"
)

func test1() {
	temperatrues := []float64{
		-28, 32.8, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	for _, t := range temperatrues {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d items\n", t, num)
	}
}

func test2() {
	temperatrues := []float64{
		-28, 32.8, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	groups := make(map[float64][]float64)

	for _, t := range temperatrues {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	for g, temps := range groups {
		fmt.Printf("%v: %v\n", g, temps)
	}
}

func test3() {
	temperatrues := []float64{
		-28, 32.8, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)

	for _, t := range temperatrues {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("set member")
	}

	fmt.Println(set)

	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}

	sort.Float64s(unique)

	fmt.Println(unique)
}

func main() {
	fmt.Println("========== map demo2 ==========")

	test1()

	fmt.Println()

	test2()

	fmt.Println()

	test3()
}
