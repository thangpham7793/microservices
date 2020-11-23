package handlers

//APIError covers all error types
type APIError string

//Error Types
const (
	ErrorRetrieveProducts APIError = "could not retrieve products"
	ErrorEncodeJSON       APIError = "could not encode json"
	ErrorDecodeJSON       APIError = "could not decode product"
	ErrorSaveProduct      APIError = "could not save product"
	ErrorPageNotFound     APIError = "page not found"
)
