// 공 바꾸기

package main

import "fmt"

func main() {

	// N, M 변수 선언
	var N, M int

	// 바구니 개수(N), 교환횟수(M) 입력받기
	fmt.Scan(&N, &M)

	// 리스트 초기화
	// 초기 상태는 [0, ... , 0] (N개)
	basket := make([]int, N)

	// i가 0부터 시작해서 N번 반복
	// i:=0 i변수 선언과 함께 0 할당
	// i<N 일 때만 반복
	// i++는 i = i+1과 동일
	// [1, 2, ... , N]
	for i := 0; i < N; i++ {
		basket[i] = i + 1
	}

	// M번 입력을 받음
	// swap 연산
	for i := 0; i < M; i++ {
		var i, j int
		fmt.Scan(&i, &j)
		basket[i-1], basket[j-1] = basket[j-1], basket[i-1]
	}

	// basket 내부의 N개 원소를 모두 출력
	// 형식 문자열을 이용해 출력
	// "%d " 공백을 넣어서 출력
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", basket[i])
	}
}
