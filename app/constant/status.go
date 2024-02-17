package constant

// ResponseStatus is a way to facilitate debugging and internal error handling
type ResponseStatus int

const (
	Success ResponseStatus = iota
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
	ParsingFailed
	DBQueryFailed
	DBNoRowsAffected
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{
		"SUCCESS",
		"DATA_NOT_FOUND",
		"UNKNOWN_ERROR",
		"INVALID_REQUEST",
		"UNAUTHORIZED",
		"PARSING_FAILED",
		"DB_QUERY_FAILED",
		"DB_NO_ROWS_AFFECTED"}[r]
}

func (r ResponseStatus) GetNumber() int {
	return int(r)
}
