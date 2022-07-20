package main

import (
	"fmt"
	"strconv"
	"strings"
)

const inputTrial = `
5 1 9 5
7 5 3
2 4 6 8`

const input = `
409  194  207  470  178  454  235  333  511  103  474  293  525  372  408  428
4321  2786  6683  3921  265  262  6206  2207  5712  214  6750  2742  777  5297
3536  2675  1298  1069  175  145  706  2614  4067  4377  146  134  1930  3850
2169  1050  3705  2424  614  3253  222  3287  3340  2637  61  216  2894  247
214  99  797  80  683  789  92  736  318  103  153  749  631  626  367  110  805
2922  1764  178  3420  3246  3456  73  2668  3518  1524  273  2237  228  1826  182
4682  642  397  5208  136  4766  180  1673  1263  4757  4680  141  4430  1098  188 
158  712  1382  170  550  913  191  163  459  1197  1488  1337  900  1182  1018  337
3858  202  1141  3458  2507  239  199  4400  3713  3980  4170  227  3968  1688  4352
4168  209 `

func main() {
	checksum := 0
	var rows = strings.Split(strings.TrimSpace(input), "\n")
	for _, row := range rows {
		numbersInStr := strings.Split(strings.TrimSpace(row), " ")
		max, _ := strconv.Atoi(numbersInStr[0])
		min, _ := strconv.Atoi(numbersInStr[0])

		for _, numberInStr := range numbersInStr {
			num, _ := strconv.Atoi(numberInStr)
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		checksum += (max - min)
	}
	fmt.Println(checksum)
}
