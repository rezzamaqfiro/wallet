package constant

type ctxKey string

const (
	ContextBirthTime ctxKey = "birth-time"
	ContextMessageID ctxKey = "nsq-message-id"
)
