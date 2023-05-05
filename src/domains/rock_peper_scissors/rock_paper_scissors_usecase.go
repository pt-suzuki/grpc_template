package rock_peper_scissors

type UseCase interface {
	PlayGame(handShape HandShapes) (*RockPaperScissor, error)
	ListByCriteria(criteria *SearchCriteria) (RockPaperScissors, error)
}

type useCase struct {
	translator Translator
	repository Repository
}

func UseCaseImpl(translator Translator, repository Repository) UseCase {
	return &useCase{
		translator: translator,
		repository: repository,
	}
}

func (u *useCase) Save(m *RockPaperScissor) error {
	if err := u.repository.Save(m); err != nil {
		return err
	}
	return nil
}

func (u *useCase) ListByCriteria(criteria *SearchCriteria) (RockPaperScissors, error) {
	result, err := u.repository.ListByCriteria(criteria)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *useCase) PlayGame(handShape HandShapes) (*RockPaperScissor, error) {
	result := &RockPaperScissor{
		YourHandShapes: handShape,
	}
	// ジャンケンの勝敗を決めています。
	result.PlayHandShape()

	// 結果を保存しています。
	if err := u.Save(result); err != nil {
		return nil, err
	}
	return result, nil
}
