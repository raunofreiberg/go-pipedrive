package pipedrive

import "net/http"

type NoteFieldsService service

type Option struct {
	Id    int    `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
}

type NoteField struct {
	Id                   int      `json:"id,omitempty"`
	Key                  string   `json:"key,omitempty"`
	Name                 string   `json:"name,omitempty"`
	ActiveFlag           bool     `json:"active_flag,omitempty"`
	FieldType            string   `json:"field_type,omitempty"`
	EditFlag             int      `json:"edit_flag,omitempty"`
	MandatoryFlag        bool     `json:"mandatory_flag,omitempty"`
	VisibleInExportsFlag bool     `json:"visible_in_exports_flag,omitempty"`
	Options              []Option `json:"options,omitempty"`
}

type NoteFields struct {
	Success        bool           `json:"success,omitempty"`
	Data           []NoteField    `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// Return list of all fields for note.
func (s *NoteFieldsService) List() (*NoteFields, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/noteFields", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *NoteFields

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
