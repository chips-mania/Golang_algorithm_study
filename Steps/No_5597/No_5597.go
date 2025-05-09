// 과제 안 내신 분..?

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 출석번호 1~30을 나타내는 맵 생성
	studentNumber := make(map[int]bool)

	// 전체 학생은 true로 초기화
	for i := 1; i <= 30; i++ {
		studentNumber[i] = true
	}

	// 제출한 학생은 true에서 false로 변경
	var n int
	for i := 0; i < 28; i++ {
		fmt.Scan(&n)
		studentNumber[n] = false
	}

	// 빈 result리스트에 studentNumber가 true면 추가
	var result []int
	for i := 1; i <= 30; i++ {
		if studentNumber[i] {

			// result라는 리스트에 i를 추가
			result = append(result, i)
		}
	}

	// 오름차순 정렬
	sort.Ints(result)

	// 한 줄 씩 출력
	fmt.Println(result[0])
	fmt.Println(result[1])
}
