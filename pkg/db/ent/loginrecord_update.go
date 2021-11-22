// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/loginrecord"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// LoginRecordUpdate is the builder for updating LoginRecord entities.
type LoginRecordUpdate struct {
	config
	hooks    []Hook
	mutation *LoginRecordMutation
}

// Where appends a list predicates to the LoginRecordUpdate builder.
func (lru *LoginRecordUpdate) Where(ps ...predicate.LoginRecord) *LoginRecordUpdate {
	lru.mutation.Where(ps...)
	return lru
}

// SetUserID sets the "user_id" field.
func (lru *LoginRecordUpdate) SetUserID(u uuid.UUID) *LoginRecordUpdate {
	lru.mutation.SetUserID(u)
	return lru
}

// SetAppID sets the "app_id" field.
func (lru *LoginRecordUpdate) SetAppID(u uuid.UUID) *LoginRecordUpdate {
	lru.mutation.SetAppID(u)
	return lru
}

// SetLoginTime sets the "login_time" field.
func (lru *LoginRecordUpdate) SetLoginTime(u uint32) *LoginRecordUpdate {
	lru.mutation.ResetLoginTime()
	lru.mutation.SetLoginTime(u)
	return lru
}

// SetNillableLoginTime sets the "login_time" field if the given value is not nil.
func (lru *LoginRecordUpdate) SetNillableLoginTime(u *uint32) *LoginRecordUpdate {
	if u != nil {
		lru.SetLoginTime(*u)
	}
	return lru
}

// AddLoginTime adds u to the "login_time" field.
func (lru *LoginRecordUpdate) AddLoginTime(u uint32) *LoginRecordUpdate {
	lru.mutation.AddLoginTime(u)
	return lru
}

// SetIP sets the "ip" field.
func (lru *LoginRecordUpdate) SetIP(s string) *LoginRecordUpdate {
	lru.mutation.SetIP(s)
	return lru
}

// SetLocation sets the "location" field.
func (lru *LoginRecordUpdate) SetLocation(s string) *LoginRecordUpdate {
	lru.mutation.SetLocation(s)
	return lru
}

// SetLat sets the "lat" field.
func (lru *LoginRecordUpdate) SetLat(f float64) *LoginRecordUpdate {
	lru.mutation.ResetLat()
	lru.mutation.SetLat(f)
	return lru
}

// AddLat adds f to the "lat" field.
func (lru *LoginRecordUpdate) AddLat(f float64) *LoginRecordUpdate {
	lru.mutation.AddLat(f)
	return lru
}

// SetLon sets the "lon" field.
func (lru *LoginRecordUpdate) SetLon(f float64) *LoginRecordUpdate {
	lru.mutation.ResetLon()
	lru.mutation.SetLon(f)
	return lru
}

// AddLon adds f to the "lon" field.
func (lru *LoginRecordUpdate) AddLon(f float64) *LoginRecordUpdate {
	lru.mutation.AddLon(f)
	return lru
}

// SetTimezone sets the "timezone" field.
func (lru *LoginRecordUpdate) SetTimezone(s string) *LoginRecordUpdate {
	lru.mutation.SetTimezone(s)
	return lru
}

// Mutation returns the LoginRecordMutation object of the builder.
func (lru *LoginRecordUpdate) Mutation() *LoginRecordMutation {
	return lru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lru *LoginRecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lru.hooks) == 0 {
		affected, err = lru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoginRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lru.mutation = mutation
			affected, err = lru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lru.hooks) - 1; i >= 0; i-- {
			if lru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lru *LoginRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := lru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lru *LoginRecordUpdate) Exec(ctx context.Context) error {
	_, err := lru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lru *LoginRecordUpdate) ExecX(ctx context.Context) {
	if err := lru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lru *LoginRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   loginrecord.Table,
			Columns: loginrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: loginrecord.FieldID,
			},
		},
	}
	if ps := lru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: loginrecord.FieldUserID,
		})
	}
	if value, ok := lru.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: loginrecord.FieldAppID,
		})
	}
	if value, ok := lru.mutation.LoginTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: loginrecord.FieldLoginTime,
		})
	}
	if value, ok := lru.mutation.AddedLoginTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: loginrecord.FieldLoginTime,
		})
	}
	if value, ok := lru.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: loginrecord.FieldIP,
		})
	}
	if value, ok := lru.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: loginrecord.FieldLocation,
		})
	}
	if value, ok := lru.mutation.Lat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: loginrecord.FieldLat,
		})
	}
	if value, ok := lru.mutation.AddedLat(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: loginrecord.FieldLat,
		})
	}
	if value, ok := lru.mutation.Lon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: loginrecord.FieldLon,
		})
	}
	if value, ok := lru.mutation.AddedLon(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: loginrecord.FieldLon,
		})
	}
	if value, ok := lru.mutation.Timezone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: loginrecord.FieldTimezone,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LoginRecordUpdateOne is the builder for updating a single LoginRecord entity.
type LoginRecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoginRecordMutation
}

// SetUserID sets the "user_id" field.
func (lruo *LoginRecordUpdateOne) SetUserID(u uuid.UUID) *LoginRecordUpdateOne {
	lruo.mutation.SetUserID(u)
	return lruo
}

// SetAppID sets the "app_id" field.
func (lruo *LoginRecordUpdateOne) SetAppID(u uuid.UUID) *LoginRecordUpdateOne {
	lruo.mutation.SetAppID(u)
	return lruo
}

// SetLoginTime sets the "login_time" field.
func (lruo *LoginRecordUpdateOne) SetLoginTime(u uint32) *LoginRecordUpdateOne {
	lruo.mutation.ResetLoginTime()
	lruo.mutation.SetLoginTime(u)
	return lruo
}

// SetNillableLoginTime sets the "login_time" field if the given value is not nil.
func (lruo *LoginRecordUpdateOne) SetNillableLoginTime(u *uint32) *LoginRecordUpdateOne {
	if u != nil {
		lruo.SetLoginTime(*u)
	}
	return lruo
}

// AddLoginTime adds u to the "login_time" field.
func (lruo *LoginRecordUpdateOne) AddLoginTime(u uint32) *LoginRecordUpdateOne {
	lruo.mutation.AddLoginTime(u)
	return lruo
}

// SetIP sets the "ip" field.
func (lruo *LoginRecordUpdateOne) SetIP(s string) *LoginRecordUpdateOne {
	lruo.mutation.SetIP(s)
	return lruo
}

// SetLocation sets the "location" field.
func (lruo *LoginRecordUpdateOne) SetLocation(s string) *LoginRecordUpdateOne {
	lruo.mutation.SetLocation(s)
	return lruo
}

// SetLat sets the "lat" field.
func (lruo *LoginRecordUpdateOne) SetLat(f float64) *LoginRecordUpdateOne {
	lruo.mutation.ResetLat()
	lruo.mutation.SetLat(f)
	return lruo
}

// AddLat adds f to the "lat" field.
func (lruo *LoginRecordUpdateOne) AddLat(f float64) *LoginRecordUpdateOne {
	lruo.mutation.AddLat(f)
	return lruo
}

// SetLon sets the "lon" field.
func (lruo *LoginRecordUpdateOne) SetLon(f float64) *LoginRecordUpdateOne {
	lruo.mutation.ResetLon()
	lruo.mutation.SetLon(f)
	return lruo
}

// AddLon adds f to the "lon" field.
func (lruo *LoginRecordUpdateOne) AddLon(f float64) *LoginRecordUpdateOne {
	lruo.mutation.AddLon(f)
	return lruo
}

// SetTimezone sets the "timezone" field.
func (lruo *LoginRecordUpdateOne) SetTimezone(s string) *LoginRecordUpdateOne {
	lruo.mutation.SetTimezone(s)
	return lruo
}

// Mutation returns the LoginRecordMutation object of the builder.
func (lruo *LoginRecordUpdateOne) Mutation() *LoginRecordMutation {
	return lruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lruo *LoginRecordUpdateOne) Select(field string, fields ...string) *LoginRecordUpdateOne {
	lruo.fields = append([]string{field}, fields...)
	return lruo
}

// Save executes the query and returns the updated LoginRecord entity.
func (lruo *LoginRecordUpdateOne) Save(ctx context.Context) (*LoginRecord, error) {
	var (
		err  error
		node *LoginRecord
	)
	if len(lruo.hooks) == 0 {
		node, err = lruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoginRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lruo.mutation = mutation
			node, err = lruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lruo.hooks) - 1; i >= 0; i-- {
			if lruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lruo *LoginRecordUpdateOne) SaveX(ctx context.Context) *LoginRecord {
	node, err := lruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lruo *LoginRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := lruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lruo *LoginRecordUpdateOne) ExecX(ctx context.Context) {
	if err := lruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lruo *LoginRecordUpdateOne) sqlSave(ctx context.Context) (_node *LoginRecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   loginrecord.Table,
			Columns: loginrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: loginrecord.FieldID,
			},
		},
	}
	id, ok := lruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing LoginRecord.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := lruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loginrecord.FieldID)
		for _, f := range fields {
			if !loginrecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != loginrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: loginrecord.FieldUserID,
		})
	}
	if value, ok := lruo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: loginrecord.FieldAppID,
		})
	}
	if value, ok := lruo.mutation.LoginTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: loginrecord.FieldLoginTime,
		})
	}
	if value, ok := lruo.mutation.AddedLoginTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: loginrecord.FieldLoginTime,
		})
	}
	if value, ok := lruo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: loginrecord.FieldIP,
		})
	}
	if value, ok := lruo.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: loginrecord.FieldLocation,
		})
	}
	if value, ok := lruo.mutation.Lat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: loginrecord.FieldLat,
		})
	}
	if value, ok := lruo.mutation.AddedLat(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: loginrecord.FieldLat,
		})
	}
	if value, ok := lruo.mutation.Lon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: loginrecord.FieldLon,
		})
	}
	if value, ok := lruo.mutation.AddedLon(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: loginrecord.FieldLon,
		})
	}
	if value, ok := lruo.mutation.Timezone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: loginrecord.FieldTimezone,
		})
	}
	_node = &LoginRecord{config: lruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loginrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
