package translate

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/localization"
	messagePkg "github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

func Translate(locale localization.Localization, message string, args ...any) string {
	switch message {
	case messagePkg.Hello:
		return locale.Translate(messagePkg.Hello, args...)
	default:
		return locale.Translate(message)
	}
}
