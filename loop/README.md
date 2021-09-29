# Loop

Go에서 반복되는 작업을 하기 위해서 `for`라는 반복문을 이용하면 깔끔하게 해결할 수 있습니다.

Go에는 `while` `do` `until` 같은 키워드가 없고 오직 `for`만 사용할 수 있습니다.

이번에는 문자를 5번 반복하는 함수를 위한 테스트를 먼저 만들어보겠습니다.

> 테스트 파일 작성
``` go
package loop

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
```

`go test`로 테스트시 역시 `Repeat` 함수가 존재하지 않아 실패하는 것을 볼 수 있습니다.

`./loop_test.go:6:14: undefined: Repeat`

> 테스트를 실행하기 위한 약간의 코드를 작성하고 실패하는 테스트 출력을 확인하기

테스트 파일이 컴파일이 될 수 있도록 Repeat 함수를 만들어 봅시다.

``` go
package loop

func Repeat(str string) string {
	return ""
}
```

`--- FAIL: TestRepeat (0.00s)`

`loop_test.go:10: expected "aaaaa" but got ""`

`go test` 를 사용해서 값을 비교했을 때 위와같이 값이 달라서 테스트가 실패하는 것을 볼 수 있습니다.

> 테스트를 통과할 수 있는 충분한 코드를 작성하기

Go의 `for` 문법은 대체로 C와 비슷한 언어들이 따르는 매우 평범한 형태입니다.

다음과 같이 for를 사용하여 코드를 수정해 봅시다.

``` go
func Repeat(str string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + str
	}
	return repeated
}
```

`C` `Java` `JavaScript`와 같은 언어들과 달리 세 개의 컴포넌트(`i := 0; i < 5; i++`)를 둘러싼 괄호가 없고 중괄호 `{}`가 항상 필요합니다.

`go test`로 테스트 파일을 실행시켜보면 잘 실행되는 것을 볼 수 있습니다.

> 단축 변수선언

위의 코드에서 우리는 변수를 다음과 같이 선언했습니다.

``` go
var repeated string
```

변수를 초기화하고 선언할때 `:=`를 사용할 수 있는데 선언과 같이 값 초기화를 하는 과정을 한번에 할 수 있게 해줍니다.

하지만 여기서는 선언만하기 때문에 명시적으로 `var repeated string`과 같이 선언했습니다.

만약 `:=`으로 선언하고 싶다면 다음과 같이 할 수 있습니다.

```go
repeated := ""
```

이런 식으로 `:=`를 이용하여 변수를 선언과 동시에 초기화하는 방법을 단축변수선언이라고 합니다. 

> 코드 리팩토링

코드에 불 필요한 부분이나 보기 쉽게 하기 위해 다음과 같이 `Repeat` 함수를 수정합니다.

``` go
package loop

const repeatCount = 5

func Repeat(str string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += str
	}
	return repeated
}
```

위의 코드에서 5를 상수 `repeatCount`로 변경하고 `+=` 연산자로 값을 더했습니다.

`i < 5` 부분을 `i < repeatCount`로 변경하면서 반복하는 수의 의미를 더 명확하게 알 수 있습니다.

`repeated = repeated + str`를 `repeated += str` 이라고 변경하면서 값을 계속 붙이는 부분을 좀 더 간단하게 표현할 수 있습니다.

> 성능 측정

Go에서 `Benchmark`를 작성하면 테스트할 함수의 성능을 측정할 수 있습니다.

성능 측정을 하기 위해서는 함수의 이름 앞부분을 `Benchmark`라고 지어 주어야 합니다.

다음과 같이 Benchmark 함수를 작성해 봅시다.

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

테스트 함수와 매우 비슷한 형태의 코드를 볼 수 있습니다.

위의 코드는 Repeat 함수를 b.N번 호출합니다.

b.N의 값은 실행할 때 마다 바뀌며 테스트를 할때 b.N의 값을 확인할 수 있습니다.

이제 성능 측정을 위해 `go test -bench=.`(linux or mac 기준) 명령을 입력하여 성능 측정을 해봅시다.

``` text
goos: linux
goarch: amd64
pkg: github.com/zxcv9203/learn-go-with-tests/loop
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkRepeat-12      10222831               118.0 ns/op
PASS
ok      github.com/zxcv9203/learn-go-with-tests/loop    1.330s
```

``` text
BenchmarkRepeat-12      10222831               118.0 ns/op
```

여기서 2번째에 나오는 `10222831`이 b.N(테스트 횟수)의 값입니다.

3번째 값 `118.0 ns/op`은 b.N 만큼 테스트하면서 걸린 평균 작업시간입니다.

> 연습

- 호출자에서 문자가 반복되는 횟수를 지정할 수 있도록 테스트를 변경하고 코드를 수정해봅시다. ✅
- 함수를 문서화하기 위하여 ExampleRepeat를 작성해봅시다. ✅
- strings 패키지를 찾아봅니다. 쓸모 있다고 생각되는 함수를 찾아보고 여기에서 한 것 같이 테스트를 작성해서 실험 해봅니다. ✅