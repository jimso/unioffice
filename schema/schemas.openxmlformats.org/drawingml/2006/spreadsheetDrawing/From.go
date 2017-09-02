// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"
	"log"
)

type From struct {
	CT_Marker
}

func NewFrom() *From {
	ret := &From{}
	ret.CT_Marker = *NewCT_Marker()
	return ret
}

func (m *From) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "xdr:from"
	return m.CT_Marker.MarshalXML(e, start)
}

func (m *From) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Marker = *NewCT_Marker()
lFrom:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "col":
				if err := d.DecodeElement(m.Col, &el); err != nil {
					return err
				}
			case "colOff":
				if err := d.DecodeElement(m.ColOff, &el); err != nil {
					return err
				}
			case "row":
				if err := d.DecodeElement(m.Row, &el); err != nil {
					return err
				}
			case "rowOff":
				if err := d.DecodeElement(m.RowOff, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on From %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lFrom
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the From and its children
func (m *From) Validate() error {
	return m.ValidateWithPath("From")
}

// ValidateWithPath validates the From and its children, prefixing error messages with path
func (m *From) ValidateWithPath(path string) error {
	if err := m.CT_Marker.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
