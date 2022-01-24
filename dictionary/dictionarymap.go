package dictionary

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("no word in the dict")
	ErrWordExists       = DictionaryErr("the word already here")
	ErrWordDoesNotExist = DictionaryErr("Nothing to update")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(s string) (string, error) {
	val, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(s string, s2 string) error {
	if _, ok := d[s]; !ok {
		d[s] = s2
		return nil
	}
	return ErrWordExists

	/*
		_, err := d.Search(word)

		switch err {
		case ErrNotFound:
			d[word] = definition
		case nil:
			return ErrWordExists
		default:
			return err
		}

		return nil
	*/
}

func (d Dictionary) Update(word string, definition string) error {
	if _, ok := d[word]; ok {
		d[word] = definition
		return nil
	}
	return ErrWordDoesNotExist
	/*
		_, err := d.Search(word)

			switch err {
			case ErrNotFound:
				return ErrWordDoesNotExist
			case nil:
				d[word] = definition
			default:
				return err
			}

			return nil
	*/
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
