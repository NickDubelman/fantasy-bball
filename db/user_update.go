// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NickDubelman/fantasy-bball/db/predicate"
	"github.com/NickDubelman/fantasy-bball/db/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where adds a new predicate for the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPicture sets the "picture" field.
func (uu *UserUpdate) SetPicture(s string) *UserUpdate {
	uu.mutation.SetPicture(s)
	return uu
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePicture(s *string) *UserUpdate {
	if s != nil {
		uu.SetPicture(*s)
	}
	return uu
}

// ClearPicture clears the value of the "picture" field.
func (uu *UserUpdate) ClearPicture() *UserUpdate {
	uu.mutation.ClearPicture()
	return uu
}

// SetJoined sets the "joined" field.
func (uu *UserUpdate) SetJoined(t time.Time) *UserUpdate {
	uu.mutation.SetJoined(t)
	return uu
}

// SetNillableJoined sets the "joined" field if the given value is not nil.
func (uu *UserUpdate) SetNillableJoined(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetJoined(*t)
	}
	return uu
}

// ClearJoined clears the value of the "joined" field.
func (uu *UserUpdate) ClearJoined() *UserUpdate {
	uu.mutation.ClearJoined()
	return uu
}

// SetLastActive sets the "lastActive" field.
func (uu *UserUpdate) SetLastActive(t time.Time) *UserUpdate {
	uu.mutation.SetLastActive(t)
	return uu
}

// SetNillableLastActive sets the "lastActive" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastActive(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLastActive(*t)
	}
	return uu
}

// ClearLastActive clears the value of the "lastActive" field.
func (uu *UserUpdate) ClearLastActive() *UserUpdate {
	uu.mutation.ClearLastActive()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Picture(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPicture,
		})
	}
	if uu.mutation.PictureCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPicture,
		})
	}
	if value, ok := uu.mutation.Joined(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldJoined,
		})
	}
	if uu.mutation.JoinedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldJoined,
		})
	}
	if value, ok := uu.mutation.LastActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastActive,
		})
	}
	if uu.mutation.LastActiveCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldLastActive,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPicture sets the "picture" field.
func (uuo *UserUpdateOne) SetPicture(s string) *UserUpdateOne {
	uuo.mutation.SetPicture(s)
	return uuo
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePicture(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPicture(*s)
	}
	return uuo
}

// ClearPicture clears the value of the "picture" field.
func (uuo *UserUpdateOne) ClearPicture() *UserUpdateOne {
	uuo.mutation.ClearPicture()
	return uuo
}

// SetJoined sets the "joined" field.
func (uuo *UserUpdateOne) SetJoined(t time.Time) *UserUpdateOne {
	uuo.mutation.SetJoined(t)
	return uuo
}

// SetNillableJoined sets the "joined" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableJoined(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetJoined(*t)
	}
	return uuo
}

// ClearJoined clears the value of the "joined" field.
func (uuo *UserUpdateOne) ClearJoined() *UserUpdateOne {
	uuo.mutation.ClearJoined()
	return uuo
}

// SetLastActive sets the "lastActive" field.
func (uuo *UserUpdateOne) SetLastActive(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastActive(t)
	return uuo
}

// SetNillableLastActive sets the "lastActive" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastActive(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLastActive(*t)
	}
	return uuo
}

// ClearLastActive clears the value of the "lastActive" field.
func (uuo *UserUpdateOne) ClearLastActive() *UserUpdateOne {
	uuo.mutation.ClearLastActive()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Picture(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPicture,
		})
	}
	if uuo.mutation.PictureCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPicture,
		})
	}
	if value, ok := uuo.mutation.Joined(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldJoined,
		})
	}
	if uuo.mutation.JoinedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldJoined,
		})
	}
	if value, ok := uuo.mutation.LastActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastActive,
		})
	}
	if uuo.mutation.LastActiveCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldLastActive,
		})
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
