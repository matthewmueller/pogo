package postgres

import (
	"database/sql"

	"github.com/matthewmueller/pogo/internal/pogo"
)

// Here we introspect the System Catalogs
// https://www.postgresql.org/docs/11/catalogs.html
//
// Most system catalogs are copied from the template database during
// database creation and are thereafter database-specific. A few catalogs
// are physically shared across all databases in a cluster; these are noted
// in the descriptions of the individual catalogs.
//
// System Catalogs:
//
// Catalog Name              Purpose
// –––––––––––––––––––––––––––––––––
// pg_aggregate              aggregate functions
// pg_am                     index access methods
// pg_amop                   access method operators
// pg_amproc                 access method support functions
// pg_attrdef                column default values
// pg_attribute              table columns (“attributes”)
// pg_authid                 authorization identifiers (roles)
// pg_auth_members           authorization identifier membership relationships
// pg_cast                   casts (data type conversions)
// pg_class                  tables, indexes, sequences, views (“relations”)
// pg_collation              collations (locale information)
// pg_constraint             check constraints, unique constraints, primary key constraints, foreign key constraints
// pg_conversion             encoding conversion information
// pg_database               databases within this database cluster
// pg_db_role_setting        per-role and per-database settings
// pg_default_acl            default privileges for object types
// pg_depend                 dependencies between database objects
// pg_description            descriptions or comments on database objects
// pg_enum                   enum label and value definitions
// pg_event_trigger          event triggers
// pg_extension              installed extensions
// pg_foreign_data_wrapper   foreign-data wrapper definitions
// pg_foreign_server         foreign server definitions
// pg_foreign_table          additional foreign table information
// pg_index                  additional index information
// pg_inherits               table inheritance hierarchy
// pg_init_privs             object initial privileges
// pg_language               languages for writing functions
// pg_largeobject            data pages for large objects
// pg_largeobject_metadata   metadata for large objects
// pg_namespace              schemas
// pg_opclass                access method operator classes
// pg_operator               operators
// pg_opfamily               access method operator families
// pg_partitioned_table      information about partition key of tables
// pg_pltemplate             template data for procedural languages
// pg_policy                 row-security policies
// pg_proc                   functions and procedures
// pg_publication            publications for logical replication
// pg_publication_rel        relation to publication mapping
// pg_range                  information about range types
// pg_replication_origin     registered replication origins
// pg_rewrite                query rewrite rules
// pg_seclabel               security labels on database objects
// pg_sequence               information about sequences
// pg_shdepend               dependencies on shared objects
// pg_shdescription          comments on shared objects
// pg_shseclabel             security labels on shared database objects
// pg_statistic              planner statistics
// pg_statistic_ext          extended planner statistics
// pg_subscription           logical replication subscriptions
// pg_subscription_rel       relation state for subscriptions
// pg_tablespace             tablespaces within this database cluster
// pg_transform              transforms (data type to procedural language conversions)
// pg_trigger                triggers
// pg_ts_config              text search configurations
// pg_ts_config_map          text search configurations' token mappings
// pg_ts_dict                text search dictionaries
// pg_ts_parser              text search parsers
// pg_ts_template            text search templates
// pg_type                   data types
// pg_user_mapping           mappings of users to foreign servers

// Introspect the database
func Introspect(db *sql.DB) (*pogo.Database, error) {
	schemas, err := introspectSchemas(db)
	if err != nil {
		return nil, err
	}
	_ = schemas
	return nil, nil
}

// Schema struct
//
// The catalog pg_namespace stores namespaces. A namespace is the structure
// underlying SQL schemas: each namespace can have a separate collection of
// relations, types, etc. without name conflicts.
//
// Name       Type        References      Description
// oid        oid                         Row identifier (hidden attribute; must be explicitly selected)
// nspname    name                        Name of the namespace
// nspowner   oid         pg_authid.oid   Owner of the namespace
// nspacl     aclitem[]                   Access privileges; see GRANT and REVOKE for details
type Schema struct {
	OID   int
	Name  string
	Owner int
	ACL   []string
}

func introspectSchemas(db *sql.DB) (schemas []*Schema, err error) {

	return schemas, nil
}

// Table struct
type Table struct {
	Type     byte   // type
	ManualPk bool   // manual_pk
	Name     string // table_name
}

// Column struct
type Column struct {
	FieldOrdinal int
	Name         string
	DataType     string
	NotNull      bool
	Comment      *string
	DefaultValue *string
	IsPrimaryKey bool
}

// Index struct
type Index struct {
	Name      string
	IsUnique  bool
	IsPrimary bool
	SeqNo     int
	Origin    string
	IsPartial bool
}

// IndexColumn struct
type IndexColumn struct {
	SeqNo    int
	Cid      int
	Name     string
	DataType string
	NotNull  bool
}

// ForeignKey struct
type ForeignKey struct {
	Name          string
	FullName      string
	DataType      string
	RefIndexName  string
	RefTableName  string
	RefColumnName string
	KeyID         int
	SeqNo         int
	OnUpdate      string
	OnDelete      string
	Match         string
}
