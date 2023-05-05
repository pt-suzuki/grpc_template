package user

type Translator interface {
}

type translator struct {
}

func TranslatorImpl() Translator {
	return &translator{}
}
