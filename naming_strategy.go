package snowflake

import (
	"strings"

	"gorm.io/gorm/schema"
)

// NamingStrategy for snowflake (always uppercase)
type NamingStrategy struct {
	defaultNS schema.Namer
}

// NewNamingStrategy create new instance of snowflake naming strat
func NewNamingStrategy() schema.Namer {
	return &NamingStrategy{
		defaultNS: schema.NamingStrategy{},
	}
}

// ColumnName snowflake edition
func (ns NamingStrategy) ColumnName(table, column string) string {
	return strings.ToUpper(ns.defaultNS.ColumnName(table, column))
}

// TableName snowflake edition
func (ns NamingStrategy) TableName(table string) string {
	return ns.defaultNS.TableName(table)
}

// SchemaName snowflake edition
func (ns NamingStrategy) SchemaName(table string) string {
	return ns.defaultNS.SchemaName(table)
}

// JoinTableName snowflake edition
func (ns NamingStrategy) JoinTableName(joinTable string) string {
	return ns.defaultNS.JoinTableName(joinTable)
}

// RelationshipFKName snowflake edition
func (ns NamingStrategy) RelationshipFKName(rel schema.Relationship) string {
	return ns.defaultNS.RelationshipFKName(rel)
}

// CheckerName snowflake edition
func (ns NamingStrategy) CheckerName(table, column string) string {
	return ns.defaultNS.CheckerName(table, column)
}

// IndexName snowflake edition
func (ns NamingStrategy) IndexName(table, column string) string {
	return ns.defaultNS.IndexName(table, column)
}

// UniqueName generate unique constraint name
func (ns NamingStrategy) UniqueName(table, column string) string {
	return ns.defaultNS.UniqueName(table, column)
}
