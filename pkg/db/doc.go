// Package db provides database connection management via GORM.
//
// It supports MySQL, PostgreSQL, and SQLServer with multi-instance
// connections (lazy-initialized on first use). Each database is configured
// through the config package with support for connection pooling, table
// prefixes, and singular table names.
package db
