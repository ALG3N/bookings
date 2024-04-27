package models

// TemplateData represents the data structure used for rendering templates.
type TemplateData struct {
	StringMap map[string]string      // StringMap stores key-value pairs of string data.
	IntMap    map[string]int         // IntMap stores key-value pairs of integer data.
	FloatMap  map[string]float32     // FloatMap stores key-value pairs of float data.
	Data      map[string]interface{} // Data stores key-value pairs of generic data.
	CSRFToken string                 // CSRFToken stores the Cross-Site Request Forgery token.
	Flash     string                 // Flash stores flash messages.
	Warning   string                 // Warning stores warning messages.
	Error     string                 // Error stores error messages.
}
