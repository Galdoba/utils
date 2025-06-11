package text

func Wrap(text string, options ...WrapOption) string {
	wo := wrapOptions{}
	for _, modify := range options {
		modify(&wo)
	}

}

type wrapOptions struct {
}

type WrapOption func(*wrapOptions)
