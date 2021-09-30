# 구조체, 메소드 & 인터페이스

높이와 너비가 주어진 사각형의 둘레를 계산하는 기하학 코드가 필요하다고 가정해봅시다.

해당 코드를 작성하기 위해 다음 과 같은 형태의 함수를 만들 수 있습니다.

`Perimeter(width float64, height float64)`

여기서 `float64`는 `123.45` 같은 부동 소수점 수에 대한 타입입니다.

> 테스트 작성하기

``` go
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
```

> 테스트 실행해보기

``` 
./perimeter_test.go:6:9: undefined: Perimeter
```

> 컴파일이 되는 최소한의 코드를 작성하고, 테스트 실패 출력을 확인하기

