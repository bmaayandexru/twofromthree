package main

import (
	"fmt"
	"os"
	"slices"
)

var (
	in      []int64
	sMul, dMul [3]int64 // последовательность и приращение последовательности
	count   int64
	result  int64
	lastseq int
)

func NextNum() int64 {
	// в sMul текущие претенденты на число
	// ищем минимальный и генерим следующий
	var res int64
	if sMul[0] < sMul[1] && sMul[0] < sMul[2] {
		res = sMul[0]
		lastseq = 0
		// вычисляем следующего претендента
		for {
			// sMul[0] = ind[0] * dMul[0]
			sMul[0] += dMul[0]
			if sMul[0] % in[2] != 0 {
				// претендент не должен делиться на цело на 3й элемент
				break
			}
		}
	} else if sMul[1] <= sMul[0] && sMul[1] < sMul[2] {
		res = sMul[1]
		lastseq = 1
		// вычисляем следующего претендента
		for {
			// sMul[1] = ind[1] * dMul[1]
			sMul[1] += dMul[1]
			if sMul[1]%in[1] != 0 {
				// претендент не должен делиться на цело на 2й элемент
				break
			}
		}
	} else /*if sMul[2] <= sMul[0] && sMul[2] <= sMul[1] */ {
		res = sMul[2]
		lastseq = 2
		// вычисляем следующего претендента
		for {
			sMul[2] += dMul[2]
			if sMul[2]%in[0] != 0 {
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
	//fmt.Println("числа ", in)
	if Check() {
		dMul[0] = in[0] * in[1]
		dMul[1] = in[0] * in[2]
		dMul[2] = in[1] * in[2]
		lastseq = -1
		St := in[0]+in[1]+in[2]-3	// количество полезных шагов на период
		Nt := count / St			// количество полных периодов
		abc := in[0]*in[1]*in[2]	
		sMul[0] = dMul[0] + Nt * abc	// числа на конец последнего периода 
		sMul[1] = dMul[1] + Nt * abc
		sMul[2] = dMul[2] + Nt * abc
		//fmt.Printf("Шагов на период %d. Полных периодов %d. Шагов за полных приодов %d\n",St, Nt, St*Nt)
		//fmt.Println("числа ",sMul)	
		var i int64
		for i = St*Nt; i < count; i++ {
			result = NextNum()
			if result > 1_000_000_000_000_000_000 {
				result = -1
				break
			}
		}
		//fmt.Println(result,sMul)
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
