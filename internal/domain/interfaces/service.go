package interfaces

import (
	"context"

	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/entity"
	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/form"
)

// WordOfWisdomService is an interface for the WordOfWisdom service.
type WordOfWisdomService interface {
	GetChallenge(ctx context.Context) ([]byte, uint8, error)
	SubmitSolution(ctx context.Context, solution form.SubmitSolution) (entity.Quote, error)
}
