package state

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

func ParsePiece(square entity.Square) (entity.Piece, error) {
	indices := GetBoardIndices(square)
	piece := Board[indices[0]][indices[1]]
	if piece == "" || len(piece) < 2 {
		return nil, errors.NewTrace(message.NotAPiece)
	}
	switch piece[0] {
	case 'P':
		return entity.Pawn{Indices: indices, Square: square}, nil
	case 'B':
		return entity.Bishop{Indices: indices, Square: square}, nil
	case 'T':
		return entity.Knight{Indices: indices, Square: square}, nil
	case 'R':
		return entity.Rook{Indices: indices, Square: square}, nil
	case 'Q':
		return entity.Queen{Indices: indices, Square: square}, nil
	case 'K':
		return entity.King{Indices: indices, Square: square}, nil
	default:
		return nil, errors.NewTrace(message.NotAPiece)
	}
}
