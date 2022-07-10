package proto

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
)

func (m *Hello) MarshalJSON() ([]byte, error) {
	b, err := protojson.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

var _ json.Marshaler = (*Hello)(nil)

func (m *Hello) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, m)
}

var _ json.Unmarshaler = (*Hello)(nil)
