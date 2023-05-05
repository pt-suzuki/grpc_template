package rock_peper_scissors

import "time"

type RockPaperScissorGorm struct {
	ID                 string
	YourHandShapes     uint32
	OpponentHandShapes uint32
	Result             uint32
	CreateTime         time.Time
	UserID             string
}

// TableName gets table name.
func (*RockPaperScissorGorm) TableName() string {
	return "rock_paper_scissor"
}
