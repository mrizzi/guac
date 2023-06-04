// Code generated by ent, DO NOT EDIT.

package artifact

import (
	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Artifact {
	return predicate.Artifact(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Artifact {
	return predicate.Artifact(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Artifact {
	return predicate.Artifact(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Artifact {
	return predicate.Artifact(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Artifact {
	return predicate.Artifact(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Artifact {
	return predicate.Artifact(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Artifact {
	return predicate.Artifact(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Artifact {
	return predicate.Artifact(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Artifact {
	return predicate.Artifact(sql.FieldLTE(FieldID, id))
}

// Algorithm applies equality check predicate on the "algorithm" field. It's identical to AlgorithmEQ.
func Algorithm(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldEQ(FieldAlgorithm, v))
}

// Digest applies equality check predicate on the "digest" field. It's identical to DigestEQ.
func Digest(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldEQ(FieldDigest, v))
}

// AlgorithmEQ applies the EQ predicate on the "algorithm" field.
func AlgorithmEQ(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldEQ(FieldAlgorithm, v))
}

// AlgorithmNEQ applies the NEQ predicate on the "algorithm" field.
func AlgorithmNEQ(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldNEQ(FieldAlgorithm, v))
}

// AlgorithmIn applies the In predicate on the "algorithm" field.
func AlgorithmIn(vs ...string) predicate.Artifact {
	return predicate.Artifact(sql.FieldIn(FieldAlgorithm, vs...))
}

// AlgorithmNotIn applies the NotIn predicate on the "algorithm" field.
func AlgorithmNotIn(vs ...string) predicate.Artifact {
	return predicate.Artifact(sql.FieldNotIn(FieldAlgorithm, vs...))
}

// AlgorithmGT applies the GT predicate on the "algorithm" field.
func AlgorithmGT(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldGT(FieldAlgorithm, v))
}

// AlgorithmGTE applies the GTE predicate on the "algorithm" field.
func AlgorithmGTE(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldGTE(FieldAlgorithm, v))
}

// AlgorithmLT applies the LT predicate on the "algorithm" field.
func AlgorithmLT(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldLT(FieldAlgorithm, v))
}

// AlgorithmLTE applies the LTE predicate on the "algorithm" field.
func AlgorithmLTE(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldLTE(FieldAlgorithm, v))
}

// AlgorithmContains applies the Contains predicate on the "algorithm" field.
func AlgorithmContains(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldContains(FieldAlgorithm, v))
}

// AlgorithmHasPrefix applies the HasPrefix predicate on the "algorithm" field.
func AlgorithmHasPrefix(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldHasPrefix(FieldAlgorithm, v))
}

// AlgorithmHasSuffix applies the HasSuffix predicate on the "algorithm" field.
func AlgorithmHasSuffix(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldHasSuffix(FieldAlgorithm, v))
}

// AlgorithmEqualFold applies the EqualFold predicate on the "algorithm" field.
func AlgorithmEqualFold(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldEqualFold(FieldAlgorithm, v))
}

// AlgorithmContainsFold applies the ContainsFold predicate on the "algorithm" field.
func AlgorithmContainsFold(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldContainsFold(FieldAlgorithm, v))
}

// DigestEQ applies the EQ predicate on the "digest" field.
func DigestEQ(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldEQ(FieldDigest, v))
}

// DigestNEQ applies the NEQ predicate on the "digest" field.
func DigestNEQ(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldNEQ(FieldDigest, v))
}

// DigestIn applies the In predicate on the "digest" field.
func DigestIn(vs ...string) predicate.Artifact {
	return predicate.Artifact(sql.FieldIn(FieldDigest, vs...))
}

// DigestNotIn applies the NotIn predicate on the "digest" field.
func DigestNotIn(vs ...string) predicate.Artifact {
	return predicate.Artifact(sql.FieldNotIn(FieldDigest, vs...))
}

// DigestGT applies the GT predicate on the "digest" field.
func DigestGT(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldGT(FieldDigest, v))
}

// DigestGTE applies the GTE predicate on the "digest" field.
func DigestGTE(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldGTE(FieldDigest, v))
}

// DigestLT applies the LT predicate on the "digest" field.
func DigestLT(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldLT(FieldDigest, v))
}

// DigestLTE applies the LTE predicate on the "digest" field.
func DigestLTE(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldLTE(FieldDigest, v))
}

// DigestContains applies the Contains predicate on the "digest" field.
func DigestContains(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldContains(FieldDigest, v))
}

// DigestHasPrefix applies the HasPrefix predicate on the "digest" field.
func DigestHasPrefix(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldHasPrefix(FieldDigest, v))
}

// DigestHasSuffix applies the HasSuffix predicate on the "digest" field.
func DigestHasSuffix(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldHasSuffix(FieldDigest, v))
}

// DigestEqualFold applies the EqualFold predicate on the "digest" field.
func DigestEqualFold(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldEqualFold(FieldDigest, v))
}

// DigestContainsFold applies the ContainsFold predicate on the "digest" field.
func DigestContainsFold(v string) predicate.Artifact {
	return predicate.Artifact(sql.FieldContainsFold(FieldDigest, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Artifact) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Artifact) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Artifact) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		p(s.Not())
	})
}
