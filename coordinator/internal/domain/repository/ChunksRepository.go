package repository

type ChunksRepository interface {
	GetChunks() ([]string, error)
}
