# 포인터 & 에러

이번에는 간단한 은행 시스템을 만들어 보겠습니다.

Wallet 구조체를 만들고 Bitcoin을 입금해 봅시다.

## 테스트부터 작성하기

```go
package pointer

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
```

이전에는 필드에 접근할 때 직접적으로 필드 이름에 접근했었지만

보안이 중요한 wallet 구조체에서는 구조체 내부의 필드를 밖으로 노출하기 원하지 않습니다.

필드 값을 변경할 때 우리는 메서드를 이용해서 제어하기를 원합니다.

## 테스트 실행해보기

테스트를 실행해보면 Wallet 구조체가 존재하지 않기 때문에 에러가 발생합니다.
```
./wallet_test.go:6:12: undefined: Wallet
```

## 컴파일이 되는 최소한의 코드를 작성하고 테스트 실패 출력을 확인하기

컴파일러는 Wallet이 무엇인지 모르기 때문에 알려줘야 합니다.

``` go
type Wallet struct {}
```

테스트를 실행해보면 필드나 메서드가 존재하지 않는다는 에러 메시지가 나옵니다.

```
./wallet_test.go:8:8: wallet.Deposit undefined (type Wallet has no field or method Deposit)
./wallet_test.go:10:15: wallet.Balance undefined (type Wallet has no field or method Balance)
```

그렇다면 이제 메서드를 정의해 봅시다.

컴파일 에러가 아닌 원하는 값과 달라 FAIL이 뜨는 메시지가 뜨도록 메서드를 추가해봅시다.

```go
func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() int {
	return 0
}
```

다음과 같이 원하는 대로 에러 메세지가 출력됩니다.

```
--- FAIL: TestWallet (0.00s)
    wallet_test.go:14: got 0 want 10
```

## 테스트를 통과하는 최소한의 코드 작성하기

우리의 구조체에 상태를 저장하기 위해 일종의 balance 변수가 필요합니다.

``` go
type Wallet struct {
	balance int
}
```

Go에서는 만약 symbol(변수, 타입, 함수 등)이 소문자로 시작한다면 그것이 정의된 패키지 밖에서는 private 합니다.

우리의 예제에서 우리의 메서드만 이 변수를 조작할 수 있도록 하고 다른 것은 조작하지 못하도록 하길 원합니다.

우리는 내부 balance 필드에 구조체 안에 있는 "receiver" 변수를 통해서만 접근할 수 있다는 것을 기억해야 합니다.

다음과 같이 메서드를 수정해서 테스트 값과 일치할 수 있도록 만듭시다.

``` go
func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
```

> ?????

테스트 문을 실행하면 이상하게도 다음과 같이 값이 변하지 않았습니다.

```
--- FAIL: TestWallet (0.00s)
    wallet_test.go:14: got 0 want 10
```

주의해야 할게 있는데 Go에서는, 함수나 메서드를 호출하는 경우 인자(arguments)는 복사됩니다.

다음 메서드를 호출할 때 `func (w Wallet) Deposit(amout int)`에 w는 메서드를 호출하는 것의 복사본입니다.

너무 컴퓨터 과학적으로 깊게 들어가지 않고 설명해보면, wallet을 선언할 때 wallet은 메모리 어딘가에 저장됩니다.

만약 wallet의 위치를 알고 싶다면 `&wallet` 같은 방식으로 값의 메모리 주소를 찾을 수 있습니다.

한번 다음 문구를 코드에 추가해봅시다.

```go
func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	fmt.Printf("test : %v\n", &wallet.balance)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
```


```go
func (w Wallet) Deposit(amount int) {
	fmt.Printf("Deposit : %v\n", &w.balance)
	w.balance += amount
}
```

두 구조체의 주소 값이 다른 것을 확인할 수 있습니다.

따라서 우리가 만약 코드 안에서 잔액 값을 바꾸어주는 것은 테스트로부터 받은 복사본에 작업하는 것입니다.

결국 테스트에서는 잔액 값이 변화하지 않습니다.

우리는 이것을 포인터로 해결할 수 있습니다.

포인터는 특정한 값을 가리키고 따라서 그 값을 변화시킬 수 있습니다.

따라서 포인터를 사용하면 복사본을 갖지 않고 wallet을 가리키는 포인터를 얻게 되어 값을 바꿀 수 있습니다.

```go
type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
```

리시버 타입의 차이는 Wallet이 아니라 *Wallet이라고 쓰고 이것은 "Wallet에 대한 포인터"라고 얘기할 수 있습니다.

테스트를 실행해보면 통과하는 것을 볼 수 있습니다.

근데 코드를 보면 의문점이 하나 생길 수 있습니다.

포인터를 역참조해서 값을 가져와야 값을 받을 수 있는거 아닌가? 라는 생각이 들 수 있습니다.

``` go
func (w *Wallet) Balance() int {
	return (*w).balance
}
```

위의 코드는 충분히 타당하며, 실제로 잘 실행되는 것을 볼 수 있습니다.

하지만 Go언어의 개발자들은 위 코드의 표기가 쓰기 귀찮다로 생각했고 그래서 Go에서는 특별한 역참조에 대한 명시 없이 사용하는 것을 허용했습니다.

이런 구조체에 대한 포인터는 다음과 같인 구조체 포인터라고 불리고 자동 역참조가 됩니다.

```go
func (w *Wallet) Balance() int {
	return w.balance
}
```

위의 코드를 보면 값을 변경하지 않기 때문에 포인터 리시버를 사용할 필요가 없습니다.

그러나 관습적으로 메서드 리시버 타입을 통일성있게 하나의 타입으로 가져가야합니다.

## 리팩터링 하기

우리는 비트코인 지갑을 만든다고 했지만 비트코인에 대한 언급은 아무데서도 되어있지 않습니다.

우리는 지금까지 int를 사용하였는데 무언가를 세는 데 있어서 int 타입은 좋은 타입이기 때문입니다.

bitcoin을 표현하기 위해 구조체를 생성해 사용하는 것은 좀 과하다고 생각이 들 수 있습니다.

그렇다고 int 만으로 표현하기 에는 동작하는데 문제는 없지만 어떤 값인지에 대해 설명을 잘 못하는 문제가 있습니다.

그런 경우 기존의 존재하는 타입으로 새로운 타입을 만들 수 있습니다.

문법은 다음과 같습니다.

```
type MyName OriginalType
```

``` go
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
```

```go 
func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
```

이렇게 새로운 타입을 만들어 그 타입 위에 메서드들을 정의할 수 있습니다.

이런 방식은 존재하는 타입에서 원하는 특정 도메인에 특화된 기능을 추가하는 경우 유용합니다.

이번에는 Stringer 인터페이스를 구현해봅시다.

``` go
type Stringer interface {
	String() string
}
```

Stringer 인터페이스는 fmt 패키지에 정의되어 있고 Printf에서 %s 포맷의 스트링을 사용하는 경우 타입이 어떻게 출력될지 정의합니다.

``` go
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
```

위에서 보이듯이, 타입 별칭(type alias)에서 새로운 메서드를 생성하는 문법과 구조체에서의 경우가 똑같은 것을 알 수 있습니다.

다음은 우리의 테스트 에서 String()을 사용하도록 포맷 스트링을 바꿔줍니다.

``` go
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
```

한번 실패 결과를 보기 위해 값을 살짝 수정해서 실패해봅시다.

```
    wallet_test.go:15: got 10 BTC want 20 BTC
```

원하는 방식으로 작동하는 것을 알 수 있습니다.

## 테스트부터 작성하기

이제 새로운 기능으로 비트코인의 수를 줄이는 기능을 추가해봅시다.

```go
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
```
## 테스트 실행해보기

```
./wallet_test.go:22:9: wallet.Withdraw undefined (type Wallet has no field or method Withdraw)
```

## 컴파일이 되는 최소한의 코드를 작성하고, 테스트 실패 출력을 확인하기

``` go
func (w *Wallet) Withdraw(amount Bitcoin) {

}

```

```
    --- FAIL: TestWallet/Withdraw (0.00s)
        wallet_test.go:29: got 20 BTC want 10 BTC
```

## 테스트를 통과하는 코드를 작성하기

``` go
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
```

## 리팩터링 하기

우리의 테스트에 중복이 있기 때문에, 중복을 리팩토링하여 제거합니다.

``` go
func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
}
```

또한 테스트해야 할 항목이 또 있는데 만약 Withdraw를 계좌에 남아있는 잔액보다 많이 시도하게 된다면 어떻게 될까요?

지금까지는 초과 인출 시설에 대해서는 가정하지 않았습니다.

Withdraw를 사용하다 문제가 생긴 경우 우리는 어떻게 알려야 할까요?

만약 에러를 알려주길 원한다면 Go에서는 관용적으로 함수에서 리턴 값으로 err을 보내주어 호출자가 확인하고 행동할 수 있도록 해줍니다.

먼저, 테스트부터 작성해봅시다.

## 테스트부터 작성하기

```go
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
```

만약 기존의 잔액보다 더 많이 인출을 시도하려고 한다면 잔액은 기존과 같게 유지하고 Withdraw 에서는 에러를 리턴하도록 해야합니다.

그래서 우리는 만약 nil이 리턴된다면 테스트가 실패하도록 하여 에러를 체크할 것입니다.

nil은 다른 프로그래밍 언어에서의 null과 동의어입니다.

에러는 nil이 될 수 있는데, 그 이유는 Withdraw의 리턴 타입이 error 이고 이것은 인터페이스이기 때문입니다.

만약 인터페이스를 인자나 리턴 값으로 받는 함수를 보게 되면 이것은 nil이 될 수 있습니다. (nillable)

null 처럼 만약 nil 값에 접근하려 하면 런타임 패닉을 발생시킵니다.

이것은 매우 위험하니 반드시 nil인지 확인해야 합니다.

## 테스트 실행해보기

```
./wallet_test.go:29:25: wallet.Withdraw(Bitcoin(100)) used as value
```

위의 말이 아마 좀 확실하지 않아 보일 수 있지만 이전의 Withdraw의 의도는 단지 호출하는 것이였고, 값을 리턴하지 않았었습니다.

컴파일을 하기 위해서 반환 타입에 error를 추가해야 합니다.

## 컴파일이 되는 최소한의 코드를 작성하고 테스트 실패 출력 확인하기

```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.balance -= amount
	return nil
}
```

일단 컴파일러를 만족시키기 위해 error를 리턴하도록 하고 nil을 무조건 리턴하게 만듭시다.

## 테스트를 통과하는 최소한의 코드 작성하기

```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}
```

코드에서 errors를 import 해주는 것을 기억해야 합니다.

errors.New는 당신이 작성한 메시지와 함께 새로운 error를 생성하여 줍니다.

## 리팩터링 하기

에러 체크를 하는 데 있어 테스트를 좀 더 명확하게 읽을 수 있도록 테스트 헬퍼 함수를 만들어 줍니다.

```go
	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}
```

그리고 테스트 코드를 다음과 같이 수정합시다.

```go
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})
```

"oh no"라는 에러를 리턴하는 것은 별로 유용하지 않기에 우리는 계속 에러를 반복할 수 있다는 것을 생각해야 합니다.

에러가 궁극적으로 유저에게 전달된다고 가정하면, 단지 에러가 존재하게 두는 것보다는 테스트에서 어떤 종류의 메시지라도 assert 하도록 개선해야 합니다.

## 테스트 부터 작성하기

helper에서 string을 비교하도록 업데이트 합니다.

```go
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
```

그 다음 테스트 코드를 수정합니다.

``` go
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot with draw, insufficient funds")
	})
```

위에서는 t.Fatal은 호출되면 해당 시점에서 테스트를 중지하게 됩니다.

만약 t.Error로 할 경우 이후의 것도 실행되어 nil 포인터 접근으로 인한 패닉이 발생합니다.

## 테스트 실행해보기

```
    --- FAIL: TestWallet/Withdraw_insufficient_funds (0.00s)
        wallet_test.go:40: got "oh no", want "cannot with draw, insufficient funds"
```

## 테스트를 통과하는 최소한의 코드를 작성하기

```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.balance -= amount
	return nil
}
```

## 리팩터링 하기

테스트 코스와 Withdraw 코드 모두 에러 메시지를 포함하고 있어 중복이 있습니다.

누군가가 테스트 메시지를 바꾸길 원한다면 테스트를 실패하도록 하는 것은 매우 귀찮은 일이 될 것이고 메시지의 내용을 바꾸는 것은 테스트에서 너무 디테일한 부분입니다.

우리는 정확히 어떤 단어인지 관심이 없고, 특정한 상황에서 일종의 의미있는 에러 메시지를 반환해주면 됩니다.

Go에서 에러는 값이기 때문에 우리는 에러를 변수로 리팩토링 할 수 있어 하나의 값으로 에러를 가지고 갈 수 있습니다.

```go
var ErrInsfficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsfficientFunds
	}
	w.balance -= amount
	return nil
}
```

var 키워드는 패키지에서 전역으로 변수를 선언할 수 있도록 허용합니다.

이제 Withdraw 함수는 매우 깔끔해졌기 때문에 이것은 그 자체로 매우 긍정적인 변화입니다.

다음은 테스트 코드 대신 특정한 스트링을 사용하는 대신 이 값을 사용하도록 리팩토링 합니다.

``` go
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsfficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

```

헬퍼 함수를 메인 테스트 함수에서 빼서 옮겼고 따라서 다른 사람들이 파일들을 열었을 때 헬퍼 함수를 먼저 읽기보다 테스트 코드를 먼저 읽을 수 있도록 만들었습니다.

테스트의 다른 특징으로는 테스트를 통하여 우리가 실제 코드의 사용법을 이해하도록 도와주고 그래서 코드를 공감할 수 있도록 만들어줍니다.

여기서 보듯이 개발자는 간단히 우리의 코드를 호출하고 ErrInsufficientFunds와 동일한지 확인하고 적절히 행동하면 됩니다.

> 확인하지 않은 에러

Go 컴파일러가 많은 것을 도와주지만, 때때로 당신이 놓치고 에러 핸들링 하기 쉽지 않은 것들이 있습니다.

우리가 테스트하지 않은 하나의 시나리오가 있습니다.

그것을 찾기 위해, 터미널에서 다음과 같은 것을 쳐서 Go에서 이용 가능한 linter 중 하나인 errcheck를 설치해봅시다.

`go get -u github.com/kisielk/errcheck`

그 뒤 코드가 있는 디렉터리 안에서 `errcheck .`을 실행해봅시다.

다음과 같은 결과가 출력될 것입니다.

```
wallet_test.go:15:18:   wallet.Withdraw(Bitcoin(10))
```

위에서 우리에게 말하고자 하는 것은 우리는 코드의 그 줄에서 반환되는 에러를 확인하지 않았다는 것입니다.

즉, Withdraw 함수가 성공적으로 실행되었는지 에러가 반환되지 않았는지 확인하지 않았다는 것을 의미합니다.

마지막으로 테스트 코드를 수정해봅시다.

``` go

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsfficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

```

## 정리

> 포인터

- Go는 함수 / 메서드에서 값을 넘겨줄때 값을 복사하기 때문에 만약 다른 함수에서 값을 바꾸길 원한다면 바꾸기 원하는 것을 포인터로 받아야 합니다.

- Go에서 값을 복사한다는 사실은 꽤 자주 유용하지만 때때로 당신의 시스템에서 어떤 것의 복사본을 만들기 원하지 않는다면 그 경우, 참조(reference)를 넘겨주어야 합니다. 예를들면, 매우 큰 데이터나 데이터베이스의 커넥션 풀이 아마 당신이 하나의 인스턴스만 가지려 하는 것들일 수 있습니다.

> nil

- 포인터는 nil일 수 있습니다.
- 만약 함수가 어떤것의 포인터를 반환하였다면 당신은 반드시 그것이 nil인지 아닌지 확인하거나 런타임 예외를 일으켜야 합니다. 컴파일러는 이런 상황에서 에러를 잡아내지 못합니다.
- 당신이 표현하려 하는 값이 없을 수도 있을 때 유용합니다.

> 에러

- 에러는 함수 / 메서드를 호출할 때 실패를 알려주는 방법입니다.

- 우리의 테스트 과정을 본다면, 에러에 string을 사용하여 체크하는 방법은 매우 유별난 테스트가 된다고 결론을 내렸습니다. 따라서 우리는 그 대신 의미 있는 값으로 리팩토링하여 테스트를 더 쉽게 할 수 있고 이렇게 사용하면 API의 사용자도 더 쉬워질 것입니다.

- 이것은 에러 처리의 끝이 아니며, 당신은 좀 더 복잡한 것을 할 수 있고 이것은 단지 시작이다. 이후 섹션에서 더 많은 전략을 다룰 것입니다.

- 에러를 체크하지마라, 에러를 우아하게 다뤄라

> 기존 타입으로 부터 새로운 타입 생성

- 값에 특정 도메인에 의미를 추가하는 데 유용합니다.

- 인터페이스를 구현할 수 있도록 합니다.

포인터와 에러는 Go를 작성하는 데 있어서 매우 큰 부분이며 당신은 이것들에 익숙해져야 합니다.

만약 당신이 실수하더라도 고맙게도 컴파일러가 보통 문제가 생긴부분을 도와주기 때문에 시간을 들여 출력된 에러 내용을 읽어봅시다!