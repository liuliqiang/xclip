package clipboard

type Clipboard interface {
	SetContent([]byte) error
	GetContent() ([]byte, error)
}
