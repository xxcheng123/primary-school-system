// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xxcheng123/primary-school-system/ent/article"
	"github.com/xxcheng123/primary-school-system/ent/predicate"
	"github.com/xxcheng123/primary-school-system/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedAt(t)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedAt(*t)
	}
	return uu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uu *UserUpdate) ClearDeletedAt() *UserUpdate {
	uu.mutation.ClearDeletedAt()
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(i int64) *UserUpdate {
	uu.mutation.ResetStatus()
	uu.mutation.SetStatus(i)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(i *int64) *UserUpdate {
	if i != nil {
		uu.SetStatus(*i)
	}
	return uu
}

// AddStatus adds i to the "status" field.
func (uu *UserUpdate) AddStatus(i int64) *UserUpdate {
	uu.mutation.AddStatus(i)
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(i int64) *UserUpdate {
	uu.mutation.ResetRole()
	uu.mutation.SetRole(i)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(i *int64) *UserUpdate {
	if i != nil {
		uu.SetRole(*i)
	}
	return uu
}

// AddRole adds i to the "role" field.
func (uu *UserUpdate) AddRole(i int64) *UserUpdate {
	uu.mutation.AddRole(i)
	return uu
}

// AddArticleIDs adds the "article" edge to the Article entity by IDs.
func (uu *UserUpdate) AddArticleIDs(ids ...int64) *UserUpdate {
	uu.mutation.AddArticleIDs(ids...)
	return uu
}

// AddArticle adds the "article" edges to the Article entity.
func (uu *UserUpdate) AddArticle(a ...*Article) *UserUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddArticleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearArticle clears all "article" edges to the Article entity.
func (uu *UserUpdate) ClearArticle() *UserUpdate {
	uu.mutation.ClearArticle()
	return uu
}

// RemoveArticleIDs removes the "article" edge to Article entities by IDs.
func (uu *UserUpdate) RemoveArticleIDs(ids ...int64) *UserUpdate {
	uu.mutation.RemoveArticleIDs(ids...)
	return uu
}

// RemoveArticle removes "article" edges to Article entities.
func (uu *UserUpdate) RemoveArticle(a ...*Article) *UserUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveArticleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if err := uu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
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

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() error {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		if user.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uu.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeInt64, value)
	}
	if uu.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticleTable,
			Columns: []string{user.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedArticleIDs(); len(nodes) > 0 && !uu.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticleTable,
			Columns: []string{user.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticleTable,
			Columns: []string{user.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedAt(t)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedAt(*t)
	}
	return uuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uuo *UserUpdateOne) ClearDeletedAt() *UserUpdateOne {
	uuo.mutation.ClearDeletedAt()
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(i int64) *UserUpdateOne {
	uuo.mutation.ResetStatus()
	uuo.mutation.SetStatus(i)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetStatus(*i)
	}
	return uuo
}

// AddStatus adds i to the "status" field.
func (uuo *UserUpdateOne) AddStatus(i int64) *UserUpdateOne {
	uuo.mutation.AddStatus(i)
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(i int64) *UserUpdateOne {
	uuo.mutation.ResetRole()
	uuo.mutation.SetRole(i)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetRole(*i)
	}
	return uuo
}

// AddRole adds i to the "role" field.
func (uuo *UserUpdateOne) AddRole(i int64) *UserUpdateOne {
	uuo.mutation.AddRole(i)
	return uuo
}

// AddArticleIDs adds the "article" edge to the Article entity by IDs.
func (uuo *UserUpdateOne) AddArticleIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.AddArticleIDs(ids...)
	return uuo
}

// AddArticle adds the "article" edges to the Article entity.
func (uuo *UserUpdateOne) AddArticle(a ...*Article) *UserUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddArticleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearArticle clears all "article" edges to the Article entity.
func (uuo *UserUpdateOne) ClearArticle() *UserUpdateOne {
	uuo.mutation.ClearArticle()
	return uuo
}

// RemoveArticleIDs removes the "article" edge to Article entities by IDs.
func (uuo *UserUpdateOne) RemoveArticleIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.RemoveArticleIDs(ids...)
	return uuo
}

// RemoveArticle removes "article" edges to Article entities.
func (uuo *UserUpdateOne) RemoveArticle(a ...*Article) *UserUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveArticleIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if err := uuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
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

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() error {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		if user.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uuo.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeInt64, value)
	}
	if uuo.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticleTable,
			Columns: []string{user.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedArticleIDs(); len(nodes) > 0 && !uuo.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticleTable,
			Columns: []string{user.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticleTable,
			Columns: []string{user.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
