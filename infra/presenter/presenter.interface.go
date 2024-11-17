package presenter

// Formatter é a interface que todos os formatos implementam.
type Formatter interface {
	Render(data interface{}) ([]byte, string, error)
}
