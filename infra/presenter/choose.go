package presenter

func Choose(format string) Formatter {
	switch format {
	case "xml":
		return &XMLPresenter{}
	case "json":
		return &JSONPresenter{}
	default:
		return &JSONPresenter{}
	}
}
