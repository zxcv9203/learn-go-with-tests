# 배열과 슬라이스

배열은 같은 타입의 여러 요소들을 한 변수에 특정한 순서대로 저장할 수 있게 해줍니다.

우리가 배열을 다룰 때에, 배열들의 요소들을 탐색하며 사용하는 것은 매우 흔한 일입니다.

그래서 우리가 이전에 배운 `for` 키워드를 사용하여 배열의 요소들을 더하는 `Sum` 함수를 만들 것입니다.

`Sum` 함수는 숫자 배열들을 가지고 그 숫자들의 총합을 반환해줄 것입니다.

프로그램을 만들면서 우리는 TDD(테스트 기반 설꼐) 스킬을 활용할 것입니다.

> 테스트를 먼저 작성해봅시다

먼저 테스트 파일을 만들어봅시다.
```go
package arrslice

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
```

배열은 변수로 선언될때 우리가 정한 고정된 크기를 가집니다.

배열을 선언할 때 다음과 같은 방법을 사용해서 선언할 수 있습니다.

```go 
numbers := [5]int{1, 2, 3, 4, 5}      // [N]type{value1, value2, ..., valueN}
numbers := [...]int{1, 2, 3, 4, 5}   // [...]type{value1, value2, ..., valueN}
```

> 테스트를 실행해보기

`go test`로 실행을 시키면 `Sum` 함수를 아직 만들지 않았기 때문에 에러를 발생 시킵니다.

`./sum_test.go:8:9: undefined: Sum`

> 테스트를 실행할 최소한의 코드를 작성하고 실패한 테스트 출력을 확인합시다.

함수를 다음과 같이 작성하고 다시 `go test`를 해봅시다.

```go
package arrslice

func Sum(numbers [5]int) int {
	return (0)
}
```

위의 코드를 추가하고 `go test`를 하면 다음과 같은 메시지를 볼 수 있습니다.

`sum_test.go:12: got 0 want 15 given, [1 2 3 4 5]`

> 테스트를 통과할 수 있는 충분한 코드를 작성하기

``` go
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
```

배열의 특정 인덱스의 값을 가져오려면 `array[index]`라는 형태로 사용하면 됩니다.

위의 경우에는 numbers를 처음부터 탐색하며 sum에 해당 값을 더할 수 있습니다.

`go test`를 사용하여 테스트를 해보면 문제없이 통과하는 것을 볼 수 있습니다.

> 리팩토링

앞서 작성했던 `Sum` 함수에 `range` 구문을 사용하면 배열의 마지막 까지 쉽게 탐색할 수 있습니다.

``` go
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
```

range 구문을 이용해서 탐색할 경우 `for index, value := range array`와 같은 형태로 사용하게 됩니다.

index에는 배열의 현재 인덱스가 저장되고 value에는 현재 인덱스가 가르키는 값 (`array[i]`)을 의미합니다.

하지만 현재 코드에서는 인덱스 값을 사용할 필요가 없고 Go에서는 선언하고 사용하지 않는 변수가 있으면 에러를 발생시키기 때문에 공백식별자(`_`)를 이용하여 값을 무시해야 합니다.

> 배열 타입의 특성

배열의 흥미로운 특성은, 배열의 크기가 배열의 타입으로 인코딩 된다는 것입니다.

예를들어 `[5]int`를 요구하는 함수에 `[4]int`가 입력될 경우 컴파일 에러가 발생할 것입니다.

`./sum_test.go:8:12: cannot use numbers (type [4]int) as type [5]int in argument to Sum`

이 이유는 바로 서로 다른 타입이기 때문입니다.

`[5]int`를 요구하는 함수에 `[4]int`를 입력하는 행위는 `int`를 요구하는 함수에 `string`을 입력하는 행동과 같습니다.

배열의 크기가 특정 크기만 동작하는 함수를 만들면 확장성이 매우 낮을 것입니다.

이런 문제를 해결하기 위해서 슬라이스를 사용할 수 있습니다.

Go언어에서 슬라이스는 크기를 같이 인코딩하지 않아 아무 크기나 가질 수 있습니다.

이제 슬라이스로 Sum 함수를 만들어 봅시다.

> 슬라이스로 함수를 수정하기전 테스트를 먼저 작성해보기

이제 크기를 자유롭게 설정할 수 있는 `slice` 타입을 사용할 것입니다.

슬라이스의 문법은 배열과 매우 유사합니다.

선언할때 크기를 생략하기만 하면 됩니다.

``` go
myArray := [3]int{1, 2, 3}
mySlice := []int{1, 2, 3}
```

테스트 파일은 다음과 같이 작성합니다.

``` go
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
```

> 테스트를 실행해봅시다.

`go test`로 실행해보면 컴파일이 되지않습니다.

`./sum_test.go:20:13: cannot use numbers (type []int) as type [5]int in argument to Sum`

`[5]int`를 받는 함수에 슬라이스 타입을 인자로 보내서 에러가 발생했습니다.

> 테스트를 실행할 최소한의 코드를 작성하고 실패한 테스트 출력을 확인해봅시다.

여기서 해결할 방법은 둘 중 하나입니다.

- 기존에 존재하던 인수 `numbers [5]int`를 슬라이스로 바꾸는 것입니다. 만약 이방법을 사용한다면 또 다른 테스트가 작동하지 않게 될 것입니다.

- 새로운 함수를 하나 만드는 것입니다.

지금 상황에서는 아무도 우리의 함수를 사용하지 않기 때문에 두 개의 함수를 유지하는 것보다 1개만 가지도록 할 것입니다.

다음과 같이 슬라이스를 받도록 변경합시다.

``` go
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
```

위와 같이 수정했더라도 테스트 파일에서 배열로 보내는 부분을 수정하지 않았기 때문에 아직도 에러가 발생할 것입니다.

테스트 파일을 다음과 같이 수정해 봅시다.

```go
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
```

테스트가 성공적으로 통과하는 것을 볼 수 있습니다.

> 리팩토링

인수를 배열에서 슬라이스로 변경하면서 Sum 함수를 리팩토링을 했습니다.

하지만 실제 작동하는 함수 뿐만 아니라 테스트 파일의 코드도 소홀히 하면 안됩니다.

직접 만든 테스트 값에 대하여 의문을 갖는 것은 상당히 중요합니다.

테스트를 많이 한다고 좋은 것이 아니며, 직접 작성한 코드에 자신감을 가지는 것은 중요합니다.

많은 양의 테스트는 문제를 낳으며 유지할 때 그저 오버헤드를 증가시키만 합니다.

테스트 함수가 하나하나가 비용입니다.

``` go
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
```

지금과 같은 경우 한 함수에 대하여 2개의 테스트를 가지고 있는 것은 매우 불필요합니다.

만약 함수가 한가지 크기의 슬라이스로 통과를 한다면, 이는 다른 크기의 슬라이스도 높은 확률로 통과할 수 있다는 뜻입니다. (단, 적당한 범위 안에서)

Go 언어의 내장된 테스팅 도구는 커버리지 도구가 있습니다.

직접 다루지 않는 부분의 코드를 확인할 수 있도록 돕습니다.

커버리지를 100% 달성하는 것이 목표는 아니며, 어느정도 커버리지가 되는지 알려주기 위함입니다.

코드를 작성하면서 TDD에 관하여 엄격했다면 높은 확률로 당신의 커버리지는 100%로 끝날 것이기 때문입니다.

`go test -cover`로 확인할 수 있습니다.

`coverage: 100.0% of statements`

이제 테스트 하나를 지우고 커버리지를 다시 확인해봅시다.

여전히 커버리지가 100% 인 것을 볼 수 있습니다.

`coverage: 100.0% of statements`

> 새로운 함수 추가하기

이번에는 여러 슬라이스를 받고 슬라이스의 합을 더해서 각각의 값을 반환할 함수를 만들어 보겠습니다.

예를들어, `SumAll([]int{1, 2}, []int[0, 9}` 는 `[]int{3, 9}`를 반환 할 것입니다.

또는, `SumAll([]int{1, 1, 1}`은 `[]int{3}`를 리턴할 것입니다.

> 테스트를 먼저 작성하세요

``` go
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
```

위의 테스트를 실행해보면 `SumAll` 함수가 없어 다음과 같은 에러가 발생합니다.

`./sum_test.go:18:9: undefined: SumAll`

> 테스트가 실행되는 최소한의 코드 작성하기

SumAll을 다음과 같이 작성해봅시다.
``` go
func SumAll(numbersTosum ...[]int) (sums []int) {
	return
}
```

go 언어에서는 인수의 수가 정해지지 않았을 경우 달라지는 개수의 인수를 받을 수 있는 `variadic functions`을 사용할 수 있습니다.

하지만 컴파일을 시도해보면 컴파일이 되지 않습니다.

`./sum_test.go:21:9: invalid operation: got != want (slice can only be compared to nil)`

Go 언어에서는 slice를 다룰 때에 등호를 사용할 수 없습니다.

`got` 슬라이스와 `want` 슬라이스를 반복하여 그들의 값을 비교하는 함수를 만들 수 있습니다.

하지만 이 방법은 편리하지 않습니다.

좀 더 편하게`reflect.DeepEqual`을 사용하면 비교할 수 있습니다.

``` go
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
```

하지만 주의해야 할 점이 있는데 `reflect.DeepEqual`은 `type safe`하지 않다는 것입니다.

코드를 짤때 실수를 했더라도 코드는 그냥 컴파일 될 것입니다.

테스트 코드를 다음과 같이 바꿔봅시다.

``` go
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := "[]int{3, 9}"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
```

위처럼 슬라이스랑 문자열이랑 비교하는데도 컴파일 에러가 발생하지 않습니다.

`reflect.DeepEqual`을 사용하여 슬라이스를 비교하는 것은 매우 편리하지만 사용할 때에 주의해야합니다.

다시 테스트를 정상적으로 되돌리고 `go test`를 실행해보면 다음과 같은 출력 값을 보게 될 것입니다.

`sum_test.go:25: got [] want [3 9]`

> 통과할 수 있도록 충분한 코드를 작성합시다.

우리는 `varages`들을 반복적으로 처리하며 우리의 `Sum` 함수를 사용하여 합을 계산한 후 우리가 리턴할 슬라이스에 추가해야 합니다.

``` go
func SumAll(numbersTosum ...[]int) []int {
	lenOfNumbers := len(numbersTosum)
	sums := make([]int, lenOfNumbers)

	for i, numbers := range numbersTosum {
		sums[i] = Sum(numbers)
	}
	return sums
}
```

위의 코드에서 `make` 함수를 이용하여 전달받은 슬라이스 만큼의 크기를 가진 슬라이스를 생성합니다.

그리고 생성한 슬라이스에 `numberTosum`들을 `Sum` 함수의 인수로 보내고 `Sum` 함수를 호출해서 반환 받은값을 집어넣습니다.

이제 `go test`를 사용하면 통과할 수 있을 것입니다.

> 리팩토링

슬라이스에도 크기가 존재합니다. 만약 크기가 2인 슬라이스에서 `mySlice[10] = 1` 을 시도한다면 런타임 에러가 날 것입니다.

이런 상황을 방지하기 위해 `append` 함수를 이용할 수 있습니다.

``` go
func SumAll(numbersTosum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersTosum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
```

위 처럼 코드를 수정하게 되면 잘못된 인덱스 접근으로 인한 런타임 에러가 발생할 일이 없어집니다.

이번에는 맨앞의 인덱스를 제외하고 계산하는 `SumAllTails` 만들어 봅시다.

> SumAllTails 테스트 작성하기

``` go
func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
```

`go test` 를 실행하면 다음 결과가 나옵니다.

`./sum_test.go:30:9: undefined: SumAllTails`

이제 함수를 작성해봅시다.

> SumAllTails 함수 작성하기

SumAll 함수를 복사해서 SumAllTails 함수로 이름을 변경하고 테스트를 다시 실행해 봅시다.

`sum_test.go:34: got [3 9] want [2 9]`

> 통과할 수 있도록 충분한 코드를 작성합시다.

슬라이스들도 `slice[low:high]`를 통해 잘릴 수 있습니다.

만약 `:` 이후에 값을 입력하지 않는다면 슬라이스의 값 끝을 가르키게 됩니다.

예를들어 1부터 끝까지 가져오라고 표현하고 싶은 경우 `numbers[1:]` 과 같이 표현할 수 있습니다.

``` go
func SumAllTails(numbersTosum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersTosum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
```

> 리팩토링

이번에는 리팩토링을 할 것이 거의 없습니다.

하지만 만약 비어있는 슬라이스를 함수에 넣는다면 어떻게 될까요?

한번 이런 상황을 테스트해봅시다.

> 테스트 작성하기

``` go
func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
```

> 테스트를 실행하기

`panic: runtime error: slice bounds out of range [1:0] [recovered]`

테스트를 실행하면 다음과 같이 런타임 에러가 발생합니다.

이런 상황이 발생하지 않도록 예외처리를 해봅시다

> 통과할수 있도록 코드를 작성하기

만약 받은 매개변수의 길이가 0이라면 0을 반환하도록 합시다.

``` go
func SumAllTails(numbersTosum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersTosum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
```

> 리팩토링 

지금 만든 테스트 파일은 다음과 같이 중복되는 코드가 존재합니다.

```go
func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)      // 중복
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want) // 중복
		}
	})
}
```

이 부분을 따로 함수로 만들어서 중복을 줄여봅시다.

``` go
func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
```

이런식으로 작성하면 중복을 줄일뿐만 아니라 `reflect.DeepEqual`의 문제인 다른 타입끼리의 비교도 막아주어 조금 더 안전하게 코드를 작성할 수 있습니다.

t.Helper()는 해당함수가 헬퍼함수(테스트 코드간에 반복되는 코드를 줄이기 위해 분리하는 함수)임을 알려줍니다.