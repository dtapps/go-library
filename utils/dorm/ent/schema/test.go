package schema

import "entgo.io/ent"

// Test holds the schema definition for the Test entity.
type Test struct {
	ent.Schema
}

// Fields of the Test.
func (Test) Fields() []ent.Field {
	return nil
}

// Edges of the Test.
func (Test) Edges() []ent.Edge {
	return nil
}
