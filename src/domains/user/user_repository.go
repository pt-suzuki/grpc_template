package user

type Repository interface {
}

type repository struct {
	translator Translator
}

func RepositoryImpl(translator Translator) Repository {
	return &repository{
		translator: translator,
	}
}
