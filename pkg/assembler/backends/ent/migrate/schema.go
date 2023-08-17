// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArtifactsColumns holds the columns for the "artifacts" table.
	ArtifactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "algorithm", Type: field.TypeString},
		{Name: "digest", Type: field.TypeString},
	}
	// ArtifactsTable holds the schema information for the "artifacts" table.
	ArtifactsTable = &schema.Table{
		Name:       "artifacts",
		Columns:    ArtifactsColumns,
		PrimaryKey: []*schema.Column{ArtifactsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "artifact_algorithm",
				Unique:  false,
				Columns: []*schema.Column{ArtifactsColumns[1]},
			},
			{
				Name:    "artifact_digest",
				Unique:  true,
				Columns: []*schema.Column{ArtifactsColumns[2]},
			},
		},
	}
	// BillOfMaterialsColumns holds the columns for the "bill_of_materials" table.
	BillOfMaterialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uri", Type: field.TypeString},
		{Name: "algorithm", Type: field.TypeString},
		{Name: "digest", Type: field.TypeString},
		{Name: "download_location", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "package_id", Type: field.TypeInt, Nullable: true},
		{Name: "artifact_id", Type: field.TypeInt, Nullable: true},
	}
	// BillOfMaterialsTable holds the schema information for the "bill_of_materials" table.
	BillOfMaterialsTable = &schema.Table{
		Name:       "bill_of_materials",
		Columns:    BillOfMaterialsColumns,
		PrimaryKey: []*schema.Column{BillOfMaterialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bill_of_materials_package_versions_package",
				Columns:    []*schema.Column{BillOfMaterialsColumns[7]},
				RefColumns: []*schema.Column{PackageVersionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bill_of_materials_artifacts_artifact",
				Columns:    []*schema.Column{BillOfMaterialsColumns[8]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "sbom_unique_package",
				Unique:  true,
				Columns: []*schema.Column{BillOfMaterialsColumns[2], BillOfMaterialsColumns[3], BillOfMaterialsColumns[1], BillOfMaterialsColumns[4], BillOfMaterialsColumns[7]},
				Annotation: &entsql.IndexAnnotation{
					Where: "package_id IS NOT NULL AND artifact_id IS NULL",
				},
			},
			{
				Name:    "sbom_unique_artifact",
				Unique:  true,
				Columns: []*schema.Column{BillOfMaterialsColumns[2], BillOfMaterialsColumns[3], BillOfMaterialsColumns[1], BillOfMaterialsColumns[4], BillOfMaterialsColumns[8]},
				Annotation: &entsql.IndexAnnotation{
					Where: "package_id IS NULL AND artifact_id IS NOT NULL",
				},
			},
		},
	}
	// BuildersColumns holds the columns for the "builders" table.
	BuildersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uri", Type: field.TypeString, Unique: true},
	}
	// BuildersTable holds the schema information for the "builders" table.
	BuildersTable = &schema.Table{
		Name:       "builders",
		Columns:    BuildersColumns,
		PrimaryKey: []*schema.Column{BuildersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "builder_uri",
				Unique:  true,
				Columns: []*schema.Column{BuildersColumns[1]},
			},
		},
	}
	// CertificationsColumns holds the columns for the "certifications" table.
	CertificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"GOOD", "BAD"}, Default: "GOOD"},
		{Name: "justification", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "source_id", Type: field.TypeInt, Nullable: true},
		{Name: "package_version_id", Type: field.TypeInt, Nullable: true},
		{Name: "package_name_id", Type: field.TypeInt, Nullable: true},
		{Name: "artifact_id", Type: field.TypeInt, Nullable: true},
	}
	// CertificationsTable holds the schema information for the "certifications" table.
	CertificationsTable = &schema.Table{
		Name:       "certifications",
		Columns:    CertificationsColumns,
		PrimaryKey: []*schema.Column{CertificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "certifications_source_names_source",
				Columns:    []*schema.Column{CertificationsColumns[5]},
				RefColumns: []*schema.Column{SourceNamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "certifications_package_versions_package_version",
				Columns:    []*schema.Column{CertificationsColumns[6]},
				RefColumns: []*schema.Column{PackageVersionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "certifications_package_names_all_versions",
				Columns:    []*schema.Column{CertificationsColumns[7]},
				RefColumns: []*schema.Column{PackageNamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "certifications_artifacts_artifact",
				Columns:    []*schema.Column{CertificationsColumns[8]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "certification_type_justification_origin_collector_source_id",
				Unique:  true,
				Columns: []*schema.Column{CertificationsColumns[1], CertificationsColumns[2], CertificationsColumns[3], CertificationsColumns[4], CertificationsColumns[5]},
				Annotation: &entsql.IndexAnnotation{
					Where: "source_id IS NOT NULL AND package_version_id IS NULL AND package_name_id IS NULL AND artifact_id IS NULL",
				},
			},
			{
				Name:    "certification_type_justification_origin_collector_package_version_id",
				Unique:  true,
				Columns: []*schema.Column{CertificationsColumns[1], CertificationsColumns[2], CertificationsColumns[3], CertificationsColumns[4], CertificationsColumns[6]},
				Annotation: &entsql.IndexAnnotation{
					Where: "source_id IS NULL AND package_version_id IS NOT NULL AND package_name_id IS NULL AND artifact_id IS NULL",
				},
			},
			{
				Name:    "certification_type_justification_origin_collector_package_name_id",
				Unique:  true,
				Columns: []*schema.Column{CertificationsColumns[1], CertificationsColumns[2], CertificationsColumns[3], CertificationsColumns[4], CertificationsColumns[7]},
				Annotation: &entsql.IndexAnnotation{
					Where: "source_id IS NULL AND package_version_id IS NULL AND package_name_id IS NOT NULL AND artifact_id IS NULL",
				},
			},
			{
				Name:    "certification_type_justification_origin_collector_artifact_id",
				Unique:  true,
				Columns: []*schema.Column{CertificationsColumns[1], CertificationsColumns[2], CertificationsColumns[3], CertificationsColumns[4], CertificationsColumns[8]},
				Annotation: &entsql.IndexAnnotation{
					Where: "source_id IS NULL AND package_version_id IS NULL AND package_name_id IS NULL AND artifact_id IS NOT NULL",
				},
			},
		},
	}
	// CertifyScorecardsColumns holds the columns for the "certify_scorecards" table.
	CertifyScorecardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "source_id", Type: field.TypeInt},
		{Name: "scorecard_id", Type: field.TypeInt},
	}
	// CertifyScorecardsTable holds the schema information for the "certify_scorecards" table.
	CertifyScorecardsTable = &schema.Table{
		Name:       "certify_scorecards",
		Columns:    CertifyScorecardsColumns,
		PrimaryKey: []*schema.Column{CertifyScorecardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "certify_scorecards_source_names_source",
				Columns:    []*schema.Column{CertifyScorecardsColumns[1]},
				RefColumns: []*schema.Column{SourceNamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "certify_scorecards_scorecards_certifications",
				Columns:    []*schema.Column{CertifyScorecardsColumns[2]},
				RefColumns: []*schema.Column{ScorecardsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "certifyscorecard_source_id_scorecard_id",
				Unique:  true,
				Columns: []*schema.Column{CertifyScorecardsColumns[1], CertifyScorecardsColumns[2]},
			},
		},
	}
	// CertifyVulnsColumns holds the columns for the "certify_vulns" table.
	CertifyVulnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "time_scanned", Type: field.TypeTime},
		{Name: "db_uri", Type: field.TypeString},
		{Name: "db_version", Type: field.TypeString},
		{Name: "scanner_uri", Type: field.TypeString},
		{Name: "scanner_version", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "vulnerability_id", Type: field.TypeInt, Nullable: true},
		{Name: "package_id", Type: field.TypeInt},
	}
	// CertifyVulnsTable holds the schema information for the "certify_vulns" table.
	CertifyVulnsTable = &schema.Table{
		Name:       "certify_vulns",
		Columns:    CertifyVulnsColumns,
		PrimaryKey: []*schema.Column{CertifyVulnsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "certify_vulns_vulnerabilities_vulnerability",
				Columns:    []*schema.Column{CertifyVulnsColumns[8]},
				RefColumns: []*schema.Column{VulnerabilitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "certify_vulns_package_versions_package",
				Columns:    []*schema.Column{CertifyVulnsColumns[9]},
				RefColumns: []*schema.Column{PackageVersionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "certifyvuln_db_uri_db_version_scanner_uri_scanner_version_origin_collector_vulnerability_id_package_id",
				Unique:  true,
				Columns: []*schema.Column{CertifyVulnsColumns[2], CertifyVulnsColumns[3], CertifyVulnsColumns[4], CertifyVulnsColumns[5], CertifyVulnsColumns[6], CertifyVulnsColumns[7], CertifyVulnsColumns[8], CertifyVulnsColumns[9]},
				Annotation: &entsql.IndexAnnotation{
					Where: "vulnerability_id IS NOT NULL",
				},
			},
			{
				Name:    "certifyvuln_db_uri_db_version_scanner_uri_scanner_version_origin_collector_package_id",
				Unique:  true,
				Columns: []*schema.Column{CertifyVulnsColumns[2], CertifyVulnsColumns[3], CertifyVulnsColumns[4], CertifyVulnsColumns[5], CertifyVulnsColumns[6], CertifyVulnsColumns[7], CertifyVulnsColumns[9]},
				Annotation: &entsql.IndexAnnotation{
					Where: "vulnerability_id IS NULL",
				},
			},
		},
	}
	// DependenciesColumns holds the columns for the "dependencies" table.
	DependenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "version_range", Type: field.TypeString},
		{Name: "dependency_type", Type: field.TypeEnum, Enums: []string{"UNSPECIFIED", "DIRECT", "INDIRECT"}},
		{Name: "justification", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "package_id", Type: field.TypeInt},
		{Name: "dependent_package_id", Type: field.TypeInt},
	}
	// DependenciesTable holds the schema information for the "dependencies" table.
	DependenciesTable = &schema.Table{
		Name:       "dependencies",
		Columns:    DependenciesColumns,
		PrimaryKey: []*schema.Column{DependenciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dependencies_package_versions_package",
				Columns:    []*schema.Column{DependenciesColumns[6]},
				RefColumns: []*schema.Column{PackageVersionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "dependencies_package_names_dependent_package",
				Columns:    []*schema.Column{DependenciesColumns[7]},
				RefColumns: []*schema.Column{PackageNamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "dependency_version_range_dependency_type_justification_origin_collector_package_id_dependent_package_id",
				Unique:  true,
				Columns: []*schema.Column{DependenciesColumns[1], DependenciesColumns[2], DependenciesColumns[3], DependenciesColumns[4], DependenciesColumns[5], DependenciesColumns[6], DependenciesColumns[7]},
			},
		},
	}
	// HasSourceAtsColumns holds the columns for the "has_source_ats" table.
	HasSourceAtsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "known_since", Type: field.TypeTime},
		{Name: "justification", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "package_version_id", Type: field.TypeInt, Nullable: true},
		{Name: "package_name_id", Type: field.TypeInt, Nullable: true},
		{Name: "source_id", Type: field.TypeInt},
	}
	// HasSourceAtsTable holds the schema information for the "has_source_ats" table.
	HasSourceAtsTable = &schema.Table{
		Name:       "has_source_ats",
		Columns:    HasSourceAtsColumns,
		PrimaryKey: []*schema.Column{HasSourceAtsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "has_source_ats_package_versions_package_version",
				Columns:    []*schema.Column{HasSourceAtsColumns[5]},
				RefColumns: []*schema.Column{PackageVersionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "has_source_ats_package_names_all_versions",
				Columns:    []*schema.Column{HasSourceAtsColumns[6]},
				RefColumns: []*schema.Column{PackageNamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "has_source_ats_source_names_source",
				Columns:    []*schema.Column{HasSourceAtsColumns[7]},
				RefColumns: []*schema.Column{SourceNamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "hassourceat_source_id_package_version_id_justification",
				Unique:  true,
				Columns: []*schema.Column{HasSourceAtsColumns[7], HasSourceAtsColumns[5], HasSourceAtsColumns[2]},
				Annotation: &entsql.IndexAnnotation{
					Where: "package_version_id IS NOT NULL AND package_name_id IS NULL",
				},
			},
			{
				Name:    "hassourceat_source_id_package_name_id_justification",
				Unique:  true,
				Columns: []*schema.Column{HasSourceAtsColumns[7], HasSourceAtsColumns[6], HasSourceAtsColumns[2]},
				Annotation: &entsql.IndexAnnotation{
					Where: "package_name_id IS NOT NULL AND package_version_id IS NULL",
				},
			},
			{
				Name:    "hassourceat_known_since",
				Unique:  false,
				Columns: []*schema.Column{HasSourceAtsColumns[1]},
			},
		},
	}
	// HashEqualsColumns holds the columns for the "hash_equals" table.
	HashEqualsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "justification", Type: field.TypeString},
	}
	// HashEqualsTable holds the schema information for the "hash_equals" table.
	HashEqualsTable = &schema.Table{
		Name:       "hash_equals",
		Columns:    HashEqualsColumns,
		PrimaryKey: []*schema.Column{HashEqualsColumns[0]},
	}
	// IsVulnerabilitiesColumns holds the columns for the "is_vulnerabilities" table.
	IsVulnerabilitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "justification", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "osv_id", Type: field.TypeInt},
		{Name: "vulnerability_id", Type: field.TypeInt},
	}
	// IsVulnerabilitiesTable holds the schema information for the "is_vulnerabilities" table.
	IsVulnerabilitiesTable = &schema.Table{
		Name:       "is_vulnerabilities",
		Columns:    IsVulnerabilitiesColumns,
		PrimaryKey: []*schema.Column{IsVulnerabilitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "is_vulnerabilities_vulnerabilities_osv",
				Columns:    []*schema.Column{IsVulnerabilitiesColumns[4]},
				RefColumns: []*schema.Column{VulnerabilitiesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "is_vulnerabilities_vulnerabilities_vulnerability",
				Columns:    []*schema.Column{IsVulnerabilitiesColumns[5]},
				RefColumns: []*schema.Column{VulnerabilitiesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "isvulnerability_origin_justification_osv_id_vulnerability_id",
				Unique:  true,
				Columns: []*schema.Column{IsVulnerabilitiesColumns[2], IsVulnerabilitiesColumns[1], IsVulnerabilitiesColumns[4], IsVulnerabilitiesColumns[5]},
			},
		},
	}
	// OccurrencesColumns holds the columns for the "occurrences" table.
	OccurrencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "justification", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "artifact_id", Type: field.TypeInt},
		{Name: "package_id", Type: field.TypeInt, Nullable: true},
		{Name: "source_id", Type: field.TypeInt, Nullable: true},
	}
	// OccurrencesTable holds the schema information for the "occurrences" table.
	OccurrencesTable = &schema.Table{
		Name:       "occurrences",
		Columns:    OccurrencesColumns,
		PrimaryKey: []*schema.Column{OccurrencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "occurrences_artifacts_artifact",
				Columns:    []*schema.Column{OccurrencesColumns[4]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "occurrences_package_versions_package",
				Columns:    []*schema.Column{OccurrencesColumns[5]},
				RefColumns: []*schema.Column{PackageVersionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "occurrences_source_names_source",
				Columns:    []*schema.Column{OccurrencesColumns[6]},
				RefColumns: []*schema.Column{SourceNamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "occurrence_unique_package",
				Unique:  true,
				Columns: []*schema.Column{OccurrencesColumns[1], OccurrencesColumns[2], OccurrencesColumns[3], OccurrencesColumns[4], OccurrencesColumns[5]},
				Annotation: &entsql.IndexAnnotation{
					Where: "package_id IS NOT NULL AND source_id IS NULL",
				},
			},
			{
				Name:    "occurrence_unique_source",
				Unique:  true,
				Columns: []*schema.Column{OccurrencesColumns[1], OccurrencesColumns[2], OccurrencesColumns[3], OccurrencesColumns[4], OccurrencesColumns[6]},
				Annotation: &entsql.IndexAnnotation{
					Where: "package_id IS NULL AND source_id IS NOT NULL",
				},
			},
		},
	}
	// PackageNamesColumns holds the columns for the "package_names" table.
	PackageNamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "namespace_id", Type: field.TypeInt},
	}
	// PackageNamesTable holds the schema information for the "package_names" table.
	PackageNamesTable = &schema.Table{
		Name:       "package_names",
		Columns:    PackageNamesColumns,
		PrimaryKey: []*schema.Column{PackageNamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "package_names_package_namespaces_names",
				Columns:    []*schema.Column{PackageNamesColumns[2]},
				RefColumns: []*schema.Column{PackageNamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "packagename_name_namespace_id",
				Unique:  true,
				Columns: []*schema.Column{PackageNamesColumns[1], PackageNamesColumns[2]},
			},
		},
	}
	// PackageNamespacesColumns holds the columns for the "package_namespaces" table.
	PackageNamespacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "namespace", Type: field.TypeString},
		{Name: "package_id", Type: field.TypeInt},
	}
	// PackageNamespacesTable holds the schema information for the "package_namespaces" table.
	PackageNamespacesTable = &schema.Table{
		Name:       "package_namespaces",
		Columns:    PackageNamespacesColumns,
		PrimaryKey: []*schema.Column{PackageNamespacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "package_namespaces_package_types_namespaces",
				Columns:    []*schema.Column{PackageNamespacesColumns[2]},
				RefColumns: []*schema.Column{PackageTypesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "packagenamespace_namespace_package_id",
				Unique:  true,
				Columns: []*schema.Column{PackageNamespacesColumns[1], PackageNamespacesColumns[2]},
			},
		},
	}
	// PackageTypesColumns holds the columns for the "package_types" table.
	PackageTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Unique: true},
	}
	// PackageTypesTable holds the schema information for the "package_types" table.
	PackageTypesTable = &schema.Table{
		Name:       "package_types",
		Columns:    PackageTypesColumns,
		PrimaryKey: []*schema.Column{PackageTypesColumns[0]},
	}
	// PackageVersionsColumns holds the columns for the "package_versions" table.
	PackageVersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "version", Type: field.TypeString, Default: ""},
		{Name: "subpath", Type: field.TypeString, Default: ""},
		{Name: "qualifiers", Type: field.TypeJSON, Nullable: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "name_id", Type: field.TypeInt},
	}
	// PackageVersionsTable holds the schema information for the "package_versions" table.
	PackageVersionsTable = &schema.Table{
		Name:       "package_versions",
		Columns:    PackageVersionsColumns,
		PrimaryKey: []*schema.Column{PackageVersionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "package_versions_package_names_versions",
				Columns:    []*schema.Column{PackageVersionsColumns[5]},
				RefColumns: []*schema.Column{PackageNamesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "packageversion_hash_name_id",
				Unique:  true,
				Columns: []*schema.Column{PackageVersionsColumns[4], PackageVersionsColumns[5]},
			},
			{
				Name:    "packageversion_qualifiers",
				Unique:  false,
				Columns: []*schema.Column{PackageVersionsColumns[3]},
				Annotation: &entsql.IndexAnnotation{
					Types: map[string]string{
						"postgres": "GIN",
					},
				},
			},
		},
	}
	// PkgEqualsColumns holds the columns for the "pkg_equals" table.
	PkgEqualsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "justification", Type: field.TypeString},
		{Name: "packages_hash", Type: field.TypeString},
	}
	// PkgEqualsTable holds the schema information for the "pkg_equals" table.
	PkgEqualsTable = &schema.Table{
		Name:       "pkg_equals",
		Columns:    PkgEqualsColumns,
		PrimaryKey: []*schema.Column{PkgEqualsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "pkgequal_packages_hash_origin_justification_collector",
				Unique:  true,
				Columns: []*schema.Column{PkgEqualsColumns[4], PkgEqualsColumns[1], PkgEqualsColumns[3], PkgEqualsColumns[2]},
			},
		},
	}
	// SlsaAttestationsColumns holds the columns for the "slsa_attestations" table.
	SlsaAttestationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "build_type", Type: field.TypeString},
		{Name: "slsa_predicate", Type: field.TypeJSON, Nullable: true},
		{Name: "slsa_version", Type: field.TypeString},
		{Name: "started_on", Type: field.TypeTime, Nullable: true},
		{Name: "finished_on", Type: field.TypeTime, Nullable: true},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "built_from_hash", Type: field.TypeString},
		{Name: "built_by_id", Type: field.TypeInt},
		{Name: "subject_id", Type: field.TypeInt},
	}
	// SlsaAttestationsTable holds the schema information for the "slsa_attestations" table.
	SlsaAttestationsTable = &schema.Table{
		Name:       "slsa_attestations",
		Columns:    SlsaAttestationsColumns,
		PrimaryKey: []*schema.Column{SlsaAttestationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "slsa_attestations_builders_built_by",
				Columns:    []*schema.Column{SlsaAttestationsColumns[9]},
				RefColumns: []*schema.Column{BuildersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "slsa_attestations_artifacts_subject",
				Columns:    []*schema.Column{SlsaAttestationsColumns[10]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "slsaattestation_subject_id_origin_collector_build_type_slsa_version_built_by_id_built_from_hash",
				Unique:  true,
				Columns: []*schema.Column{SlsaAttestationsColumns[10], SlsaAttestationsColumns[6], SlsaAttestationsColumns[7], SlsaAttestationsColumns[1], SlsaAttestationsColumns[3], SlsaAttestationsColumns[9], SlsaAttestationsColumns[8]},
			},
			{
				Name:    "slsaattestation_started_on",
				Unique:  false,
				Columns: []*schema.Column{SlsaAttestationsColumns[4]},
			},
			{
				Name:    "slsaattestation_finished_on",
				Unique:  false,
				Columns: []*schema.Column{SlsaAttestationsColumns[5]},
			},
		},
	}
	// ScorecardsColumns holds the columns for the "scorecards" table.
	ScorecardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "checks", Type: field.TypeJSON},
		{Name: "aggregate_score", Type: field.TypeFloat64, Default: 0},
		{Name: "time_scanned", Type: field.TypeTime},
		{Name: "scorecard_version", Type: field.TypeString},
		{Name: "scorecard_commit", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
	}
	// ScorecardsTable holds the schema information for the "scorecards" table.
	ScorecardsTable = &schema.Table{
		Name:       "scorecards",
		Columns:    ScorecardsColumns,
		PrimaryKey: []*schema.Column{ScorecardsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "scorecard_origin_collector_scorecard_version_scorecard_commit_aggregate_score",
				Unique:  true,
				Columns: []*schema.Column{ScorecardsColumns[6], ScorecardsColumns[7], ScorecardsColumns[4], ScorecardsColumns[5], ScorecardsColumns[2]},
			},
		},
	}
	// SourceNamesColumns holds the columns for the "source_names" table.
	SourceNamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "commit", Type: field.TypeString, Nullable: true},
		{Name: "tag", Type: field.TypeString, Nullable: true},
		{Name: "namespace_id", Type: field.TypeInt},
	}
	// SourceNamesTable holds the schema information for the "source_names" table.
	SourceNamesTable = &schema.Table{
		Name:       "source_names",
		Columns:    SourceNamesColumns,
		PrimaryKey: []*schema.Column{SourceNamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "source_names_source_namespaces_namespace",
				Columns:    []*schema.Column{SourceNamesColumns[4]},
				RefColumns: []*schema.Column{SourceNamespacesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "sourcename_namespace_id_name_commit_tag",
				Unique:  true,
				Columns: []*schema.Column{SourceNamesColumns[4], SourceNamesColumns[1], SourceNamesColumns[2], SourceNamesColumns[3]},
			},
		},
	}
	// SourceNamespacesColumns holds the columns for the "source_namespaces" table.
	SourceNamespacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "namespace", Type: field.TypeString},
		{Name: "source_id", Type: field.TypeInt},
	}
	// SourceNamespacesTable holds the schema information for the "source_namespaces" table.
	SourceNamespacesTable = &schema.Table{
		Name:       "source_namespaces",
		Columns:    SourceNamespacesColumns,
		PrimaryKey: []*schema.Column{SourceNamespacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "source_namespaces_source_types_source_type",
				Columns:    []*schema.Column{SourceNamespacesColumns[2]},
				RefColumns: []*schema.Column{SourceTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "sourcenamespace_namespace_source_id",
				Unique:  true,
				Columns: []*schema.Column{SourceNamespacesColumns[1], SourceNamespacesColumns[2]},
			},
		},
	}
	// SourceTypesColumns holds the columns for the "source_types" table.
	SourceTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Unique: true},
	}
	// SourceTypesTable holds the schema information for the "source_types" table.
	SourceTypesTable = &schema.Table{
		Name:       "source_types",
		Columns:    SourceTypesColumns,
		PrimaryKey: []*schema.Column{SourceTypesColumns[0]},
	}
	// VulnerabilitiesColumns holds the columns for the "vulnerabilities" table.
	VulnerabilitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ghsa_id", Type: field.TypeString, Nullable: true},
		{Name: "cve_id", Type: field.TypeString, Nullable: true},
		{Name: "cve_year", Type: field.TypeInt, Nullable: true},
		{Name: "osv_id", Type: field.TypeString, Nullable: true},
	}
	// VulnerabilitiesTable holds the schema information for the "vulnerabilities" table.
	VulnerabilitiesTable = &schema.Table{
		Name:       "vulnerabilities",
		Columns:    VulnerabilitiesColumns,
		PrimaryKey: []*schema.Column{VulnerabilitiesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "vulnerability_ghsa_id",
				Unique:  true,
				Columns: []*schema.Column{VulnerabilitiesColumns[1]},
				Annotation: &entsql.IndexAnnotation{
					Where: "osv_id IS NULL AND cve_id IS NULL AND ghsa_id IS NOT NULL",
				},
			},
			{
				Name:    "vulnerability_cve_id",
				Unique:  true,
				Columns: []*schema.Column{VulnerabilitiesColumns[2]},
				Annotation: &entsql.IndexAnnotation{
					Where: "osv_id IS NULL AND cve_id IS NOT NULL AND ghsa_id IS NULL",
				},
			},
			{
				Name:    "vulnerability_osv_id",
				Unique:  true,
				Columns: []*schema.Column{VulnerabilitiesColumns[4]},
				Annotation: &entsql.IndexAnnotation{
					Where: "osv_id IS NOT NULL AND cve_id IS NULL AND ghsa_id IS NULL",
				},
			},
		},
	}
	// HashEqualArtifactsColumns holds the columns for the "hash_equal_artifacts" table.
	HashEqualArtifactsColumns = []*schema.Column{
		{Name: "hash_equal_id", Type: field.TypeInt},
		{Name: "artifact_id", Type: field.TypeInt},
	}
	// HashEqualArtifactsTable holds the schema information for the "hash_equal_artifacts" table.
	HashEqualArtifactsTable = &schema.Table{
		Name:       "hash_equal_artifacts",
		Columns:    HashEqualArtifactsColumns,
		PrimaryKey: []*schema.Column{HashEqualArtifactsColumns[0], HashEqualArtifactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hash_equal_artifacts_hash_equal_id",
				Columns:    []*schema.Column{HashEqualArtifactsColumns[0]},
				RefColumns: []*schema.Column{HashEqualsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "hash_equal_artifacts_artifact_id",
				Columns:    []*schema.Column{HashEqualArtifactsColumns[1]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PkgEqualPackagesColumns holds the columns for the "pkg_equal_packages" table.
	PkgEqualPackagesColumns = []*schema.Column{
		{Name: "pkg_equal_id", Type: field.TypeInt},
		{Name: "package_version_id", Type: field.TypeInt},
	}
	// PkgEqualPackagesTable holds the schema information for the "pkg_equal_packages" table.
	PkgEqualPackagesTable = &schema.Table{
		Name:       "pkg_equal_packages",
		Columns:    PkgEqualPackagesColumns,
		PrimaryKey: []*schema.Column{PkgEqualPackagesColumns[0], PkgEqualPackagesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pkg_equal_packages_pkg_equal_id",
				Columns:    []*schema.Column{PkgEqualPackagesColumns[0]},
				RefColumns: []*schema.Column{PkgEqualsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "pkg_equal_packages_package_version_id",
				Columns:    []*schema.Column{PkgEqualPackagesColumns[1]},
				RefColumns: []*schema.Column{PackageVersionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SlsaAttestationBuiltFromColumns holds the columns for the "slsa_attestation_built_from" table.
	SlsaAttestationBuiltFromColumns = []*schema.Column{
		{Name: "slsa_attestation_id", Type: field.TypeInt},
		{Name: "artifact_id", Type: field.TypeInt},
	}
	// SlsaAttestationBuiltFromTable holds the schema information for the "slsa_attestation_built_from" table.
	SlsaAttestationBuiltFromTable = &schema.Table{
		Name:       "slsa_attestation_built_from",
		Columns:    SlsaAttestationBuiltFromColumns,
		PrimaryKey: []*schema.Column{SlsaAttestationBuiltFromColumns[0], SlsaAttestationBuiltFromColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "slsa_attestation_built_from_slsa_attestation_id",
				Columns:    []*schema.Column{SlsaAttestationBuiltFromColumns[0]},
				RefColumns: []*schema.Column{SlsaAttestationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "slsa_attestation_built_from_artifact_id",
				Columns:    []*schema.Column{SlsaAttestationBuiltFromColumns[1]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArtifactsTable,
		BillOfMaterialsTable,
		BuildersTable,
		CertificationsTable,
		CertifyScorecardsTable,
		CertifyVulnsTable,
		DependenciesTable,
		HasSourceAtsTable,
		HashEqualsTable,
		IsVulnerabilitiesTable,
		OccurrencesTable,
		PackageNamesTable,
		PackageNamespacesTable,
		PackageTypesTable,
		PackageVersionsTable,
		PkgEqualsTable,
		SlsaAttestationsTable,
		ScorecardsTable,
		SourceNamesTable,
		SourceNamespacesTable,
		SourceTypesTable,
		VulnerabilitiesTable,
		HashEqualArtifactsTable,
		PkgEqualPackagesTable,
		SlsaAttestationBuiltFromTable,
	}
)

func init() {
	BillOfMaterialsTable.ForeignKeys[0].RefTable = PackageVersionsTable
	BillOfMaterialsTable.ForeignKeys[1].RefTable = ArtifactsTable
	CertificationsTable.ForeignKeys[0].RefTable = SourceNamesTable
	CertificationsTable.ForeignKeys[1].RefTable = PackageVersionsTable
	CertificationsTable.ForeignKeys[2].RefTable = PackageNamesTable
	CertificationsTable.ForeignKeys[3].RefTable = ArtifactsTable
	CertifyScorecardsTable.ForeignKeys[0].RefTable = SourceNamesTable
	CertifyScorecardsTable.ForeignKeys[1].RefTable = ScorecardsTable
	CertifyVulnsTable.ForeignKeys[0].RefTable = VulnerabilitiesTable
	CertifyVulnsTable.ForeignKeys[1].RefTable = PackageVersionsTable
	DependenciesTable.ForeignKeys[0].RefTable = PackageVersionsTable
	DependenciesTable.ForeignKeys[1].RefTable = PackageNamesTable
	HasSourceAtsTable.ForeignKeys[0].RefTable = PackageVersionsTable
	HasSourceAtsTable.ForeignKeys[1].RefTable = PackageNamesTable
	HasSourceAtsTable.ForeignKeys[2].RefTable = SourceNamesTable
	IsVulnerabilitiesTable.ForeignKeys[0].RefTable = VulnerabilitiesTable
	IsVulnerabilitiesTable.ForeignKeys[1].RefTable = VulnerabilitiesTable
	OccurrencesTable.ForeignKeys[0].RefTable = ArtifactsTable
	OccurrencesTable.ForeignKeys[1].RefTable = PackageVersionsTable
	OccurrencesTable.ForeignKeys[2].RefTable = SourceNamesTable
	PackageNamesTable.ForeignKeys[0].RefTable = PackageNamespacesTable
	PackageNamespacesTable.ForeignKeys[0].RefTable = PackageTypesTable
	PackageVersionsTable.ForeignKeys[0].RefTable = PackageNamesTable
	SlsaAttestationsTable.ForeignKeys[0].RefTable = BuildersTable
	SlsaAttestationsTable.ForeignKeys[1].RefTable = ArtifactsTable
	SlsaAttestationsTable.Annotation = &entsql.Annotation{
		Table: "slsa_attestations",
	}
	SourceNamesTable.ForeignKeys[0].RefTable = SourceNamespacesTable
	SourceNamespacesTable.ForeignKeys[0].RefTable = SourceTypesTable
	HashEqualArtifactsTable.ForeignKeys[0].RefTable = HashEqualsTable
	HashEqualArtifactsTable.ForeignKeys[1].RefTable = ArtifactsTable
	PkgEqualPackagesTable.ForeignKeys[0].RefTable = PkgEqualsTable
	PkgEqualPackagesTable.ForeignKeys[1].RefTable = PackageVersionsTable
	SlsaAttestationBuiltFromTable.ForeignKeys[0].RefTable = SlsaAttestationsTable
	SlsaAttestationBuiltFromTable.ForeignKeys[1].RefTable = ArtifactsTable
}
