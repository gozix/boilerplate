// +build tools
// See https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md :(

package tools

import (
	"encoding/json"

	_ "github.com/mailru/easyjson"
)

func handleMessage() (interface{}, error) {
	var msg interface{}
	if len(b) == 0 {
		return msg, nil
	}

	err = json.Unmarshal(body, &msg)
	return msg, err
}
