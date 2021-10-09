package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

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

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		define := "this is just a test"
		dict := Dict{word: define}
		newDefine := "new define"
		err := dict.Update(word, newDefine)
		assertError(t, err, nil)
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

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dict{word: "test define"}
	dict.Delete(word)

	_, err := dict.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
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

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
