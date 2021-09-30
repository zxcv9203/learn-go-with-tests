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

``` go
func Perimeter(width float64, height float64) float64 {
	return 0
}
```

> 테스트를 통과하는 최소한의 코드 작성하기

``` go
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}
```

지금까지는 정말 쉽게 해결되었습니다.

이제 직사각형의 면적을 반환하는 `Area(width, height float64)` 함수를 만들어 봅시다.

위에서 했던 것처럼 TDD 주기에 따라 직접 작성해봅시다.

``` go
func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
```

``` go
func Area(width, height float64) float64 {
	return width * height
}
```

> 리팩터링 하기

코드는 제대로 작동하지만, 직사각형에 대한 명시적인 내용이 없습니다.

부주의한 개발자가 삼각형의 너비와 높이를 이러한 함수에 사용할 수 있는데, 함수가 잘못된 값을 반환할 것입니다.

이러한 상황을 막기위해 RectangleArea와 같이 기능을 좀 더 구체적으로 지정할 수 있습니다.

더 나은 해결책은 Rectangle이라는 사용자지정 type을 정의하여 캡슐화하는 것입니다.

struct 키워드를 사용해서 간단한 유형을 만들 수 있습니다.

struct는 데이터를 저장할 수 있는 명명도니 필드의 집합입니다.

struct는 다음과 같이 선언할 수 있습니다.

```go
type Rectangle struct {
	Width float64
	Height float64
}
```

이제 float64 대신 Rectangle을 사용하도록 코드를 리팩터링 해봅시다.