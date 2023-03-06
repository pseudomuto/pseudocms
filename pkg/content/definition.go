package content

// Definition describes the details and metadata about content. It can be thought
// of like a schema for content.
type Definition struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Fields      []Field `json:"fields"`
}

// Item defines a single piece of content.
type Item struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Data        map[string]interface{} `json:"data"`
}

// List defines a list of content.
type List []Item
