package dao

type MemoryImpl struct {
	memory map[int]DTO
}

func NewMemoryImpl() MemoryImpl {
	return MemoryImpl{
		memory: make(map[int]DTO),
	}
}

func (m MemoryImpl) Get(id int) (error, DTO) {
	return m.memory[id]
}
