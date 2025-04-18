package state

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

func parsePiece(square entity.Square) (entity.Piece, error) {
	indices := getBoardIndices(square)
	piece := Board[indices[0]][indices[1]]
	if piece == "" || len(piece) < 2 {
		return nil, errors.NewTrace(message.NotAPiece)
	}
	switch piece[0] {
	case 'P':
		return entity.Pawn(indices), nil
	case 'B':
		return entity.Bishop{Indices: indices, Square: square}, nil
	case 'T':
		return entity.Knight(indices), nil
	case 'R':
		return entity.Rook(indices), nil
	case 'Q':
		return entity.Queen(indices), nil
	case 'K':
		return entity.King(indices), nil
	default:
		return nil, errors.NewTrace(message.NotAPiece)
	}
}
