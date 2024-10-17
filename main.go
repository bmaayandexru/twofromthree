package main

import (
	"fmt"
	"os"
	"slices"
)

var (
	in      []int64
	mul, im [3]int64
	ind     = [3]int64{1, 1, 1}
	count   int64
	result  int64
)

func NextNum() int64 {
	// в mul текущие претенденты на число
	// ищем минимальный и генерим следующий
	var res int64
	if mul[0] < mul[1] && mul[0] < mul[2] {
		res = mul[0]
		// вычисляем следующего претендента
		for {
			//ind[0]++
			// mul[0] = ind[0] * im[0]
			mul[0] += im[0]
			if mul[0]%in[2] != 0 {
				// претендент не должен делиться на цело на 3й элемент
				break
			}
		}
	} else if mul[1] <= mul[0] && mul[1] < mul[2] {
		res = mul[1]
		// вычисляем следующего претендента
		for {
			//ind[1]++
			// mul[1] = ind[1] * im[1]
			mul[1] += im[1]
			if mul[1]%in[1] != 0 {
				// претендент не должен делиться на цело на 2й элемент
				break
			}
		}
	} else /*if mul[2] <= mul[0] && mul[2] <= mul[1] */ {
		res = mul[2]
		// вычисляем следующего претендента
		for {
			//ind[2]++
			//mul[2] = ind[2] * im[2]
			mul[2] += im[2]
			if mul[2]%in[0] != 0 {
				// претендент не должен делиться на цело на 1й элемент
				break
			}
		}
	}
	return res
}

func Check() bool {
	if in[0] < 1 || in[0] > 1000000 || in[1] < 1 || in[1] > 1000000 || in[2] < 1 || in[2] > 1000000 || count < 1 || count > 1_000_000_000_000_000_000 {
		return false
	}
	if in[1]%in[0] == 0 || in[2]%in[0] == 0 || in[2]%in[1] == 0 {
		return false
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	in = make([]int64, 3)
	for i, _ := range in {
		if i < len(in)-1 {
			_, err = fmt.Fscanf(file, "%d", &in[i])
		} else {
			_, err = fmt.Fscanf(file, "%d\n", &in[i])
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	_, err = fmt.Fscanf(file, "%d\n", &count)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	slices.Sort(in)
	fmt.Println(in)
	if Check() {
		mul[0] = in[0] * in[1]
		mul[1] = in[0] * in[2]
		mul[2] = in[1] * in[2]
		im[0] = mul[0]
		im[1] = mul[1]
		im[2] = mul[2]
		var i int64
		for i = 0; i < count; i++ {
			result = NextNum()
			/*
				if i%100_000_000 == 0 {
					fmt.Println(i, result)
				}
			*/
			//result *= count
			if result > 1_000_000_000_000_000_000 {
				result = -1
				break
			}
		}
	} else {
		result = -1
	}
	file, err = os.Create("output.txt")
	if err != nil {
		fmt.Printf("Error creating file %v", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d", result))
}
