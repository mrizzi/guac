package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SecurityAdvisory holds the schema definition for the SecurityAdvisory entity.
// This wraps OSV, GHSA and CVE nodes. This is synonymous with the GQL Vulnerability interface.
type SecurityAdvisory struct {
	ent.Schema
}

// Fields of the GitHubSecurityAdvisory.
func (SecurityAdvisory) Fields() []ent.Field {
	return []ent.Field{
		field.String("ghsa_id").Optional().Nillable().Comment("GHSA represents GitHub security advisories"),
		field.String("cve_id").Optional().Nillable().Comment("CVE represents Common Vulnerabilities and Exposures"),
		field.Int("cve_year").Optional().Nillable().Comment("CVE year"),
		field.String("osv_id").Optional().Nillable().Comment("OSV represents Open Source Vulnerabilities"),
	}
}

// Edges of the GitHubSecurityAdvisory.
func (SecurityAdvisory) Edges() []ent.Edge {
	return nil
}

// Indexes of the GitHubSecurityAdvisory.
func (SecurityAdvisory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ghsa_id").Unique().Annotations(
			entsql.IndexWhere("osv_id IS NULL AND cve_id IS NULL AND ghsa_id IS NOT NULL"),
		),
		index.Fields("cve_id").Unique().Annotations(
			entsql.IndexWhere("osv_id IS NULL AND cve_id IS NOT NULL AND ghsa_id IS NULL"),
		),
		index.Fields("osv_id").Unique().Annotations(
			entsql.IndexWhere("osv_id IS NOT NULL AND cve_id IS NULL AND ghsa_id IS NULL"),
		),
	}
}