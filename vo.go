package easydb

// Where where
type Where struct {
	Opt   LogicalOptType
	Key   string
	Value string
	Ins   []string
}

// On on
type On struct {
	Opt   LogicalOptType
	Key   string
	Value string
	Ins   []string
}

// Having having
type Having struct {
	Opt   LogicalOptType
	Key   string
	Value string
	Ins   []string
}

// Order order
type Order struct {
	Type OrderType
	Key  string
}

// Table table
type Table struct {
	Name string
	As   string
}

// Column column
type Column struct {
	Name  string
	As    string
	Value string
}

// QueryFunc query func
type QueryFunc struct {
	Type  QueryType
	Names []string
	As    string
}
