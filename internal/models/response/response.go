package models

// Standard API response struct
type Response struct {
	Success bool        `json:"success"` // Request status
	Data    interface{} `json:"data"`    // Dynamic payload (can be struct, slice, string, etc.)
	Message string      `json:"message"` // Message (success or error)
	Code    int         `json:"code"`    // HTTP status code (e.g., 200, 400)
}

const (
	// Common error messages
	MsgMethodNotAllowed    = "Method not allowed"
	MsgInvalidJSON         = "Invalid JSON format"
	MsgInternalServerError = "An internal server error occurred"
	MsgInvalidIDFormat     = "Invalid ID format"
	MsgFieldRequired       = "All required fields must be filled"

	// Quote messages
	MsgQuoteAdded    = "Quote successfully added"
	MsgQuoteExists   = "Quote already exists"
	MsgQuoteNotFound = "Quote not found"
	MsgQuoteList     = "Quotes retrieved successfully"
	MsgQuoteDeleted  = "Quote successfully deleted"
	MsgQuoteUpdated  = "Quote successfully updated"

	// Author messages
	MsgAuthorAdded    = "Author successfully added"
	MsgAuthorExists   = "Author already exists"
	MsgAuthorNotFound = "Author not found"
	MsgAuthorList     = "Authors retrieved successfully"
	MsgAuthorDeleted  = "Author successfully deleted"
	MsgAuthorUpdated  = "Author successfully updated"

	// Rate limiting message
	MsgTooManyRequests = "Too many requests, please try again later"
)
