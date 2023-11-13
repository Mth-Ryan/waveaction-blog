package eventhandlers

import (
	events "github.com/Mth-Ryan/waveaction-blog/internal/domain/events/books"
)

type BooksEventHandler interface {
	Handle(event events.BookEvent) 
}
