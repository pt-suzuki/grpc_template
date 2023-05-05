package user

type UseCase interface {
}

type useCase struct {
	translator Translator
}

func UseCaseImpl(translator Translator) UseCase {
	return &useCase{
		translator: translator,
	}
}
