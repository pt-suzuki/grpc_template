package rock_peper_scissors

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Save(m *RockPaperScissor) error
	ListByCriteria(criteria *SearchCriteria) (RockPaperScissors, error)
}

type repository struct {
	translator Translator
	db         *gorm.DB
}

func RepositoryImpl(translator Translator, db *gorm.DB) Repository {
	return &repository{
		translator: translator,
		db:         db,
	}
}

func (r *repository) ListByCriteria(criteria *SearchCriteria) (RockPaperScissors, error) {
	var data []*RockPaperScissorGorm

	if err := r.db.
		Where("user_id = ?", criteria.UserID).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return r.translator.MappersToModels(data), nil
}

// Save oauth_results.
func (r *repository) Save(m *RockPaperScissor) error {
	if len(m.ID) > 0 {
		return r.update(m)
	}
	return r.create(m)
}

func (r *repository) create(m *RockPaperScissor) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	data := r.translator.ModelToMapper(m)
	data.ID = id.String()
	if err := r.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) update(m *RockPaperScissor) error {
	var data RockPaperScissor
	data.ID = m.ID

	updateData := r.translator.ModelToUpdateInterface(m)

	if err := r.db.Model(data).Updates(updateData).Where("id = ?", data.ID).Error; err != nil {
		return err
	}
	return nil
}
