package translate

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/localization"
	messagePkg "github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

func Translate(locale localization.Localization, message string, args ...any) string {
	switch message {
	case messagePkg.Hello:
		return locale.Translate(messagePkg.Hello, args...)
	case messagePkg.SquareOutOfBounds:
		return locale.Translate(messagePkg.SquareOutOfBounds, args...)
	case messagePkg.NotAPiece:
		return locale.Translate(messagePkg.NotAPiece, args...)
	case messagePkg.InvalidDirection:
		return locale.Translate(messagePkg.InvalidDirection, args...)
	case messagePkg.InvalidSquareToMove:
		return locale.Translate(messagePkg.InvalidSquareToMove, args...)
	case messagePkg.TheresAPieceBlocking:
		return locale.Translate(messagePkg.TheresAPieceBlocking, args...)
	default:
		return locale.Translate(message)
	}
}
