package data

type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXML() string {
	return "<xml></xml>"
}

// inject json struct
type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {

	adapter.jsonDocument.ConvertToXML()

}
