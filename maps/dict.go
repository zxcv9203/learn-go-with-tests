package maps

type Dict map[string]string
type DictErr string

const (
	ErrNotFound         = DictErr("could not find the word you were looking for")
	ErrWordExists       = DictErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictErr("cannot update word because it does not exist")
)

func (e DictErr) Error() string {
	return string(e)
}

func (d Dict) Search(word string) (string, error) {
	define, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return define, nil
}

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

func (d Dict) Delete(word string) {
	delete(d, word)
}
