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

``` go
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
```

테스트 파일을 위와 같이 수정하고 실행해보면 다음과 같은 에러를 볼 수 있습니다.

```
./perimeter_test.go:7:18: not enough arguments in call to Perimeter
        have (Rectangle)
        want (float64, float64)
```

이제 Perimeter와 Area를 Rectangle 구조체를 사용하도록 다음과 같이 수정해봅시다.

```go 
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}
```

함수에 Rectangle을 전달하는 것이 의도에 더 명확하게 전달하지만 앞으로 배워가는 구조체를 사용하는 것이 더 많은 이점이 있다는 것을 알아야 합니다.

다음은 원에 대한 Area 함수를 작성해봅시다.

> 테스트 부터 작성하기

``` go
func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := Area(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
```
`%g` 플래그 같은 경우 실수에서 지수 부분이 크기면 %e로 표현하고 지수가 작으면 %f 로 표현합니다.

> 테스트 실행해보기

테스트를 실행하면 구조체 Circle을 작성하지 않았기 때문에 컴파일 에러가 발생합니다.

```
./perimeter_test.go:27:13: undefined: Circle
```

> 컴파일이 되는 최소한의 코드를 작성하고 테스트 실패 출력을 확인하기

Circle 타입을 정의해봅시다.

``` go
type Circle struct {
	Radius float64
}
```

이제 테스트를 다시 실행해봅시다.

```
./perimeter_test.go:28:14: cannot use circle (type Circle) as type Rectangle in argument to Area
```

일부 프로그래밍 언어를 사용하면 다음과 같은 코드를 사용할 수 있습니다.

``` go
func Area(circle Circle) float64 { ... }
func Area(rectangle Rectangle) float64 { ... }
```

하지만 Go에서는 에러가 발생합니다.

```
./perimeter.go:20:6: Area redeclared in this block
```

이 문제를 해결하기 위해서는 두가지 방법이 있습니다.

- 동일한 이름의 함수를 다른 package로 선언할 수 있습니다. 그래서 새로운 패키지로 Area(Circle)을 만들 수 있지만, 적은 양의 코드를 패키지로 분리하기에는 너무 과한 느낌이 듭니다.

- 새로 정의된 유형을 method를 정의합니다.

> 메서드란 무엇인가?

여태까지 써왔던 t.Errof가 t (testing.T)의 Errorf 메서드입니다.

메서드는 리시버가 있는 함수입니다.

메서드를 선언할 때는 메서드 이름인 식별자를 메서드에 바인딩하고 메서드를 리시버 타입과 연결합니다.

예를 들면 쉽게 이해할 수 있으므로 먼저 테스트를 변경하여 메서드를 호출한 다음 코드를 수정합니다.

```go
func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

```

만약 테스트를 실행한다면

```
./perimeter_test.go:18:19: rectangle.Area undefined (type Rectangle has no field or method Area)
./perimeter_test.go:28:16: circle.Area undefined (type Circle has no field or method Area)
```

여기서 컴파일러가 얼마나 훌륭한지 다시 한번 강조하고 싶습니다.

오류 메시지를 천천히 읽는 것은 매우 중요하기 때문에 장기적으로 도움이 될 것입니다.

읽으면 `type Rectangle has no field or method Area` Rectangle 구조체에 해당 필드나 메서드가 없다고 알려줍니다.

> 컴파일이 되는 최소한의 코드를 작성하고 테스트 실패 출력을 확인하기

다음과 같이 메서드를 추가하겠습니다.

``` go
package structs

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0
}
```

메서드를 선언하는 구문은 함수와 거의 동일하며, 이는 메서드가 함수와 정말 유사하기 때문입니다.

유일한 차이점은 메서드 리시버 구문입니다.

`func (receiverName ReceiverType) MethodName(args)`

메서드가 해당 유형의 변수에 호출되면 receiverName 변수를 통해 해당 데이터에 대한 참조를 얻습니다.

다른 많은 프로그래밍 언어에서 이것은 암묵적으로 수행되며 사용자는 receiverName을 통해 리시버에 접근합니다.

리시버 매개변수를 타입의 첫 번째 문자로 지정하는 것이 Go의 컨벤션입니다.

```
r Rectangle
```

테스트를 다시 실행하려고 하면 값이 달라서 FAIL이 발생합니다.

```
--- FAIL: TestArea (0.00s)
    --- FAIL: TestArea/rectangles (0.00s)
        perimeter_test.go:22: got 0.00 want 72.00
    --- FAIL: TestArea/circles (0.00s)
        perimeter_test.go:32: got 0 want 314.1592653589793
```

> 테스트를 통과하는 최소한의 코드 작성하기

이제 메서드를 원하는 값을 반환할 수 있도록 수정해봅시다.

``` go
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```

```go
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
```

> 리팩터링 하기

테스트에는 몇가지 중복된 것이 있습니다.

우리가 원하는 것은 각종 도형들의 모음을 가져와서 각 도형에 대한 Area()를 호출한 다음 결과를 확인하는 것입니다.

도형인 `Rectagle` 과 `Circle` 테스트를 통과할 수 있지만 도형이 아닌 것을 전달하려고 하면 컴파일하지 못하도록 하고 싶습니다.

Go를 사용하면 interface를 통해 목적을 달성할 수 있습니다.

인터페이스는 Go와 같이 정적 형식의 언어에서는 매우 강력한 개념으로 다양한 타입과 함께 사용할 수 있는 함수를 만들고 여전히 type-safely를 유지하면서 고도로 세분화된 코드를 만들 수 있기 때문입니다.

테스트 파일을 리팩토링 해봅시다.

```go
func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
```

다른 연습문제처럼 헬퍼 함수를 만들고 있는데 Shape 타입으로 입력을 받습니다.

만약 Shape 타입이 아닌 것이 온다면 컴파일에 실패할 것입니다.

어떤 것이 Shape 타입이 되는지 Go에 다음과 같이 인터페이스로 선언하면 됩니다.

```go
type Shape interface {
	Area() float64
}
```

`Rectangle`과 `Circle`을 만들었던 것 처럼 새로운 type을 만들고 있지만 이번에는 `struct`가 아닌 `interface` 입니다.

인터페이스를 추가하고 테스트를 해보면 테스트가 통과하는 것을 볼 수 있습니다.

> Go의 인터페이스

Go의 인터페이스는 대부분의 다른 프로그래밍 언어의 인터페이스와는 상당히 다릅니다.

다른 언어에서는 보통 `My Type Foo implements interface Bar` 라는 코드를 작성해야합니다.

하지만 우리의 경우에는 해당 메서드가 있는 경우에만 인터페이스를 만족시킬 수 있습니다.

- Rectangle은 Area 메서드를 호출하여 float64를 반환하므로 Shape 인터페이스를 만족시킵니다.

- Circle은 Area 메서드를 호출하여 float64를 반환하므로 Shape 인터페이스를 만족시킵니다.

- string은 해당하는 메서드가 없으므로 인터페이스를 만족하지 않습니다.

Go에서 인터페이스 타입은 암시적입니다.

전달하는 타입이 인터페이스가 요청하는 타입과 일치하면 컴파일 할 수 있습니다.

> 디커플링(Decoupling)

인터페이스가 그 도형이 Rectangle인지 Circle인지 Triangle의 도형인지에 대해 어떻게 신경쓸 필요가 없는지 주목합시다.

인터페이스를 선언함으로써 구체적인 타입으로부터 분리되고 단지 해당 작업을 할 수 있는지(해당 타입에 명시된 메서드가 존재하는지)만 판단합니다.

필요한 것만 선언하는 인터페이스를 사용하는 접근 방식은 소프트웨어 설계에서 매우 중요하며, 이후 섹션에서 자세히 설명하겠습니다.

## 추가 리팩터링 하기

이제 구조체에 대해 어느정도 이해했으므로 "테이블 기반 테스트"를 소개하겠습니다.

https://github.com/golang/go/wiki/TableDrivenTests (테이블 기반 테스트)

테이블 기반 테스트는 동일한 방법으로 테스트 할 수 있는 테스트 사례 목록을 작성하는 경우 유용합니다.

``` go
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
```

여기서 새로운 구문이 하나있는데, 바로 익명구조체를 만드는 것입니다.

익명구조체 areaTests는 shape와 want라는 두개의 필드가 존재하는 `[]struct`를 사용하여 구조체를 선언합니다.

그런 다음 구조체 슬라이스를 테스트 케이스로 채웁니다.

다른 슬라이스 처럼 테스트를 실행하기 위해 구조체 필드를 사용하여 반복합니다.

개발자가 새로운 도형을 도입하고 Area를 구현한 후 테스트 케이스에 추가하는 것은 쉽습니다.

게다가 만약 Area에 있는 버그가 발견된다면, 버그를 고치기 전에 새로운 테스트 케이스를 추가하여 테스트하는 것이 매우 쉽습니다.

## 테스트부터 작성하기

새로운 도형을 위한 테스트를 추가하는 것은 매우 쉽습니다.

코드 `{Triangle{12, 6}, 36.0},`를 추가하기만 하면 됩니다.

## 테스트 실행해보기

테스트를 실행해봅시다.

`./perimeter_test.go:22:4: undefined: Triangle`

다음과 같이 에러가 발생합니다.

### 컴파일이 되는 최소한의 코드를 작성하고, 테스트 실패 출력을 확인하기

Triangle 타입을 정의해봅시다.

``` go
type Triangle struct {
	Base   float64
	Height float64
}
```

다시 테스트를 해봅시다.

```
./perimeter_test.go:22:12: cannot use Triangle{...} (type Triangle) as type Shape in field value:
        Triangle does not implement Shape (missing Area method)
```

Area 메서드가 Triangle에 없기때문에 Shape interface에서 도형으로 사용할 수 없음을 알려주고 있으므로 빈 메서드를 추가하여 테스트 해봅시다.

```
func (t Triangle) Area() float64 {
	return 0
}
```

테스트를 실행해보면 코드가 실행되지만 값이 틀려 FAIL이 발생합니다.

```
--- FAIL: TestArea (0.00s)
    perimeter_test.go:27: got 0 want 36
```

## 테스트를 통과하는 최소한의 코드 작성하기

```go
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
```

# 리팩터링 하기

테스트 코드를 약간 개선해봅시다.

익명 구조체에 값을 집어넣을 때 현재는 다음과 같이 집어 넣고 있습니다.

``` go
	{Rectangle{12, 6}, 72.0},
	{Circle{10}, 314.1592653589793},
	{Triangle{12, 6}, 36.0},
```

모든 숫자가 무엇을 나타내는지 바로 분명하지 않으며 쉽게 이해될수 있는 테스트를 목표로 만들어야 합니다.

지금까지 `MyStruct{val1, val2}` 구조의 인스턴스를 생성하는 구문만 표시되었지만 선택적으로 필드 이름을 지정할 수 있습니다.

다음과 같이 수정해봅시다.

```go
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
```

각 필드가 어떤 역할인지 좀 더 보기 쉬워졌습니다.

## 테스트 출력이 유용한지 확인하기

아까 Triangle을 실행하다가 실패한 테스트에서 `perimeter_test.go:27: got 0 want 36` 와 같이 출력 됐었습니다.

그때는  Triangle을 가지고 작업하고 있기 때문에 Triangle과 관련 있다는 것을 알 수 있었습니다.

하지만 테이블에 있는 20가지 테스트 케이스 중 하나에서 버그가 발생한다면 어떤 경우에 실패했는지 개발자가 어떻게 알 수 있을까요?

먼저, 실제로 실폐한 테스트 케이스를 찾기 위해 어떤 케이스에서 실패했는지 확인해야 합니다.

오류 메시지를 `%#v got %g want %g`로 변경해봅시다.

`%#v` 형식 문자열은 필드 값이 있는 구조를 출력하여 개발자가 테스트 중인 속성을 한눈에 볼 수 있도록 합니다.

테스트 사례의 가독성을 높이기 위해 want 필드를 hasArea와 같이 좀 더 의미있는 이름으로 바꿀 수 있습니다.

테이블 기반 테스트의 마지막 팁은 t.Run을 사용하고 테스트 케이스 이름을 지정하는 것입니다.

각각의 케이스를 t.Run으로 실행하면 케이스 이름을 출력하기 때문에 실패시 테스트 출력이 명확해집니다.

또한 go test -run TestArea/Rectangle 과 같이 사용하여 특정 테스트만 실행할 수도 있습니다.

위의 팁들을 적용하여 작성하면 최종적으로 다음과 같은 코드를 작성할 수 있습니다.

``` go
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %g want %g", got, tt.hasArea)
			}
		})
	}
}
```


