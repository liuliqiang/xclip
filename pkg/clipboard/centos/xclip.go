package centos

// https://github.com/atotto/clipboard
type xclipClipboard struct {
}

func (x *xclipClipboard) SetContent([]byte) error {
	panic("implement me")
}

func (x *xclipClipboard) GetContent() ([]byte, error) {
	panic("implement me")
}
