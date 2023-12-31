// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (d *Definition) Creator(ctx context.Context) (*User, error) {
	result, err := d.Edges.CreatorOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryCreator().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (d *Definition) Word(ctx context.Context) (*Word, error) {
	result, err := d.Edges.WordOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryWord().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (gr *Group) Words(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WordOrder, where *WordWhereInput,
) (*WordConnection, error) {
	opts := []WordPaginateOption{
		WithWordOrder(orderBy),
		WithWordFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := gr.Edges.totalCount[0][alias]
	if nodes, err := gr.NamedWords(alias); err == nil || hasTotalCount {
		pager, err := newWordPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WordConnection{Edges: []*WordEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return gr.QueryWords().Paginate(ctx, after, first, before, last, opts...)
}

func (gr *Group) Users(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = gr.NamedUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = gr.Edges.UsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = gr.QueryUsers().All(ctx)
	}
	return result, err
}

func (u *User) Groups(ctx context.Context) (result []*Group, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedGroups(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.GroupsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryGroups().All(ctx)
	}
	return result, err
}

func (u *User) Definitions(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *DefinitionOrder, where *DefinitionWhereInput,
) (*DefinitionConnection, error) {
	opts := []DefinitionPaginateOption{
		WithDefinitionOrder(orderBy),
		WithDefinitionFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[1][alias]
	if nodes, err := u.NamedDefinitions(alias); err == nil || hasTotalCount {
		pager, err := newDefinitionPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &DefinitionConnection{Edges: []*DefinitionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryDefinitions().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Words(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WordOrder, where *WordWhereInput,
) (*WordConnection, error) {
	opts := []WordPaginateOption{
		WithWordOrder(orderBy),
		WithWordFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[2][alias]
	if nodes, err := u.NamedWords(alias); err == nil || hasTotalCount {
		pager, err := newWordPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WordConnection{Edges: []*WordEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryWords().Paginate(ctx, after, first, before, last, opts...)
}

func (w *Word) Creator(ctx context.Context) (*User, error) {
	result, err := w.Edges.CreatorOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryCreator().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Word) Group(ctx context.Context) (*Group, error) {
	result, err := w.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryGroup().Only(ctx)
	}
	return result, err
}

func (w *Word) Definitions(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *DefinitionOrder, where *DefinitionWhereInput,
) (*DefinitionConnection, error) {
	opts := []DefinitionPaginateOption{
		WithDefinitionOrder(orderBy),
		WithDefinitionFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := w.Edges.totalCount[2][alias]
	if nodes, err := w.NamedDefinitions(alias); err == nil || hasTotalCount {
		pager, err := newDefinitionPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &DefinitionConnection{Edges: []*DefinitionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return w.QueryDefinitions().Paginate(ctx, after, first, before, last, opts...)
}

func (w *Word) Parents(ctx context.Context) (result []*Word, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = w.NamedParents(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = w.Edges.ParentsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = w.QueryParents().All(ctx)
	}
	return result, err
}

func (w *Word) Descendants(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WordOrder, where *WordWhereInput,
) (*WordConnection, error) {
	opts := []WordPaginateOption{
		WithWordOrder(orderBy),
		WithWordFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := w.Edges.totalCount[4][alias]
	if nodes, err := w.NamedDescendants(alias); err == nil || hasTotalCount {
		pager, err := newWordPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WordConnection{Edges: []*WordEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return w.QueryDescendants().Paginate(ctx, after, first, before, last, opts...)
}
