package handler

type Handler struct {
	HandlerManager HandlerManager
}

func New() Handler {
	return Handler{}
}
