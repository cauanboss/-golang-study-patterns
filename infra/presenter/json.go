package presenter

import (
	"encoding/json"
)

type JSONPresenter struct{}

func (j *JSONPresenter) Render(data interface{}) ([]byte, string, error) {
	content, err := json.Marshal(data)
	if err != nil {
		return nil, "", err
	}
	return content, "application/json", nil
}
