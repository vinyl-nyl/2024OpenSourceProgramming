package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("점수입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	// counts := 0
	var isPrime bool = true // 가독성, 메모리 사이즈 개선

	if n < 2 { // 1보다 큰 자연수 중 1과 자기 자신만을 약수로 가지는 수
		isPrime = false
	} else {
		for j := 2; j <= int(math.Sqrt(float64(n))); j++ { // 1부터 입력된 수까지 반복
			if n%j == 0 { // 약수면
				// counts++ // 나누어 떨어지는 횟수 카운트
				isPrime = false
				break
			}
			fmt.Printf("%d ", j) // 반복문 횟수 확인용 코드
		}
	}

	// if counts == 0 { // 나누어 떨어지는 수가 있으면 안됨
	if isPrime { // 비교 연산 제거
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다.", n)
	}
}
