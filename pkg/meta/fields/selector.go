package fields

type Selector interface {
	Matches(Fields) bool

	Empty() bool

	RequiresExactMatch(field string) (value string, found bool)

	Transform(TransformFunc) (Selector, error)

	Requirement() Requirement

	String() string
	DeepCopy() Selector
}

type TransformFunc func(field, value string) (newField, newValue string, err error)

