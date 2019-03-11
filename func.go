package mid

// func
type HandlerFunc func(*Context)

// chains
type HandlersChain []HandlerFunc

func (h HandlersChain) Last() HandlerFunc {
	if length := len(h); length > 0 {
		return h[length - 1]
	}
	return nil
}