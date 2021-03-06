# 의존성 주입

의존성 주입을 위해서는 인터페이스에 대한 이해가 필요합니다.

프로그래밍 커뮤니티에는 의존성 주입과 관련해서 많은 오해가 있습니다.

이 가이드에서는 다음 항목들이 당신에게 어떻게 가능한지 알려줄 것입니다.

- 프레임워크가 필요하지 않습니다.

- 디자인을 지나치게 복잡하게 하지 않습니다.

- 테스트를 용이하게합니다.

- 그것이 훌륭한 범용함수를 작성하게 할 것입니다.

누군가에게 인사하는 함수를 작성하고 싶다고 가정해봅시다.

함수는 다음과 같은 형태를 가지게 될 것입니다.

```go
func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}
```

fmt.Printf를 호출하면 stdout으로 출력됩니다.

이는 테스트 프레임워크를 사용하여 캡처하기 매우 어렵습니다.

우리가 해야할 일은 print 하는 것의 의존성을 주입(인자를 넘기는 것을 표현한 것입니다.)할 수 있도록 합니다.

우리의 함수는 어디에서 또는 어떻게 print가 발생하는 지를 신경 쓸 필요 없어야합니다.

그래서 좀 더 구체적인 type보다는 interface type을 허용해야 합니다.

그렇게 한다면 우리가 제어하는 어떤 것을 출력하도록 구현을 변경하여 테스트할 수 있습니다.

"실생황"에서는 stdout에 쓰는것을 주입합니다.

fmt.Printf의 소스 코드를 보면 연결하는 방법을 알 수 있습니다.

```go
func Printf(format string, a ...interface{}) (n int, err error) {    
	return Fprintf(os.Stdout, format, a...)
	}
```

내부적으로 Printf는 os.Stdout을 전달하는 Fprintf를 호출할 뿐입니다.

os.Stdout은 정확히 무엇일까요? Fprintf는 첫 번째 인자로 무엇을 전달받을 것으로 예상할까요?

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {    
	p := newPrinter()    
	p.doPrintf(format, a)    
	n, err = w.Write(p.buf)   
	p.free()    
	return
}
```

바로 io.writer 입니다.

```go
type Writer interface {    
	Write(p []byte) (n int, err error)
}
```

당신이 Go 코드를 더 많이 작성하면 이 interface를 많이 보게 될 것입니다.

왜냐하면 data를 어딘가에 넣는 것을 잘 표현하는 좋은 general purpose interface이기 때문입니다.

그래서 우리는 어딘가에 인사말을 보내기 위해 궁극적으로 Writer를 사용하고 있다는 것을 압니다.

기존 추상화를 사용하여 코드를 테스트할 수 있고 더 재사용 가능하게 만들어 봅시다.

## 테스트 먼저 만들기

```
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

bytes 패키지의 buffer 타입은 Writer 인터페이스를 구현합니다.

그래서 우리는 테스트에서 buffer를 Writer로서 사용할 것입니다.

그리고 Greet 호출한 후에 무엇이 buffer에 쓰여졌는지 확인할 수 있습니다.

## 테스트 실행해보기

```
./greet_test.go:10:2: undefined: Greet
```

테스트는 컴파일 되지 않습니다.

## 테스트 실행을 위해 최소한의 코드를 작성하고 실패한 테스트 출력을 확인하기

```go
func Greet(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}
```
테스트를 실행해보면 stdout으로 출력이 되어서 다음과 같이 실패합니다.

```
Hello, Chris--- FAIL: TestGreet (0.00s)
    greet_test.go:16: got "" want "Hello, Chris"
```

## 테스트를 통과하는 최소한의 코드 작성하기

테스트에서 writer를 사용하여 버퍼에 문자열을 보냅니다.

fmt.Fprintf는 fmt.Printf와 비슷하지만 문자열을 보내는 곳은 writer가 가리키는 곳이 됩니다.

반면, fmt.Printf는 기본적으로 stdout으로 보냅니다.

```go
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
```

이제 테스트가 통과되는 것을 볼 수 있습니다.

## 리팩터링 하기

이전에 컴파일러는 bytes.Buffer에 대한 포인터를 전달하라고 했습니다.

기술적으로는 정확하지만, 그다지 유용하지는 않습니다.

이를 증명하기 위해 Greet 함수를 표준 출력으로 인쇄하려는 Go 애플리케이션에 연결해봅시다.

```go
func main() {
	Greet(os.Stdout, "Elodie")
}
```

```
cannot use os.Stdout (type *os.File) as type *bytes.Buffer in argument to Greet
```
에러가 발생하는 것을 볼 수 있습니다.

io.Writer로 os.Stdout과 bytes.Buffer를 둘다 전달 받을 수 있도록 해봅시다.

좀 더 범용적인 인터페이스를 사용하도록 코드를 변경하면 이제 테스트와 애플리케이션 모두에서 사용할 수 있습니다.

```go
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
```

## io.Writer에 대해 알아보기

io.Writer를 사용하여 데이터를 쓸 수 있는 다른곳이 있을까요?

Greet 함수는 얼마나 general purpose 한가요?

> The internet

다음을 실행해봅시다.

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

```

프로그램을 실행시키고 http://localhost:5000 주소로 가보면 Greet 함수가 잘 사용되는 것을 볼 수 있습니다.

HTTP 서버는 이후에 다룰 것이므로 세부사항에 대해서 자세히 알지 못해도 괜찮습니다.

HTTP 핸들러를 작성할 때 요청을 위해 만들어진 http.ResponseWriter와 http.Request를 받습니다.

서버를 구현할 때 writer를 사용하여 응답을 write 합니다.

http.ResponseWriter 또한 io.Writer를 구현한다고 추측할 것이고 그것이 우리가 handler 내에서 Greet 함수를 재사용할 수 있는 이유입니다.

# 정리

첫번째 코드는 제어할 수 없는 곳(stdout)에 데이터를 기록했기 때문에 테스트하기가 쉽지 않았었습니다.

테스트를 통해 코드를 리팩토링하고 어느곳에 데이터가 쓰여질지를 종속성 주입을 통해 제어할 수 있게 되었습니다.

> 코드를 테스트해라

함수를 쉽게 테스트할 수 없다면, 이는 일반적으로 함수 또는 전역 상태에 연결된 종속성 때문입니다.

예를들어, 서비스 계층에서 사용되는 글로벌 데이터베이스 연결 풀이 있는 경우 테스트하기가 어려울 수 있으며 실행속도가 느려집니다.

의존성 주입은 당신이 데이터베이스 의존성에(인터페이스를 통해) 주입하도록할 수 있도록 해줍니다.

그런다음 테스트에서 제어할 수 있는 무언가로 mock할 수 있습니다.

> 관심사를 분리해라

데이터가 가는 목적지를 데이터를 어떻게 만들것인가와 분리합니다.

메소드/함수가 너무 많은 책임이 있다고 느낀적이 있다면 (데이터 생성 그리고 db에 쓰기? HTTP 요청 처리 그리고 도메인 레벨 로직 수행?) 의존성 주입이 아마도 필요한 도구일 것입니다.

- 다른 컨텍스트에서 코드를 재사용할 수 있도록 허용하라

우리 코드를 사용할 수 있는 첫 번째 "새로운" 컨텍스트는 테스트 내부입니다.

그러나 향후 누군가가 당신의 함수로 새로운 것을 시도하고 싶다면 그들 자신의 의존성을 주입할 수 있습니다.

## Mocking은 어떻습니가? 의존성 주입을 위해 필요하면 또한 그것은 악마라고 들었습니다.

Mocking은 나중에 자세히 다룰 것입니다. (그리고 악마가 아닙니다.)

Mocking을 사용하면 당신이 주입하는 실제 내용을 가짜 버전으로 대체해서 테스트에서 당신이 제어하고 조사할 수 있도록 합시다.

우리의 경우에는 표준 라이브러리에 사용하도록 어떤 것이 준비되어 있습니다.

## Go 표준 라이브러리는 정말 훌륭합니다. 그것을 공부하세요

io.Writer 인터페이스에 익숙해지려면 테스트에서 bytes.Buffer를 Writer로 사용할 수 있으며 명령 줄 앱 또는 웹 서버에서 표준 라이브러리에 다른 Writer를 사용하여 함수를 사용할 수 있습니다.

표준 라이브러리와 친해질 수록 이러한 범용 인터페이스를 더 많이 볼 수 있습니다.

그러면 당신의 코드에서 그것을 재사용하여 당신의 소프트웨어를 많은 맥락에서 재사용 가능하게 만들 수 있습니다.
