package fields

import (
	"fmt"
	"strings"

	"github.com/chhz0/goiam/pkg/meta/selection"
)

// Selector 是一个字段选择器，用以在一组数据中根据给定的字段和条件进行筛选和匹配
type Selector interface {

	// Matches 判断给定的 Fields 是否与 Selector 匹配
	Matches(Fields) bool

	// Empty 判断 Selector 是否为空
	Empty() bool

	// RequiresExactMatch 判断 Selector 是否需要精确匹配给定的字段，并返回该字段的值和是否找到该字段
	RequiresExactMatch(field string) (value string, found bool)

	// Transform 将 Selector 中的字段值进行转换，并返回新的 Selector 和错误信息
	Transform(TransformFunc) (Selector, error)

	// Requirement 返回 Selector 中的字段和值，并返回一个 Requirement 结构体
	// 返回值 Requirement 结构体提供更详细的信息，包含字段名、操作符和值
	Requirements() Requirements

	// String 返回 Selector 的字符串表示形式
	String() string

	// DeepCopy 返回 Selector 的深度拷贝
	DeepCopy() Selector
}

type TransformFunc func(field, value string) (newField, newValue string, err error)

type nothingSelector struct{}

func (n nothingSelector) Matches(_ Fields) bool                       { return false }
func (n nothingSelector) Empty() bool                                 { return false }
func (n nothingSelector) RequiresExactMatch(_ string) (string, bool)  { return "", false }
func (n nothingSelector) Transform(_ TransformFunc) (Selector, error) { return n, nil }
func (n nothingSelector) Requirements() Requirements                  { return nil }
func (n nothingSelector) String() string                              { return "" }
func (n nothingSelector) DeepCopy() Selector                          { return n }

// Nothing 返回一个不拥有任何Fields的Selector
func Nothing() Selector {
	return nothingSelector{}
}

type hasTerm struct {
	field, value string
}

func (h *hasTerm) Matches(ls Fields) bool {
	return ls.Get(h.field) == h.value
}

func (h *hasTerm) Empty() bool {
	return false
}

func (h *hasTerm) RequiresExactMatch(field string) (string, bool) {
	if h.field == field {
		return h.value, true
	}

	return "", false
}

func (h *hasTerm) Transform(f TransformFunc) (Selector, error) {
	field, value, err := f(h.field, h.value)
	if err != nil {
		return nil, err
	}

	if len(field) == 0 && len(value) == 0 {
		return nil, nil
	}

	return &hasTerm{field, value}, nil
}

func (h *hasTerm) Requirements() Requirements {
	return []Requirement{{
		Field:    h.field,
		Operator: selection.Equals,
		Value:    h.value,
	}}
}

func (h *hasTerm) String() string {
	return fmt.Sprintf("%v=%v", h.field, EscapeValue(h.value))
}

func (h *hasTerm) DeepCopy() Selector {
	if h == nil {
		return nil
	}
	out := new(hasTerm)
	*out = *h

	return out
}

type notHasTerm struct {
	field, value string
}

func (h *notHasTerm) Matches(ls Fields) bool {
	return ls.Get(h.field) != h.value
}

func (h *notHasTerm) Empty() bool {
	return false
}

func (h *notHasTerm) RequiresExactMatch(field string) (string, bool) {
	return "", false
}

func (h *notHasTerm) Transform(f TransformFunc) (Selector, error) {
	field, value, err := f(h.field, h.value)
	if err != nil {
		return nil, err
	}

	if len(field) == 0 && len(value) == 0 {
		return nil, nil
	}

	return &notHasTerm{field, value}, nil
}

func (h *notHasTerm) Requirements() Requirements {
	return []Requirement{{
		Field:    h.field,
		Operator: selection.NotEquals,
		Value:    h.value,
	}}
}

func (h *notHasTerm) String() string {
	return fmt.Sprintf("%v!=%v", h.field, EscapeValue(h.value))
}

func (h *notHasTerm) DeepCopy() Selector {
	if h == nil {
		return nil
	}
	out := new(notHasTerm)
	*out = *h

	return out
}

func Everything() Selector {
	return andTerm{}
}

type andTerm []Selector

func (at andTerm) Matches(ls Fields) bool {
	for _, s := range at {
		if !s.Matches(ls) {
			return false
		}
	}
	return true
}

// Empty implements Selector.
func (at andTerm) Empty() bool {
	if at == nil {
		return false
	}

	// TODO: len([]Selector(at)) == 0
	if len(at) == 0 {
		return true
	}
	for _, s := range at {
		if !s.Empty() {
			return false
		}
	}

	return true
}

// RequiresExactMatch implements Selector.
func (at andTerm) RequiresExactMatch(field string) (value string, found bool) {
	if at.Empty() || len(at) == 0 {
		return "", false
	}
	for _, s := range at {
		if value, found = s.RequiresExactMatch(field); found {
			return value, found
		}
	}

	return "", false
}

// Transform implements Selector.
func (at andTerm) Transform(TransformFunc) (Selector, error) {
	new := make([]Selector, 0, len(at))
	for _, s := range at {
		newS, err := s.Transform(nil)
		if err != nil {
			return nil, err
		}
		if !newS.Empty() {
			new = append(new, newS)
		}
	}

	return andTerm(new), nil
}

// Requirements implements Selector.
func (at andTerm) Requirements() Requirements {
	reqs := make([]Requirement, 0, len(at))
	for _, s := range at {
		rs := s.Requirements()
		reqs = append(reqs, rs...)
	}

	return reqs
}

// String implements Selector.
func (at andTerm) String() string {
	terms := make([]string, 0, len(at))
	for _, s := range at {
		terms = append(terms, s.String())
	}

	return strings.Join(terms, ",")
}

// DeepCopy implements Selector.
func (at andTerm) DeepCopy() Selector {
	if at.Empty() {
		return nil
	}

	out := make([]Selector, 0, len(at))
	for _, s := range at {
		out = append(out, s.DeepCopy())
	}

	return andTerm(out)
}

func SelectorFromSet(ls Set) Selector {
	if ls == nil {
		return Everything()
	}

	terms := make([]Selector, 0, len(ls))
	for field, value := range ls {
		terms = append(terms, &hasTerm{field, value})
	}

	if len(terms) == 1 {
		return terms[0]
	}

	return andTerm(terms)
}

func EscapeValue(rStr string) string {
	return strings.NewReplacer(`\`, `\\`, `,`, `\,`, `=`, `\=`).Replace(rStr)
}
