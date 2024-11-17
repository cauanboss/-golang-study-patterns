package presenter

import (
	"encoding/xml"
)

type XMLPresenter struct{}

func (x *XMLPresenter) Render(data interface{}) ([]byte, string, error) {
	content, err := xml.Marshal(data)
	if err != nil {
		return nil, "", err
	}
	return content, "application/xml", nil
}
