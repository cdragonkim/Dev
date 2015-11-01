# 패키지 만들기
	- GOPATH
	  |__ bin
	  |__ pkg
	  |__ src
	  	   |
	  	   |-- calc
	  	   |    |
	  	   |    |-- sum.go // calc 패키지
	  	   |
	  	   |__ hello
	  	        |
	  	        |_ hello.go // main 패키지

	- 디렉토리 이름과 패키지 명 동일
	- 소스 파일 이름은 패키지 명과 같지 않아도 됨
	  . 소스 파일 첫번째 package calc

	- 패키지 안에서 변수, 상수 이름 규칙
	  1. 첫 글자 영문 소문자 : 패키지내 사용
	  2. 첫 글자 영문 대문자 : 외부에서 사용

	- 기준이 되는 디렉토리는 GOPATH/src

	- 패키지 컴파일하여 라이브러리로 만들려면
	  go install 명령
	  또는 go install <패키지 명> : 위치 상관없음

	- GOPATH/pkg/운영체계_아키텍처 디렉토리에 XX.a 라이브러리 파일 생성