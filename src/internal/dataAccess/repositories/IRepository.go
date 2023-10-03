package repositories

/*type IDB interface {
	DB() (*sql.DB, error)
}*/

type IRepository interface {
	Delete(Entity interface{}) error
	SelectById(id uint64) (interface{}, bool)
}
