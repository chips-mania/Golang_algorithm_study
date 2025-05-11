// 나머지
package main

import (
	"fmt"
)

func main() {
	// go에는 set 데이터 타입이 없으므로 map으로 대체
	modSet := make(map[int]bool)

	// 10회 입력받아 반복
	for i := 0; i < 10; i++ {
		var n int
		fmt.Scan(&n)

		// mod정의: 입력받은 n을 42로 나눈 나머지
		mod := n % 42

		// 나머지에 대응해 mod:ture 형태로 modSet에 저장
		modSet[mod] = true
	}

	// modSet의 길이를 출력
	fmt.Println(len(modSet))
}
