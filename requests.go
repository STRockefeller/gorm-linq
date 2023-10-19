package linq

type LockOption int

const (
	OptionWait LockOption = iota
	OptionNoWait
)

type LockableRequest interface {
	// returns true if you want to lock the query data
	Lock() bool
	LockOption() LockOption
}

// Contains a function called Where which returns an array of QueryString objects. Each QueryString object represents a condition to be applied to the container's db field.
type QueryRequest interface {
	Where() []QueryString
}

type UpdateRequest interface {
	Update() map[string]any
}
