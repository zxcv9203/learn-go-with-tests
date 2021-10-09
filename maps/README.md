# 맵

배열과 슬라이스에서 값을 순서대로 저장하는 방법을 다뤘습니다.

이번에는 항목을 key에 따라 저장하고 이렇게 저장한 key를 빠르게 찾는 방법을 살펴볼 것이다.

맵은 사전과 비슷한 방식으로 항목을 저장할 수 있어서, key는 단어이고 value는 정의라는 식으로 생각할 수 있습니다.

그러므로 우리만의 사전을 만드는 것이 맵을 배우는 가장 좋은 방법이라고 생각합니다.

우선 몇 개의 단어와 이들의 정의가 있는 사전이 있다고 가정해 봅시다.

단어를 검색하면 사전은 그 단어의 정의를 반환해야 합니다.

## 테스트부터 작성하기

```go
package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "this is just a test"}

	got := Search(dict, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

```

맵을 선언하는 것은 배열을 선언하는 것과 비슷하지만 다음과 같은 점에 있어서 다릅니다.

맵을 선언하려면 map이라는 키워드로 시작하고 두 개의 타입이 있어야 합니다. 첫번째 타입은 키 타입으로 `[]`안에 쓰여야 하고 두번째 타입은 `[]` 다음에 옵니다.

``` go
// map[type]type
map[string]string
```

키 타입에는 오직 비교 가능한 타입만이 올 수 있는데 왜냐하면 두 개의 키가 동일한지 판별할 수 없다면 올바른 값을 가져왔는지 확신할 수 있는 방법이 없기 때문입니다.

반면에 값 타입으로는 무엇이든 원하는 값이 가능합니다, 심지어 또 다른 맵도 가능합니다.

## 테스트 실행해보기

```
./dict_test.go:8:9: undefined: Search
```

## 테스트를 실행할 최소한의 코드를 작성하고 테스트 실패 결과를 확인하기

```go
package maps

func Search(dict map[string]string, word string) string {
	return ""
}
```

실행해보면 명확한 에러 메시지와 함께 실패할 것입니다.

```
--- FAIL: TestSearch (0.00s)
    dict_test.go:12: got "" want "this is just a test" given, "test"
```

## 테스트를 통과하는 최소한의 코드 작성하기

```go
func Search(dict map[string]string, word string) string {
	return dict[word]
}
```

이제 실행해보면 테스트에 통과하는 것을 볼 수 있습니다.

## 리팩터링하기

``` go
func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "this is just a test"}

	got := Search(dict, "test")
	want := "this is just a test"

	assertStrings(t, got, want)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
```

헬퍼 함수를 만들어 테스트가 늘어날 경우 중복된 코드가 증가하지 않도록 분리해주었습니다.

> 커스텀 타입 사용하기

map에 대한 새로운 타입을 만들고 Search 함수를 만들어 의미파악을 좀 더 쉽게 하도록 합니다.

```go
func TestSearch(t *testing.T) {
	dict := dict{"test": "this is just a test"}

	got := dict.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
```

dict 타입이 아직 선언되지 않았으므로 다음과 같이 메인을 수정합니다.

```go
type dict map[string]string

func (d dict) Search(word string) string {
	return d[word]
}
```

## 테스트부터 작성하기

기본 검색은 구현하기 매우쉬웠습니다.

그러나 만약 사전에 없는단어를 검색하면 어떻게 될까요?

실제로 아무것도 가져올 수 없습니다.

이래도 프로그램이 계속 동작하기는 하기때문에 괜찮지만 더 좋은 방법이 있습니다.

바로 Search 함수에서 단어가 사전에 존재하지 않을 경우 존재하지 않는다고 알려주는 것입니다.

이 방법으로 사용자가 단어가 존재하지 않는건지 아니면 단지 정의가 없는건지 궁금해하지 않게 됩니다.

Go에서 이런 상황을 다루는 방법은 두번째 인자로 Error 타입을 활용하는 것입니다.

Error는 .Error() 메서드를 통해 문자열로 변환될 수 있습니다.

만약 Error가 nil일 경우를 대비해서 if 조건문을 통해 nil일 경우는 Error 메서드를 호출하지 않게끔 보호해줘야 합니다.

``` go
func TestSearch(t *testing.T) {
	dict := dict{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertStrings(t, err.Error(), want)
	})
}
```

## 테스트 실행해보기

함수의 반환이 아직 하나이기 때문에 에러가 발생합니다.

```
./dict_test.go:9:10: assignment mismatch: 2 variables but dict.Search returns 1 value
./dict_test.go:15:10: assignment mismatch: 2 variables but dict.Search returns 1 value
```

## 테스트를 실행할 최소한의 코드를 작성하고 테스트 실패 결과를 확인하기

``` go
func (d dict) Search(word string) (string, error) {
	return d[word], nil
}
```

## 테스트를 통과하는 최소한의 코드 작성하기

``` go
package maps

import "errors"

type dict map[string]string

func (d dict) Search(word string) (string, error) {
	define, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return define, nil
}
```

테스트를 통과하기 위해서 맵 탐색의 흥미로운 특성을 사용했습니다.

맵은 두 개의 값을 반환합니다.

두번째 값은 boolean으로 키를 찾는데 성공했는지를 가리킵니다.

이 성질을 이용해서 단어가 존재하지 않는 것과 단어에 정의가 없는 것을 구분할 수 있습니다.

## 리팩터링하기

``` go
package maps

import "errors"

type dict map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d dict) Search(word string) (string, error) {
	define, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return define, nil
}
```

에러 메시지를 따로 분리하면서 코드의 중복을 줄일 수 있습니다.

``` go
	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
```

ErrNotFound라는 변수를 사용함으로써 나중에 에러 문자열을 바꿨을때 하나하나 바꿀 필요가 없어졌습니다.

## 테스트부터 작성하기

사전을 검색하는 훌륭한 방법을 갖추었습니다.

그러나 사전에 새 단어를 추가하는 방법이 없습니다.

```go

func TestAdd(t *testing.T) {
	dict := dict{}
	dict.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dict.Search("test")
	if err != nil {
		t.Fatal("should find added word : ", err)
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

## 테스트 실행해보기

```go
func (d dict) Add(word, define string) {

}
```

```
--- FAIL: TestAdd (0.00s)
    dict_test.go:28: should find added word :  could not find the word you were looking for
FAIL
```

## 테스트를 통과하는 최소한의 코드 작성하기

``` go
func (d dict) Add(word, define string) {
	d[word] = define
}
```

맵에 추가하는 것은 배열과 유사합니다. 키를 명시하고 값을 같게 하면됩니다.

> 포인터, 복사, 그 외

맵의 흥미로운 특성은 주소를 전달(&myMap)하지 않고서도 수정할 수 있다는 것입니다.

"레퍼런스 타입"처럼 느껴질 수 있는데, Dave Cheney가 설명한 바에 따르면 그렇지 않습니다.

맵에 함수/메서드를 전달하게 되면 실제로 맵을 복사하지만 단지 포인터 부분만 해당합니다.

데이터를 가지고 있는 하부 자료 구조는 복사하지 않습니다.

맵에 관해 유의할 점은 nil 값이 가능하다는 점입니다.

읽기 작업을 수행할때 nil 맵은 빈 맵과 동일하게 동작하지만 nil 맵에 쓰기 작업을 시도한다면 이는 런타임 패닉을 일으키는 원인이 됩니다.

https://go.dev/blog/maps (맵에대한 설명)

따라서 절대로 빈 맵을 초기화해서는 안됩니다.

```go
var m map[string]string
```

대신에 다음과 같은 방법으로 빈 맵을 만들 수 있습니다.

``` go
var dict = map[string]string{}
// OR
var dict = make(map[string]string)
```

두 방법은 빈 hash map을 생성하고 dict가 이를 가리키게 합니다.

이것은 절대로 런타임 패닉을 발생하지 않도록 보장하는 방법입니다.

## 리팩터링하기

구현에 리팩터링 할게 그리 많지 않지만 테스트는 보다 간결하게 만들 수 있습니다.

```go
func TestAdd(t *testing.T) {
	dict := Dict{}
	word := "test"
	define := "this is just a test"
	dict.Add("test", "this is just a test")

	assertDefine(t, dict, word, define)
}

func assertDefine(t testing.TB, dict Dict, word, define string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word : ", err)
	}
	if define != got {
		t.Errorf("got %q want %q", got, define)
	}
}
```

단어와 정의를 위한 변수를 만들었고 정의를 검사하는 로직을 별도의 헬퍼 함수로 빼내었습니다.

## 테스트부터 작성하기

Add 함수는 잘 동작하지만 추가하고자 하는 값이 이미 존재하는 경우 값을 덮어 씌워버립니다.

이 점이 편리하게 느껴지는 경우도 있겠지만 함수 이름이 덜 정밀해지는 문제가 있습니다.

Add 함수를 단어가 없을 경우에 뜻을 추가하도록 구현합시다.

``` go
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dict{}
		word := "test"
		define := "this is just a test"
		err := dict.Add("test", "this is just a test")

		assertError(t, err, nil)
		assertDefine(t, dict, word, define)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		define := "this is just a test"
		dict := Dict{word: define}
		err := dict.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefine(t, dict, word, define)
	})
}
```

## 테스트 실행해보기

Add 함수에서 값을 반환하도록 하지 않게 만들었기 때문에 컴파일러는 실패할 것입니다.

```
./dict_test.go:26:18: dict.Add("test", "this is just a test") used as value
./dict_test.go:35:18: dict.Add(word, "new test") used as value
```

## 테스트를 실행할 최소한의 코드를 작성하고 테스트 실패 결과를 확인하기

``` go
var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dict) Add(word, define string) error {
	d[word] = define
	return nil
}
```

## 테스트를 통과하는 최소한의 코드 작성하기

```go
func (d Dict) Add(word, define string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = define
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
```

switch 키워드를 이용해서 에러를 매치해보겠습니다.

이런식으로 switch 구문을 이용하면 더 안전한 코드를 작성할 수 있는데 Search 함수가 ErrNotFound가 아닌 에러를 반환하는 경우가 이에 해당됩니다.

## 리팩터링 하기

리팩터링 할게 그리 많지 않습니다.

그러나 에러 활용도가 커져감에 따라 약간의 수정을 해보겠습니다.

``` go
type DictErr string

const (
	ErrNotFound   = DictErr("could not find the word you were looking for")
	ErrWordExists = DictErr("cannot add word because it already exists")
)

func (e DictErr) Error() string {
	return string(e)
}
```

에러를 상수로 만들었습니다.

이는 error 인터페이스를 구현하는 우리만의 DictErr 타입을 만드는데 필요합니다.

## 테스트부터 작성하기

이번에는 단어의 정의를 Update하는 함수를 만들어 봅시다.

```go
func TestUpdate(t *testing.T) {
	word := "test"
	define := "this is just a test"
	dict := Dict{word: define}
	newDefine := "new define"
	dict.Update(word, newDefine)
	assertDefine(t, dict, word, newDefine)
}
```

## 테스트 실행해보기

```
./dict_test.go:46:6: dict.Update undefined (type Dict has no field or method Update)
```

## 테스트를 실행할 최소한의 코드를 작성하고 테스트 실패 결과를 확인하기

우리는 이와같은 에러를 어떻게 처리해야할 지 이미 알고 있습니다.

다음과 같이 함수를 정의해야 합니다.

```go
func (d Dict) Update(word, define string) {

}
```

테스트를 실행하면 다음과 같은 에러 메시지가 나옵니다.

```
--- FAIL: TestUpdate (0.00s)
    dict_test.go:47: got "this is just a test" want "new define"
```

## 테스트를 통과하는 최소한의 코드 작성하기

``` go
func (d Dict) Update(word, define string) {
	d[word] = define
}
```

현재는 없는 단어가 들어오더라도 새로 생성되기 때문에 없는 단어가 들어오면 에러를 출력하도록 바꿔줍시다.

## 테스트부터 작성하기

```go
func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		define := "this is just a test"
		dict := Dict{word: define}
		newDefine := "new define"
		dict.Update(word, newDefine)
		assertDefine(t, dict, word, newDefine)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		define := "this is just a test"
		dict := Dict{}

		err := dict.Update(word, define)
		assertError(t, err, ErrWordDoesNotExist)
	})
}
```

단어가 존재하지 않는 경우에 관한 또 다른 에러 타입을 추가했습니다.

또한 Update 함수를 수정해서 err 값을 반환받을 수 있도록 구성했습니다.

## 테스트 실행해보기

```
./dict_test.go:47:21: dict.Update(word, newDefine) used as value
./dict_test.go:56:21: dict.Update(word, define) used as value
./dict_test.go:57:23: undefined: ErrWordDoesNotExist
```

## 테스트를 실행할 최소한의 코드를 작성하고 테스트 실패 결과를 확인하기

``` go
const (
	ErrNotFound         = DictErr("could not find the word you were looking for")
	ErrWordExists       = DictErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictErr("cannot update word because it does not exist")
)

func (d Dict) Update(word, define string) error {
	d[word] = define
	return nil
}
```

실행하면 다음과 같은 결과가 나옵니다.

```
    --- FAIL: TestUpdate/new_word (0.00s)
        dict_test.go:57: got error %!q(<nil>) want "cannot update word because it does not exist"
```

## 테스트를 통과하는 최소한의 코드 작성하기

``` go
func (d Dict) Update(word, define string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = define
	default:
		return err
	}
	return nil
}
```

> 업데이트 함수를 위한 새로운 에러 타입을 선언하는 것에 관한 note

ErrNotFound를 사용하고 새로운 에러타입을 추가하지 않을 수도 있습니다.

그러나 업데이트에 실패했을 때 정확한 에러를 받는 것이 종종 더 나을 때가 있습니다.

구체적인 에러는 무엇이 잘못됐는지 더 많은 정보를 줍니다.

다음은 사전에서 단어를 delete 하는 함수를 만들어 봅시다.

## 테스트부터 작성하기

``` go
func TestDelete(t *testing.T) {
	word := "test"
	dict := Dict{word: "test define"}
	dict.Delete(word)

	_, err := dict.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
```

## 테스트 실행해보기

```
./dict_test.go:65:6: dict.Delete undefined (type Dict has no field or method Delete)
```

## 테스트를 실행할 최소한의 코드를 작성하고 테스트 실패 결과를 확인하기

```go
func (d Dict) Delete(word string) {

}
```

```
--- FAIL: TestDelete (0.00s)
    dict_test.go:69: Expected "test" to be deleted
```

## 테스트를 통과하는 최소한의 코드 작성하기

```
func (d Dict) Delete(word string) {
	delete(d, word)
}
```

Go에는 맵에 사용 가능한 delete라는 내장 함수가 있습니다.

이 함수는 두개의 인자를 받습니다.

첫번째 인자는 맵이고 두번째 인자는 삭제할 키입니다.

delete 함수는 아무것도 반환하지 않으며, 우리의 Delete 함수도 같은 방식에 기초할 것입니다.

존재하지 않는 값을 삭제하는 일은 아무런 영향이 없습니다.

Update와 Add 함수와는 달리 API를 에러를 포함해서 복잡하게 할 필요는 없습니다.
