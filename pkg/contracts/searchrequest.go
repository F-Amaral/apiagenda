package contracts

type SearchRequest struct {
	Id      *string
	Name    *string
	Email   *string
	Deleted *bool
}
