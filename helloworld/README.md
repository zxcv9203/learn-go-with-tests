# Hello, World!

전통적으로 새로운 언어를 배울때 만드는 첫 번째 프로그램은 "Hello, World!"를 출력하는 프로그램을 만듭니다.

먼저 원하는 위치에 폴더를 만들고 `hello.go`라고 불리는 새로운 파일을 만듭시다.

새로 만든 파일에 다음 코드를 작성해봅시다.

``` go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```

위의 파일을 실행하기 위해서 `go run hello.go`를 입력해봅시다.

> 위 코드의 동작 방식

go에서 프로그램을 작성할 때 그 안에 메인 함수안에 정의된 메인 패키지가 있을 것입니다.

패키지는 관련된 go 코드를 그룹화하는 방법입니다.

func 키워드는 함수를 정의하는 방법입니다.

import "fmt"를 사용하여 출력에 사용되는 Println을 가져왔습니다.

> 코드 테스트 해보기

해당 코드가 문제없는지 어떻게 테스트해야 하나요?

일단 우리의 목적은 `Hello, World`를 출력하는 것입니다.

테스트를 하기 위해 `Hello, World` 문자열을 반환하는 함수를 만들어 메인함수와 분리를 먼저 해봅시다.

``` go
func Hello() string {
	return "Hello, World"
}
func main() {
	fmt.Println(Hello())
}
```

이제 위 함수를 테스트를 할 파일을 만듭시다.

hello_test.go 라는 파일을 생성한 뒤 다음과 같이 작성합니다.

``` go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

> Go 모듈?

이제 테스트 파일을 만들었으니 테스트를 한 번 해봅시다.

터미널에서 `go test` 를 입력해봅시다.

만약 golang 1.16 이상의 버전을 받는다면 아래와 같이 에러 메시지가 발생할 수 있습니다.

```
go: cannot find main module; see 'go help modules'
```

문제가 발생한다면 `go mod init hello` 명령을 입력합시다.

```
module hello

go 1.16
```

go.mod 파일은 작성한 코드에 대한 필수적인 정보를 go 도구에 알려줍니다.

응용프로그램을 배포할 계획인 경우 종속성에 대한 정보뿐만 아니라 다운로드 할 수 있는 코드 위치도 포함할 수 있습니다.

현재로서는 모듈 파일이 최소 수준이므로 그대로 둡니다.

모듈에 대한 자세한 내용은 https://golang.org/doc/modules/gomod-ref 에서 확인하세요

앞으로 `go test`나 `go build` 같은 명령어를 실행하기 이전에 항상 새폴더에서 `go mod init <name>`을 실행해야 합니다.

> 테스트로 돌아가기

터미널에서 `go test`를 실행합니다.

```
ok      github.com/zxcv9203/learn-go-with-tests/helloworld      0.001s
```

통과하는 것을 볼 수 있습니다.

이번에는 테스트가 제대로 작동하는지 일부러 문자열을 변경하여 Fail을 발생시켜 보세요

```
--- FAIL: TestHello (0.00s)
    hello_test.go:10: got "Hello, World" want "aHello, World"
```

> go test 문법

go test를 쓰기 위해서는 몇가지 규칙이 존재합니다.

1. `xxx_test.go`와 같이 파일의 이름을 지정해야 합니다.

2. 테스트 함수는 `Testxxx`와 같이 `Test`라는 문구로 시작해야 합니다.

3. 테스트 함수는 `*Testing.t` 이라는 매개변수를 받습니다.

4. `*Testing.t`를 사용하기 위해서는 `testing` 패키지를 `import` 해와야 합니다.

