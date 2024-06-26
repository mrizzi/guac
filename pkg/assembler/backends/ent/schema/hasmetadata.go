//
// Copyright 2023 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// HasMetadata holds the schema definition for the HasMetadata entity.
type HasMetadata struct {
	ent.Schema
}

// Fields of the HasMetadata.
func (HasMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(getUUIDv7).
			Unique().
			Immutable(),
		field.UUID("source_id", getUUIDv7()).Optional().Nillable(),
		field.UUID("package_version_id", getUUIDv7()).Optional().Nillable(),
		field.UUID("package_name_id", getUUIDv7()).Optional().Nillable(),
		field.UUID("artifact_id", getUUIDv7()).Optional().Nillable(),
		field.Time("timestamp"),
		field.String("key"),
		field.String("value"),
		field.String("justification"),
		field.String("origin"),
		field.String("collector"),
		field.String("document_ref"),
	}
}

// Edges of the HasMetadata.
func (HasMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("source", SourceName.Type).Unique().Field("source_id"),
		edge.To("package_version", PackageVersion.Type).Unique().Field("package_version_id"),
		edge.To("all_versions", PackageName.Type).Unique().Field("package_name_id"),
		edge.To("artifact", Artifact.Type).Unique().Field("artifact_id"),
	}
}

func (HasMetadata) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key", "value", "justification", "origin", "collector", "timestamp", "document_ref", "source_id").Unique().
			Annotations(entsql.IndexWhere("source_id IS NOT NULL AND package_version_id IS NULL AND package_name_id IS NULL AND artifact_id IS NULL")).
			StorageKey("has_metadata_source_id"),
		index.Fields("key", "value", "justification", "origin", "collector", "timestamp", "document_ref", "package_version_id").Unique().
			Annotations(entsql.IndexWhere("source_id IS NULL AND package_version_id IS NOT NULL AND package_name_id IS NULL AND artifact_id IS NULL")).
			StorageKey("has_metadata_package_version_id"),
		index.Fields("key", "value", "justification", "origin", "collector", "timestamp", "document_ref", "package_name_id").Unique().
			Annotations(entsql.IndexWhere("source_id IS NULL AND package_version_id IS NULL AND package_name_id IS NOT NULL AND artifact_id IS NULL")).
			StorageKey("has_metadata_package_name_id"),
		index.Fields("key", "value", "justification", "origin", "collector", "timestamp", "document_ref", "artifact_id").Unique().
			Annotations(entsql.IndexWhere("source_id IS NULL AND package_version_id IS NULL AND package_name_id IS NULL AND artifact_id IS NOT NULL")).
			StorageKey("has_metadata_artifact_id"),
	}
}
