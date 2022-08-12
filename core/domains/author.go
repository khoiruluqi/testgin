package domains

type Author struct {
	Id				int			`gorm:"primaryKey"`
	Name 			string
	Books	 		[]Book	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}