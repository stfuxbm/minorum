package models

// Standard API response struct
type Response struct {
	Success bool        `json:"success"` // Status of the request
	Data    interface{} `json:"data"`    // Dynamic data (can be struct, slice, string, etc.)
	Message string      `json:"message"` // Message (success/failure)
	Code    int         `json:"code"`    // HTTP status code (200, 400, etc.)
}

const (
	// Common error messages
	MsgMethodNotAllowed    = "Method not allowed"
	MsgInvalidJSON         = "Invalid JSON format"
	MsgInternalServerError = "An internal server error occurred"
	MsgInvalidIDFormat     = "Invalid ID format"
	MsgFieldRequired       = "Required fields must be filled"

	// Quote messages
	MsgQuoteAdded    = "Quote successfully added"
	MsgQuoteExists   = "Quote already exists"
	MsgQuoteNotFound = "Quote not found"
	MsgQuoteList     = "List of quotes retrieved successfully"
	MsgQuoteDeleted  = "Quote successfully deleted"
	MsgQuoteUpdated  = "Quote successfully updated"

	// Author messages
	MsgAuthorAdded    = "Author successfully added"
	MsgAuthorExists   = "Author already exists"
	MsgAuthorNotFound = "Author not found"
	MsgAuthorList     = "List of authors retrieved successfully"
	MsgAuthorDeleted  = "Author successfully deleted"
	MsgAuthorUpdated  = "Author successfully updated"
)
