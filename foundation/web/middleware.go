package web

type Middleware func(handler Handler) Handler

func wrapMiddlewares(handler Handler, mids ...Middleware) Handler {
	for _, mid := range mids {
		if mid == nil {
			continue
		}
		handler = mid(handler)
	}

	return handler
}
