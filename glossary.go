package glossary

import (
	"encoding/json"
	"fmt"

	"github.com/GoRethink/gorethink/encoding"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Glossary are essentially a JSON object as []byte
type Glossary []byte

// Get allows you to get the value of a path
func (ann *Glossary) Get(path string) (*gjson.Result, error) {
	if ann == nil {
		return nil, fmt.Errorf("no value found at path %s", path)
	}

	val := gjson.GetBytes(*ann, path)
	if !val.Exists() {
		return nil, fmt.Errorf("no value found at path %s", path)
	}

	return &val, nil
}

// Set places the value at the provided path
func (ann *Glossary) Set(path string, val interface{}) error {
	if ann == nil {
		*ann = []byte{}
	}

	res, err := sjson.SetBytes(*ann, path, val)
	if err != nil {
		return err
	}

	*ann = res

	return nil
}

// MarshalJSON implements json.Marshaler
func (ann *Glossary) MarshalJSON() ([]byte, error) {
	var m map[string]interface{}

	if ann == nil {
		return json.Marshal(m)
	}

	if len(*ann) == 0 {
		*ann = []byte(`{}`)
	}

	if err := json.Unmarshal(*ann, &m); err != nil {
		return nil, err
	}

	return json.Marshal(m)
}

// UnmarshalJSON implements json.Unmarshaler
func (ann *Glossary) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}

	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	*ann = b

	return nil
}

var _ encoding.Marshaler = (*Glossary)(nil)
var _ encoding.Unmarshaler = (*Glossary)(nil)

// MarshalRQL implements gorethink.RQLMarshaler
func (ann Glossary) MarshalRQL() (interface{}, error) {
	var m map[string]interface{}

	if ann == nil {
		return m, nil
	}

	if err := json.Unmarshal(ann, &m); err != nil {
		return nil, err
	}

	return m, nil
}

// UnmarshalRQL implements gorethink.RQLUnmarshaler
func (ann *Glossary) UnmarshalRQL(val interface{}) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}

	*ann = b

	return nil
}
