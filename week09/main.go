package main

import (
	"fmt"
)

func main() {

	var result string
	result = fmt.Sprintf("%0.1f 나누기 %0.1f는 %0.2f입니다.\n", 1.0, 3.0, 1.0/3.0)

	fmt.Print(result)

	i := 1
	for i <= 10 {
		fmt.Printf("%5d\n", i)
		i++
	}

	// rand.Seed(time.Now().Unix())
	// answer := rand.Intn(6) + 1

	// win := false

	// for guesses := 0; guesses < 3; guesses++ {
	// 	fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력 : ", 3-guesses)
	// 	r := bufio.NewReader(os.Stdin)
	// 	i, err := r.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	i = strings.TrimSpace(i)
	// 	guess, err := strconv.Atoi(i)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(guess)

	// 	if answer == guess {
	// 		fmt.Printf("정답입니다. ")
	// 		win = true
	// 		break
	// 	} else if answer > guess {
	// 		fmt.Println("입력하신 수는 정답보다 작은 수입니다. LOW")
	// 	} else {
	// 		fmt.Println("입력하신 수는 정답보다 큰 수입니다. HIGH")
	// 	}
	// }

	// if win {
	// 	fmt.Print("당신이 이겼습니다! ")
	// } else {
	// 	fmt.Print("당신이 졌습니다. ")
	// }
	// fmt.Printf("정답은 %d입니다.", answer)
}
