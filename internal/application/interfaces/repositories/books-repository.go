package repositories

import (
	"github.com/Mth-Ryan/waveaction-blog/internal/domain/entities"
	"github.com/google/uuid"
)

type BooksRepository interface {
	Get(id uuid.UUID) (entities.Book, error)
	GetAll() ([]entities.Book, error)
	Create(entity *entities.Book) error
	Update(id uuid.UUID, entity *entities.Book) error
	Delete(id uuid.UUID) error
}
