# 인터넷 소스 저장소의 패키지 사용하기
	- import : local , internet 패키지 사용
	- 지원 소스저장소와 버전관리 시스템
	  . Github : Git
	  . BitBucket : Mercurial
	  . Launchpad : Bazaar
	  . IBM DevOps Services: Git
	- 설치
	  . $ sudo apt-get install git (우분트)
	  . $ sudo yum install git(CentOS)	  
	- 마지막 데렉토리가 패키지 이름

	- go get 
	  . import로 설정한 패키지 자동으로 받아온다

	- go get <저장소 주소>로도 패키지 받아 옮

	- 저장소 규칙
	  . GitHub
	    github.com/계정/프로젝트
	    github.com/계정/프로젝트/디렉토리
	    github.com/계정/프로젝트/디렉토리/하위 디렉토리
	    