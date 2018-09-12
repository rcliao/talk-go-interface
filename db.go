package dao

import "database/sql"

type DBImpl struct {
	db *sql.DB
}

func (d DBImpl) Get(id int) (error, DTO) {
	var (
		i    int
		name string
	)
	err := d.db.QueryRow("SELECT * FROM DTO WHERE id=?", id).Scan(&i, &name)
	if err != nil {
		return err, DTO{}
	}

	return nil, DTO{i, name}
}
