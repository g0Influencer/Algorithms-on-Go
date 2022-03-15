package main

import (
	"fmt"
)

// Этап 1: формирование таблицы смещений

func BMH(t, a string) {
	N := len(a)
	set := make(map[string]bool) // уникальные символы в образе
	M := len(t)                  // число символов в образе
	dic := make(map[string]int)  // словарь смещений

	for i := M - 2; i >= 0; i-- {
		ok := set[string(t[i])]
		if !ok {
			dic[string(t[i])] = M - i - 1
			set[string(t[i])] = true
		}
	}
	ok := set[string(t[M-1])]
	if !ok {
		dic[string(t[M-1])] = M

	}
	dic["*"] = M
	fmt.Println(dic)

	// Этап 2: поиск образа в строке
	var off int
	var flBreak bool
	if N >= M {
		i := M - 1
		for i < N {
			k := 0
			j := 0
			flBreak = false
			for j = M - 1; j >= 0; j-- {
				if a[i-k] != t[j] {
					if j == M-1 {
						_, aw := dic[string(a[i])]
						if !aw {
							off = dic["*"]
						} else {
							off = dic[string(a[i])]
						}
					} else {
						off = dic[string(t[j])]
					}
					i += off
					flBreak = true
					break
				}
				k += 1
			}
			if !flBreak {
				fmt.Println("образ найден по индексу: ", i-k+1)
				i += dic["*"]
				fmt.Println(dic)
			}

		}
	} else {
		fmt.Println("Образ не надйен")
	}
}
func main() {

	t := "data"
	a := "big data data data "

	BMH(t, a)

}
