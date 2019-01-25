package postgres

import (
	"database/sql"
	"fmt"

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
// https://www.postgresql.org/docs/11/catalog-pg-namespace.html
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
	fmt.Println("introspecting...")
	return schemas, nil
}

// pg_class
// https://www.postgresql.org/docs/11/catalog-pg-class.html
//
// Name                  Type            References        Description
// oid                   oid                               Row identifier (hidden attribute; must be explicitly selected)
// relname               name                              Name of the table, index, view, etc.
// relnamespace          oid            pg_namespace.oid   The OID of the namespace that contains this relation
// reltype               oid            pg_type.oid        The OID of the data type that corresponds to this table's row type, if any (zero for indexes, which have no pg_type entry)
// reloftype             oid            pg_type.oid        For typed tables, the OID of the underlying composite type, zero for all other relations
// relowner              oid            pg_authid.oid      Owner of the relation
// relam                 oid            pg_am.oid          If this is an index, the access method used (B-tree, hash, etc.)
// relfilenode           oid                               Name of the on-disk file of this relation; zero means this is a “mapped” relation whose disk file name is determined by low-level state
// reltablespace         oid            pg_tablespace.oid  The tablespace in which this relation is stored. If zero, the database's default tablespace is implied. (Not meaningful if the relation has no on-disk file.)
// relpages              int4                              Size of the on-disk representation of this table in pages (of size BLCKSZ). This is only an estimate used by the planner. It is updated by VACUUM, ANALYZE, and a few DDL commands such as CREATE INDEX.
// reltuples             float4                            Number of live rows in the table. This is only an estimate used by the planner. It is updated by VACUUM, ANALYZE, and a few DDL commands such as CREATE INDEX.
// relallvisible         int4                              Number of pages that are marked all-visible in the table's visibility map. This is only an estimate used by the planner. It is updated by VACUUM, ANALYZE, and a few DDL commands such as CREATE INDEX.
// reltoastrelid         oid            pg_class.oid       OID of the TOAST table associated with this table, 0 if none. The TOAST table stores large attributes “out of line” in a secondary table.
// relhasindex           bool                              True if this is a table and it has (or recently had) any indexes
// relisshared           bool                              True if this table is shared across all databases in the cluster. Only certain system catalogs (such as pg_database) are shared.
// relpersistence        char                              p = permanent table, u = unlogged table, t = temporary table
// relkind               char                              r = ordinary table, i = index, S = sequence, t = TOAST table, v = view, m = materialized view, c = composite type, f = foreign table, p = partitioned table, I = partitioned index
// relnatts              int2                              Number of user columns in the relation (system columns not counted). There must be this many corresponding entries in pg_attribute. See also pg_attribute.attnum.
// relchecks             int2                              Number of CHECK constraints on the table; see pg_constraint catalog
// relhasoids            bool                              True if we generate an OID for each row of the relation
// relhasrules           bool                              True if table has (or once had) rules; see pg_rewrite catalog
// relhastriggers        bool                              True if table has (or once had) triggers; see pg_trigger catalog
// relhassubclass        bool                              True if table has (or once had) any inheritance children
// relrowsecurity        bool                              True if table has row level security enabled; see pg_policy catalog
// relforcerowsecurity   bool                              True if row level security (when enabled) will also apply to table owner; see pg_policy catalog
// relispopulated        bool                              True if relation is populated (this is true for all relations other than some materialized views)
// relreplident          char                              Columns used to form “replica identity” for rows: d = default (primary key, if any), n = nothing, f = all columns i = index with indisreplident set, or default
// relispartition        bool                              True if table or index is a partition
// relrewrite            oid            pg_class.oid       For new relations being written during a DDL operation that requires a table rewrite, this contains the OID of the original relation; otherwise 0. That state is only visible internally; this field should never contain anything other than 0 for a user-visible relation.
// relfrozenxid          xid                               All transaction IDs before this one have been replaced with a permanent (“frozen”) transaction ID in this table. This is used to track whether the table needs to be vacuumed in order to prevent transaction ID wraparound or to allow pg_xact to be shrunk. Zero (InvalidTransactionId) if the relation is not a table.
// relminmxid            xid                               All multixact IDs before this one have been replaced by a transaction ID in this table. This is used to track whether the table needs to be vacuumed in order to prevent multixact ID wraparound or to allow pg_multixact to be shrunk. Zero (InvalidMultiXactId) if the relation is not a table.
// relacl                aclitem[]                         Access privileges; see GRANT and REVOKE for details
// reloptions            text[]                            Access-method-specific options, as “keyword=value” strings
// relpartbound          pg_node_tree                      If table is a partition (see relispartition), internal representation of the partition bound

// Table struct
type Table struct {
	Type     byte   // type
	ManualPk bool   // manual_pk
	Name     string // table_name
}

// pg_attribute
// https://www.postgresql.org/docs/11/catalog-pg-attribute.html
//
// The catalog pg_attribute stores information about table columns.
// There will be exactly one pg_attribute row for every column in every
// table in the database. (There will also be attribute entries for indexes,
// and indeed all objects that have pg_class entries.)
//
// Name            Type         References        Description
// attrelid        oid         pg_class.oid       The table this column belongs to
// attname         name                           The column name
// atttypid        oid         pg_type.oid        The data type of this column
// attstattarget   int4                           attstattarget controls the level of detail of statistics accumulated for this column by ANALYZE. A zero value indicates that no statistics should be collected. A negative value says to use the system default statistics target. The exact meaning of positive values is data type-dependent. For scalar data types, attstattarget is both the target number of “most common values” to collect, and the target number of histogram bins to create.
// attlen          int2                           A copy of pg_type.typlen of this column's type
// attnum          int2                           The number of the column. Ordinary columns are numbered from 1 up. System columns, such as oid, have (arbitrary) negative numbers.
// attndims        int4                           Number of dimensions, if the column is an array type; otherwise 0. (Presently, the number of dimensions of an array is not enforced, so any nonzero value effectively means “it's an array”.)
// attcacheoff     int4                           Always -1 in storage, but when loaded into a row descriptor in memory this might be updated to cache the offset of the attribute within the row
// atttypmod       int4                           atttypmod records type-specific data supplied at table creation time (for example, the maximum length of a varchar column). It is passed to type-specific input functions and length coercion functions. The value will generally be -1 for types that do not need atttypmod.
// attbyval        bool                           A copy of pg_type.typbyval of this column's type
// attstorage      char                           Normally a copy of pg_type.typstorage of this column's type. For TOAST-able data types, this can be altered after column creation to control storage policy.
// attalign        char                           A copy of pg_type.typalign of this column's type
// attnotnull      bool                           This represents a not-null constraint.
// atthasdef       bool                           This column has a default value, in which case there will be a corresponding entry in the pg_attrdef catalog that actually defines the value.
// atthasmissing   bool                           This column has a value which is used where the column is entirely missing from the row, as happens when a column is added with a non-volatile DEFAULT value after the row is created. The actual value used is stored in the attmissingval column.
// attidentity     char                           If a zero byte (''), then not an identity column. Otherwise, a = generated always, d = generated by default.
// attisdropped    bool                           This column has been dropped and is no longer valid. A dropped column is still physically present in the table, but is ignored by the parser and so cannot be accessed via SQL.
// attislocal      bool                           This column is defined locally in the relation. Note that a column can be locally defined and inherited simultaneously.
// attinhcount     int4                           The number of direct ancestors this column has. A column with a nonzero number of ancestors cannot be dropped nor renamed.
// attcollation    oid         pg_collation.oid   The defined collation of the column, or zero if the column is not of a collatable data type.
// attacl          aclitem[]                      Column-level access privileges, if any have been granted specifically on this column
// attoptions      text[]                         Attribute-level options, as “keyword=value” strings
// attfdwoptions   text[]                         Attribute-level foreign data wrapper options, as “keyword=value” strings
// attmissingval   anyarray                       This column has a one element array containing the value used when the column is entirely missing from the row, as happens when the column is added with a non-volatile DEFAULT value after the row is created. The value is only used when atthasmissing is true. If there is no value the column is null.

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
