package repository

type chunksRepositoryImpl struct{}

func (c *chunksRepositoryImpl) GetChunks() ([]string, error) {
	panic("not implemented")
}

func NewChunksRepositoryImpl() *chunksRepositoryImpl {
	return &chunksRepositoryImpl{}
}
