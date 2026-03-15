package main

import (
	"encoding/xml"
	"io"
)

func filter(r io.Reader, w io.Writer) error {
	decoder := xml.NewDecoder(r)
	encoder := xml.NewEncoder(w)

	var stack []bool

	for {
		tok, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			attrs, allowed := allowlist[t.Name.Local]
			stack = append(stack, allowed)

			if !allowed {
				continue
			}

			var filtered []xml.Attr
			for _, attr := range t.Attr {
				for _, a := range attrs {
					if attr.Name.Local == a {
						filtered = append(filtered, attr)
						break
					}
				}
			}

			t.Attr = filtered
			if err := encoder.EncodeToken(t); err != nil {
				return err
			}

		case xml.EndElement:
			if len(stack) == 0 {
				continue
			}

			allowed := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if !allowed {
				continue
			}

			if err := encoder.EncodeToken(t); err != nil {
				return err
			}

		case xml.CharData:
			if len(stack) > 0 && stack[len(stack)-1] {
				if err := encoder.EncodeToken(t); err != nil {
					return err
				}
			}

		case xml.Comment:
		default:
		}
	}

	return encoder.Flush()
}
