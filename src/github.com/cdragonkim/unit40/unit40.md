# 문서화 하기
	- // 계산 패키지
	  package calc

	  // 두 정수를 더함
	  func Sum(a int, b int) int {
	  	return a + b
	  }

	- GOPATH로  가서 godoc calc
	  godoc <패키지 명>

	  godoc <패키지 명> <함수명>

	- 웹 브라우저로 볼수 있게 웹서버 실행
	  godoc -http=:6060 
