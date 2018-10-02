package wiki

import (
	"encoding/xml"
	"io"
)

type XMLSiteInfo struct {
	Sitename   string `xml:"sitename"`
	Base       string `xml:"base"`
	Generator  string `xml:"generator"`
	Case       string `xml:"case"`
	Namespaces []struct {
		Key   string `xml:"key,attr"`
		Case  string `xml:"case,attr"`
		Value string `xml:",chardata"`
	} `xml:"namespaces>namespace"`
}

type XMLPage struct {
	ID        uint64        `xml:"id"`
	Title     string        `xml:"title"`
	Namespace uint          `xml:"ns"`
	Revisions []XMLRevision `xml:"revision"`
}

type XMLRevision struct {
	ID        uint64 `xml:"id"`
	Timestamp string `xml:"timestamp"`
	Text      string `xml:"text"`
}

type xmlParser struct {
	x *xml.Decoder
}

func NewXMLParser(r io.Reader) (*xmlParser, error) {
	d := xml.NewDecoder(r)
	_, err := d.Token()
	if err != nil {
		return nil, err
	}
	s := new(XMLSiteInfo)
	if err := d.Decode(&s); err != nil {
		return nil, err

	}
	return &xmlParser{x: d}, nil
}

func (p *xmlParser) Next() (*XMLPage, error) {
	rv := &XMLPage{}
	return rv, p.x.Decode(rv)
}
