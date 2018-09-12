package cookies

type DTO struct {
	ID   int
	Name string
}

type DAO interface {
	List() (error, []DTO)
	Get(id int) (error, DTO)
	Add(DTO) error
	Update(DTO) error
	Remove(int id) error
}
