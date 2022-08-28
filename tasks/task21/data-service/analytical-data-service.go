package dataservice

import "log"

type AnalyticalDataService interface {
	SendXmlData()
}

type XmlDocuent struct {
}

func (doc XmlDocuent) SendXmlData() {
	log.Println("[Info] XML document was sucsesfully sent")
}
