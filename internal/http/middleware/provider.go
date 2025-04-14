package middleware

func ProvideMiddlewares(
	log *Log,
) []MiddleWare {
	return []MiddleWare{log}
}
