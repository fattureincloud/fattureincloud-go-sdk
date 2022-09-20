package filter

type Expression interface {
	BuildQuery() string
	String() string
}
