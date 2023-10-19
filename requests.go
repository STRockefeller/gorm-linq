package linq

type LockOption int

const (
	OptionWait LockOption = iota
	OptionNoWait
)

type LockableRequest interface {
	Lock() bool
	LockOption() LockOption
}

type QueryRequest interface {
	Where() []QueryString
}

type UpdateRequest interface {
	Update() map[string]any
}
