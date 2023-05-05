package rock_peper_scissors

type ReportUseCase interface {
	GetByUserId(userId string) (*Report, error)
}

type reportUseCase struct {
	translator ReportTranslator
	useCase    UseCase
}

func ReportUseCaseImpl(translator ReportTranslator, useCase UseCase) ReportUseCase {
	return &reportUseCase{
		translator: translator,
		useCase:    useCase,
	}
}

func (u *reportUseCase) GetByUserId(userId string) (*Report, error) {
	criteria := &SearchCriteria{
		UserID: userId,
	}

	results, err := u.useCase.ListByCriteria(criteria)
	if err != nil {
		return nil, err
	}

	m := &Report{
		MatchResults: results,
	}
	m.CalcGameCount()

	return m, nil
}
