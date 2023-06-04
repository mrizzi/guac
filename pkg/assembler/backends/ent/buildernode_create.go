// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/buildernode"
)

// BuilderNodeCreate is the builder for creating a BuilderNode entity.
type BuilderNodeCreate struct {
	config
	mutation *BuilderNodeMutation
	hooks    []Hook
}

// SetURI sets the "uri" field.
func (bnc *BuilderNodeCreate) SetURI(s string) *BuilderNodeCreate {
	bnc.mutation.SetURI(s)
	return bnc
}

// Mutation returns the BuilderNodeMutation object of the builder.
func (bnc *BuilderNodeCreate) Mutation() *BuilderNodeMutation {
	return bnc.mutation
}

// Save creates the BuilderNode in the database.
func (bnc *BuilderNodeCreate) Save(ctx context.Context) (*BuilderNode, error) {
	return withHooks(ctx, bnc.sqlSave, bnc.mutation, bnc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bnc *BuilderNodeCreate) SaveX(ctx context.Context) *BuilderNode {
	v, err := bnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bnc *BuilderNodeCreate) Exec(ctx context.Context) error {
	_, err := bnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bnc *BuilderNodeCreate) ExecX(ctx context.Context) {
	if err := bnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bnc *BuilderNodeCreate) check() error {
	if _, ok := bnc.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "BuilderNode.uri"`)}
	}
	return nil
}

func (bnc *BuilderNodeCreate) sqlSave(ctx context.Context) (*BuilderNode, error) {
	if err := bnc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bnc.mutation.id = &_node.ID
	bnc.mutation.done = true
	return _node, nil
}

func (bnc *BuilderNodeCreate) createSpec() (*BuilderNode, *sqlgraph.CreateSpec) {
	var (
		_node = &BuilderNode{config: bnc.config}
		_spec = sqlgraph.NewCreateSpec(buildernode.Table, sqlgraph.NewFieldSpec(buildernode.FieldID, field.TypeInt))
	)
	if value, ok := bnc.mutation.URI(); ok {
		_spec.SetField(buildernode.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	return _node, _spec
}

// BuilderNodeCreateBulk is the builder for creating many BuilderNode entities in bulk.
type BuilderNodeCreateBulk struct {
	config
	builders []*BuilderNodeCreate
}

// Save creates the BuilderNode entities in the database.
func (bncb *BuilderNodeCreateBulk) Save(ctx context.Context) ([]*BuilderNode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bncb.builders))
	nodes := make([]*BuilderNode, len(bncb.builders))
	mutators := make([]Mutator, len(bncb.builders))
	for i := range bncb.builders {
		func(i int, root context.Context) {
			builder := bncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BuilderNodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bncb *BuilderNodeCreateBulk) SaveX(ctx context.Context) []*BuilderNode {
	v, err := bncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bncb *BuilderNodeCreateBulk) Exec(ctx context.Context) error {
	_, err := bncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bncb *BuilderNodeCreateBulk) ExecX(ctx context.Context) {
	if err := bncb.Exec(ctx); err != nil {
		panic(err)
	}
}
