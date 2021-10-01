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

