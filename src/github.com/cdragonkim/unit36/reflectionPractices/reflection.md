# Reflection 이란?
  - runtime 시점에 Interface , struct 등의 type 정보를 얻어내거나 결정하는 기능
  - Java, C# VM 언어나 Python, Ruby등의 Script언어 주로 사용
  - Go : 기본 패키지에서 제공기

# Reflection으로 구조체 태그 가져오기

# Pointer 와 Interface 값 가져오기
  - 일반 포인터와 인터페이스 값 가져오기
    . 포인터는 값을 가져오려면 
      : reflect.ValueOf 함수로 값 정보reflect.Value 주소 얻어온 후 
        다시 Elem 함수로 값 정보 가져와야 함 
        (값 정보에서 바로 Int함수등으로 값 가져오려면 Runtime Error)
        즉, Elem 함수로 포인터의 메모리에 저장된 실제 값 정보 가져와야 함
        그리고 타입에 맞는 Int, Float, String 함수로 값 가져 옮
      :   

# 동적으로 함수 생성하기
  - reflect.MakeFunc 함수 사용
  - 매개변수와 리턴값은 반드시 []reflect.Value 사용

