// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/devexps/go-monolithic-demo/app/admin/service/internal/data/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 1)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreateBy:   {Type: field.TypeUint32, Column: user.FieldCreateBy},
			user.FieldCreateTime: {Type: field.TypeTime, Column: user.FieldCreateTime},
			user.FieldUpdateTime: {Type: field.TypeTime, Column: user.FieldUpdateTime},
			user.FieldDeleteTime: {Type: field.TypeTime, Column: user.FieldDeleteTime},
			user.FieldStatus:     {Type: field.TypeEnum, Column: user.FieldStatus},
			user.FieldUserName:   {Type: field.TypeString, Column: user.FieldUserName},
			user.FieldPassword:   {Type: field.TypeString, Column: user.FieldPassword},
			user.FieldNickName:   {Type: field.TypeString, Column: user.FieldNickName},
			user.FieldRealName:   {Type: field.TypeString, Column: user.FieldRealName},
			user.FieldEmail:      {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldPhone:      {Type: field.TypeString, Column: user.FieldPhone},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *UserFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreateBy applies the entql uint32 predicate on the create_by field.
func (f *UserFilter) WhereCreateBy(p entql.Uint32P) {
	f.Where(p.Field(user.FieldCreateBy))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *UserFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *UserFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldUpdateTime))
}

// WhereDeleteTime applies the entql time.Time predicate on the delete_time field.
func (f *UserFilter) WhereDeleteTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldDeleteTime))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *UserFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(user.FieldStatus))
}

// WhereUserName applies the entql string predicate on the user_name field.
func (f *UserFilter) WhereUserName(p entql.StringP) {
	f.Where(p.Field(user.FieldUserName))
}

// WherePassword applies the entql string predicate on the password field.
func (f *UserFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(user.FieldPassword))
}

// WhereNickName applies the entql string predicate on the nick_name field.
func (f *UserFilter) WhereNickName(p entql.StringP) {
	f.Where(p.Field(user.FieldNickName))
}

// WhereRealName applies the entql string predicate on the real_name field.
func (f *UserFilter) WhereRealName(p entql.StringP) {
	f.Where(p.Field(user.FieldRealName))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WherePhone applies the entql string predicate on the phone field.
func (f *UserFilter) WherePhone(p entql.StringP) {
	f.Where(p.Field(user.FieldPhone))
}
