package main

import (
	"fmt"
	"os"
	"slices"
)

var (
	in      []int64
	mul, im [3]int64
	ind     [3]int64
	count   int64
	result  int64
)

func NextNum() int64 {
	// в mul текущие претенденты на число
	// ищем минимальный и генерим следующий
	var res int64
	if mul[0] < mul[1] && mul[0] < mul[2] {
		res = mul[0]
		ind[0]++
		// вычисляем следующего претендента
		for {
			// mul[0] = ind[0] * im[0]
			mul[0] += im[0]
			if mul[0]%in[2] != 0 {
				// претендент не должен делиться на цело на 3й элемент
				break
			}
		}
	} else if mul[1] <= mul[0] && mul[1] < mul[2] {
		res = mul[1]
		ind[1]++
		// вычисляем следующего претендента
		for {
			// mul[1] = ind[1] * im[1]
			mul[1] += im[1]
			if mul[1]%in[1] != 0 {
				// претендент не должен делиться на цело на 2й элемент
				break
			}
		}
	} else /*if mul[2] <= mul[0] && mul[2] <= mul[1] */ {
		res = mul[2]
		ind[2]++
		// вычисляем следующего претендента
		for {
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

func Steps(num int64) int64 {
	var abc int64 = in[0] * in[1] * in[2]
	return num/(in[0]*in[1]) + num/(in[0]*in[2]) + num/(in[1]*in[2]) - 3*num/abc
}

func CMax(c int64) int64 {
	return c * in[0] * in[1] * in[2] / (in[0] + in[1] + in[2] - 3)
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
	fmt.Println("числа ", in)
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
			if result > 1_000_000_000_000_000_000 {
				result = -1
				break
			}
		}
		fmt.Println("i ", i, " result ", result)
	} else {
		result = -1
	}
	fmt.Println("итерации", ind, "текущие числа", mul)
	Max := CMax(count)
	sMax := Steps(Max)
	fmt.Println("максимальное по шагам ", Max, "шагов до числа", sMax)
	file, err = os.Create("output.txt")
	if err != nil {
		fmt.Printf("Error creating file %v", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d", result))
}
