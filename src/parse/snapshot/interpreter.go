package snapshot

// Interpreter structure for context-aware and time-sensitive snapshot information
type Interpreter struct {
	Filename        string
	InitialisedCall *InitialisedCall
}

// InitialisedCall contains function call resolved arguments just before body execution
type InitialisedCall struct {
	Name string
	Args map[string]Argument
}

// Argument ...
type Argument struct {
	Str     *StringArgument
	StrList *StringListArgument
	Other   interface{}
}

// StringArgument ...
type StringArgument struct {
	Value string
}

// StringListArgument ...
type StringListArgument struct {
	Value []string
}
