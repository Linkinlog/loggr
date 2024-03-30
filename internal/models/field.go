package models

func NewField(t, d string) *Field {
	return &Field{
		Title:       t,
		Description: d,
	}
}

type Field struct {
	Title       string
	Description string
}
