package maps

const (
	ErrUnknownTerm       = DictionaryErr("Unknown term")
	ErrWordAlreadyExists = DictionaryErr("Word Already exists")
)

type Dictionary map[string]string

type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}
func (d Dictionary) Search(term string) (string, error) {
	result, ok := d[term]
	if !ok {
		return "", ErrUnknownTerm
	}
	return result, nil
}

func (d Dictionary) Add(term, meaning string) error {
	_, ok := d[term]
	if !ok {
		d[term] = meaning
		return nil
	}
	return ErrWordAlreadyExists

}

func (d Dictionary) Update(word, newDef string) error {
	_, ok := d[word]

	if ok {
		d[word] = newDef
		return nil
	}
	return ErrUnknownTerm
}
