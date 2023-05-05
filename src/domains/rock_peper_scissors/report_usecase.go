package rock_peper_scissors

type ReportService interface {
	AddRockPaperScissorPlay(rps *RockPaperScissor)
	GetByUserId(userId string) *Report
}

type reportService struct {
	translator       ReportTranslator
	reportRepository ReportRepository
}

func reportServiceImpl(translator ReportTranslator, reportRepository ReportRepository) ReportService {
	return &reportService{
		translator:       translator,
		reportRepository: reportRepository,
	}
}

func (s *reportService) GetByUserId(userId string) *Report {
	return s.reportRepository.GetByUserId(userId)
}

func (s *reportService) Save(content *Report) {
	s.reportRepository.Save(content)
}

func (s *reportService) AddRockPaperScissorPlay(rps *RockPaperScissor) {
	r := s.GetByUserId(rps.UserId)
	r.AddGameCount(rps)
	s.Save(r)
}
