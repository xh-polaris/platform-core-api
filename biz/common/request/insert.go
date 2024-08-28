package request

type InsertRequest struct {
	Body []*InsertItem
}

type InsertItem struct {
	EventName string
	Tags      map[string]interface{}
}
