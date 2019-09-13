package mailyak

import "bytes"

// BodyPart is a buffer holding the contents of an email MIME part.
type BodyPart struct{
	bytes.Buffer
	MimeType string
}

// Set accepts a string s as the contents of a BodyPart, replacing any existing
// data.
func (w *BodyPart) Set(s string) {
	w.Reset()
	w.WriteString(s)
}

// SetMimeType sets the mime type
func (w *BodyPart) SetMimeType(m string) {
	w.MimeType = m
}
