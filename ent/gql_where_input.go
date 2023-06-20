// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"shrektionary_api/ent/definition"
	"shrektionary_api/ent/group"
	"shrektionary_api/ent/predicate"
	"shrektionary_api/ent/user"
	"shrektionary_api/ent/word"
)

// DefinitionWhereInput represents a where input for filtering Definition queries.
type DefinitionWhereInput struct {
	Predicates []predicate.Definition  `json:"-"`
	Not        *DefinitionWhereInput   `json:"not,omitempty"`
	Or         []*DefinitionWhereInput `json:"or,omitempty"`
	And        []*DefinitionWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "description" field predicates.
	Description             *string  `json:"description,omitempty"`
	DescriptionNEQ          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGT           *string  `json:"descriptionGT,omitempty"`
	DescriptionGTE          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLT           *string  `json:"descriptionLT,omitempty"`
	DescriptionLTE          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`

	// "word" edge predicates.
	HasWord     *bool             `json:"hasWord,omitempty"`
	HasWordWith []*WordWhereInput `json:"hasWordWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *DefinitionWhereInput) AddPredicates(predicates ...predicate.Definition) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the DefinitionWhereInput filter on the DefinitionQuery builder.
func (i *DefinitionWhereInput) Filter(q *DefinitionQuery) (*DefinitionQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyDefinitionWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyDefinitionWhereInput is returned in case the DefinitionWhereInput is empty.
var ErrEmptyDefinitionWhereInput = errors.New("ent: empty predicate DefinitionWhereInput")

// P returns a predicate for filtering definitions.
// An error is returned if the input is empty or invalid.
func (i *DefinitionWhereInput) P() (predicate.Definition, error) {
	var predicates []predicate.Definition
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, definition.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Definition, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, definition.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Definition, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, definition.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, definition.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, definition.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, definition.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, definition.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, definition.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, definition.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, definition.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, definition.IDLTE(*i.IDLTE))
	}
	if i.Description != nil {
		predicates = append(predicates, definition.DescriptionEQ(*i.Description))
	}
	if i.DescriptionNEQ != nil {
		predicates = append(predicates, definition.DescriptionNEQ(*i.DescriptionNEQ))
	}
	if len(i.DescriptionIn) > 0 {
		predicates = append(predicates, definition.DescriptionIn(i.DescriptionIn...))
	}
	if len(i.DescriptionNotIn) > 0 {
		predicates = append(predicates, definition.DescriptionNotIn(i.DescriptionNotIn...))
	}
	if i.DescriptionGT != nil {
		predicates = append(predicates, definition.DescriptionGT(*i.DescriptionGT))
	}
	if i.DescriptionGTE != nil {
		predicates = append(predicates, definition.DescriptionGTE(*i.DescriptionGTE))
	}
	if i.DescriptionLT != nil {
		predicates = append(predicates, definition.DescriptionLT(*i.DescriptionLT))
	}
	if i.DescriptionLTE != nil {
		predicates = append(predicates, definition.DescriptionLTE(*i.DescriptionLTE))
	}
	if i.DescriptionContains != nil {
		predicates = append(predicates, definition.DescriptionContains(*i.DescriptionContains))
	}
	if i.DescriptionHasPrefix != nil {
		predicates = append(predicates, definition.DescriptionHasPrefix(*i.DescriptionHasPrefix))
	}
	if i.DescriptionHasSuffix != nil {
		predicates = append(predicates, definition.DescriptionHasSuffix(*i.DescriptionHasSuffix))
	}
	if i.DescriptionEqualFold != nil {
		predicates = append(predicates, definition.DescriptionEqualFold(*i.DescriptionEqualFold))
	}
	if i.DescriptionContainsFold != nil {
		predicates = append(predicates, definition.DescriptionContainsFold(*i.DescriptionContainsFold))
	}

	if i.HasWord != nil {
		p := definition.HasWord()
		if !*i.HasWord {
			p = definition.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasWordWith) > 0 {
		with := make([]predicate.Word, 0, len(i.HasWordWith))
		for _, w := range i.HasWordWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasWordWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, definition.HasWordWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyDefinitionWhereInput
	case 1:
		return predicates[0], nil
	default:
		return definition.And(predicates...), nil
	}
}

// GroupWhereInput represents a where input for filtering Group queries.
type GroupWhereInput struct {
	Predicates []predicate.Group  `json:"-"`
	Not        *GroupWhereInput   `json:"not,omitempty"`
	Or         []*GroupWhereInput `json:"or,omitempty"`
	And        []*GroupWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "description" field predicates.
	Description             *string  `json:"description,omitempty"`
	DescriptionNEQ          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGT           *string  `json:"descriptionGT,omitempty"`
	DescriptionGTE          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLT           *string  `json:"descriptionLT,omitempty"`
	DescriptionLTE          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`

	// "words" edge predicates.
	HasWords     *bool             `json:"hasWords,omitempty"`
	HasWordsWith []*WordWhereInput `json:"hasWordsWith,omitempty"`

	// "users" edge predicates.
	HasUsers     *bool             `json:"hasUsers,omitempty"`
	HasUsersWith []*UserWhereInput `json:"hasUsersWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *GroupWhereInput) AddPredicates(predicates ...predicate.Group) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the GroupWhereInput filter on the GroupQuery builder.
func (i *GroupWhereInput) Filter(q *GroupQuery) (*GroupQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyGroupWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyGroupWhereInput is returned in case the GroupWhereInput is empty.
var ErrEmptyGroupWhereInput = errors.New("ent: empty predicate GroupWhereInput")

// P returns a predicate for filtering groups.
// An error is returned if the input is empty or invalid.
func (i *GroupWhereInput) P() (predicate.Group, error) {
	var predicates []predicate.Group
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, group.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Group, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, group.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Group, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, group.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, group.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, group.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, group.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, group.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, group.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, group.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, group.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, group.IDLTE(*i.IDLTE))
	}
	if i.Description != nil {
		predicates = append(predicates, group.DescriptionEQ(*i.Description))
	}
	if i.DescriptionNEQ != nil {
		predicates = append(predicates, group.DescriptionNEQ(*i.DescriptionNEQ))
	}
	if len(i.DescriptionIn) > 0 {
		predicates = append(predicates, group.DescriptionIn(i.DescriptionIn...))
	}
	if len(i.DescriptionNotIn) > 0 {
		predicates = append(predicates, group.DescriptionNotIn(i.DescriptionNotIn...))
	}
	if i.DescriptionGT != nil {
		predicates = append(predicates, group.DescriptionGT(*i.DescriptionGT))
	}
	if i.DescriptionGTE != nil {
		predicates = append(predicates, group.DescriptionGTE(*i.DescriptionGTE))
	}
	if i.DescriptionLT != nil {
		predicates = append(predicates, group.DescriptionLT(*i.DescriptionLT))
	}
	if i.DescriptionLTE != nil {
		predicates = append(predicates, group.DescriptionLTE(*i.DescriptionLTE))
	}
	if i.DescriptionContains != nil {
		predicates = append(predicates, group.DescriptionContains(*i.DescriptionContains))
	}
	if i.DescriptionHasPrefix != nil {
		predicates = append(predicates, group.DescriptionHasPrefix(*i.DescriptionHasPrefix))
	}
	if i.DescriptionHasSuffix != nil {
		predicates = append(predicates, group.DescriptionHasSuffix(*i.DescriptionHasSuffix))
	}
	if i.DescriptionEqualFold != nil {
		predicates = append(predicates, group.DescriptionEqualFold(*i.DescriptionEqualFold))
	}
	if i.DescriptionContainsFold != nil {
		predicates = append(predicates, group.DescriptionContainsFold(*i.DescriptionContainsFold))
	}

	if i.HasWords != nil {
		p := group.HasWords()
		if !*i.HasWords {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasWordsWith) > 0 {
		with := make([]predicate.Word, 0, len(i.HasWordsWith))
		for _, w := range i.HasWordsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasWordsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasWordsWith(with...))
	}
	if i.HasUsers != nil {
		p := group.HasUsers()
		if !*i.HasUsers {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUsersWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasUsersWith))
		for _, w := range i.HasUsersWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUsersWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasUsersWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyGroupWhereInput
	case 1:
		return predicates[0], nil
	default:
		return group.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "authID" field predicates.
	AuthID             *string  `json:"authid,omitempty"`
	AuthIDNEQ          *string  `json:"authidNEQ,omitempty"`
	AuthIDIn           []string `json:"authidIn,omitempty"`
	AuthIDNotIn        []string `json:"authidNotIn,omitempty"`
	AuthIDGT           *string  `json:"authidGT,omitempty"`
	AuthIDGTE          *string  `json:"authidGTE,omitempty"`
	AuthIDLT           *string  `json:"authidLT,omitempty"`
	AuthIDLTE          *string  `json:"authidLTE,omitempty"`
	AuthIDContains     *string  `json:"authidContains,omitempty"`
	AuthIDHasPrefix    *string  `json:"authidHasPrefix,omitempty"`
	AuthIDHasSuffix    *string  `json:"authidHasSuffix,omitempty"`
	AuthIDEqualFold    *string  `json:"authidEqualFold,omitempty"`
	AuthIDContainsFold *string  `json:"authidContainsFold,omitempty"`

	// "groups" edge predicates.
	HasGroups     *bool              `json:"hasGroups,omitempty"`
	HasGroupsWith []*GroupWhereInput `json:"hasGroupsWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("ent: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.AuthID != nil {
		predicates = append(predicates, user.AuthIDEQ(*i.AuthID))
	}
	if i.AuthIDNEQ != nil {
		predicates = append(predicates, user.AuthIDNEQ(*i.AuthIDNEQ))
	}
	if len(i.AuthIDIn) > 0 {
		predicates = append(predicates, user.AuthIDIn(i.AuthIDIn...))
	}
	if len(i.AuthIDNotIn) > 0 {
		predicates = append(predicates, user.AuthIDNotIn(i.AuthIDNotIn...))
	}
	if i.AuthIDGT != nil {
		predicates = append(predicates, user.AuthIDGT(*i.AuthIDGT))
	}
	if i.AuthIDGTE != nil {
		predicates = append(predicates, user.AuthIDGTE(*i.AuthIDGTE))
	}
	if i.AuthIDLT != nil {
		predicates = append(predicates, user.AuthIDLT(*i.AuthIDLT))
	}
	if i.AuthIDLTE != nil {
		predicates = append(predicates, user.AuthIDLTE(*i.AuthIDLTE))
	}
	if i.AuthIDContains != nil {
		predicates = append(predicates, user.AuthIDContains(*i.AuthIDContains))
	}
	if i.AuthIDHasPrefix != nil {
		predicates = append(predicates, user.AuthIDHasPrefix(*i.AuthIDHasPrefix))
	}
	if i.AuthIDHasSuffix != nil {
		predicates = append(predicates, user.AuthIDHasSuffix(*i.AuthIDHasSuffix))
	}
	if i.AuthIDEqualFold != nil {
		predicates = append(predicates, user.AuthIDEqualFold(*i.AuthIDEqualFold))
	}
	if i.AuthIDContainsFold != nil {
		predicates = append(predicates, user.AuthIDContainsFold(*i.AuthIDContainsFold))
	}

	if i.HasGroups != nil {
		p := user.HasGroups()
		if !*i.HasGroups {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasGroupsWith) > 0 {
		with := make([]predicate.Group, 0, len(i.HasGroupsWith))
		for _, w := range i.HasGroupsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasGroupsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasGroupsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}

// WordWhereInput represents a where input for filtering Word queries.
type WordWhereInput struct {
	Predicates []predicate.Word  `json:"-"`
	Not        *WordWhereInput   `json:"not,omitempty"`
	Or         []*WordWhereInput `json:"or,omitempty"`
	And        []*WordWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "description" field predicates.
	Description             *string  `json:"description,omitempty"`
	DescriptionNEQ          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGT           *string  `json:"descriptionGT,omitempty"`
	DescriptionGTE          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLT           *string  `json:"descriptionLT,omitempty"`
	DescriptionLTE          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`

	// "definitions" edge predicates.
	HasDefinitions     *bool                   `json:"hasDefinitions,omitempty"`
	HasDefinitionsWith []*DefinitionWhereInput `json:"hasDefinitionsWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *WordWhereInput) AddPredicates(predicates ...predicate.Word) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the WordWhereInput filter on the WordQuery builder.
func (i *WordWhereInput) Filter(q *WordQuery) (*WordQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyWordWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyWordWhereInput is returned in case the WordWhereInput is empty.
var ErrEmptyWordWhereInput = errors.New("ent: empty predicate WordWhereInput")

// P returns a predicate for filtering words.
// An error is returned if the input is empty or invalid.
func (i *WordWhereInput) P() (predicate.Word, error) {
	var predicates []predicate.Word
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, word.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Word, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, word.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Word, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, word.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, word.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, word.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, word.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, word.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, word.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, word.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, word.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, word.IDLTE(*i.IDLTE))
	}
	if i.Description != nil {
		predicates = append(predicates, word.DescriptionEQ(*i.Description))
	}
	if i.DescriptionNEQ != nil {
		predicates = append(predicates, word.DescriptionNEQ(*i.DescriptionNEQ))
	}
	if len(i.DescriptionIn) > 0 {
		predicates = append(predicates, word.DescriptionIn(i.DescriptionIn...))
	}
	if len(i.DescriptionNotIn) > 0 {
		predicates = append(predicates, word.DescriptionNotIn(i.DescriptionNotIn...))
	}
	if i.DescriptionGT != nil {
		predicates = append(predicates, word.DescriptionGT(*i.DescriptionGT))
	}
	if i.DescriptionGTE != nil {
		predicates = append(predicates, word.DescriptionGTE(*i.DescriptionGTE))
	}
	if i.DescriptionLT != nil {
		predicates = append(predicates, word.DescriptionLT(*i.DescriptionLT))
	}
	if i.DescriptionLTE != nil {
		predicates = append(predicates, word.DescriptionLTE(*i.DescriptionLTE))
	}
	if i.DescriptionContains != nil {
		predicates = append(predicates, word.DescriptionContains(*i.DescriptionContains))
	}
	if i.DescriptionHasPrefix != nil {
		predicates = append(predicates, word.DescriptionHasPrefix(*i.DescriptionHasPrefix))
	}
	if i.DescriptionHasSuffix != nil {
		predicates = append(predicates, word.DescriptionHasSuffix(*i.DescriptionHasSuffix))
	}
	if i.DescriptionEqualFold != nil {
		predicates = append(predicates, word.DescriptionEqualFold(*i.DescriptionEqualFold))
	}
	if i.DescriptionContainsFold != nil {
		predicates = append(predicates, word.DescriptionContainsFold(*i.DescriptionContainsFold))
	}

	if i.HasDefinitions != nil {
		p := word.HasDefinitions()
		if !*i.HasDefinitions {
			p = word.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasDefinitionsWith) > 0 {
		with := make([]predicate.Definition, 0, len(i.HasDefinitionsWith))
		for _, w := range i.HasDefinitionsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasDefinitionsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, word.HasDefinitionsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyWordWhereInput
	case 1:
		return predicates[0], nil
	default:
		return word.And(predicates...), nil
	}
}
