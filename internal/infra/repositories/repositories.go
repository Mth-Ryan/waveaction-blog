package repositories

import (
	"github.com/Mth-Ryan/waveaction-blog/internal/application/interfaces/repositories"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewBooksRepository,
		fx.As(new(repositories.BooksRepository)),
	),
)
