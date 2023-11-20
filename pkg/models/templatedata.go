package models

type TemplateData struct {
	Data    map[string]interface{}
	CSFR    string
	Flash   string
	Warning string
	Error   string
}
