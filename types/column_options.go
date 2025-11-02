package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ColumnOptions struct {
	Name               string
	Type               string
	Length             int
	DefaultValue       string
	AutoIncrement      int
	IsAutoIncrementing bool
	IsPrimaryKey       bool
	Serial             bool
	IsUnique           bool
	IsUnsigned         bool
	IsRequired         bool
	Ignore             bool
	TableName          string
	CurrentTimestamp   bool
}

func NewColumnOptions(input string) (*ColumnOptions, error) {
	options := &ColumnOptions{}
	input = strings.TrimSpace(input)

	if input == "" {
		return options, nil
	}

	if input == "-" {
		options.Ignore = true
		return options, nil
	}

	parts := strings.Split(input, " ")

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if options.ReadType(part) {
			continue
		}

		if options.ReadTypeAttr(part) {
			continue
		}

		if options.ReadName(part) {
			continue
		}

		if options.ReadTableName(part) {
			continue
		}

		if options.ReadDefaultValue(part) {
			continue
		}

		if options.ReadAutoIncrement(part) {
			continue
		}

		if part == "primary-key" || part == "primary_key" || part == "primarykey" {
			options.IsPrimaryKey = true
			continue
		}

		if part == "required" {
			options.IsRequired = true
			continue
		}

		if part == "serial" {
			options.Serial = true
			continue
		}

		if part == "unsigned" {
			options.IsUnsigned = true
			continue
		}

		if part == "unique" {
			options.IsUnique = true
			continue
		}

		if part == "now" {
			options.CurrentTimestamp = true
			continue
		}

		return nil, errors.New(fmt.Sprintf("Unrecognized SQL column option: %s", part))
	}

	return options, nil
}

func (options *ColumnOptions) ReadAttr(input string, names ...string) (string, bool) {
	parts := strings.Split(input, "=")

	for _, name := range names {
		if parts[0] == name && len(parts) > 1 {
			return parts[1], true
		} else if parts[0] == name {
			return "", true
		}
	}

	return "", false
}

func (options *ColumnOptions) ReadTypeAttr(input string) bool {
	value, ok := options.ReadAttr(input, "type")
	if !ok || len(value) == 0 {
		return false
	}

	return options.ReadType(value)
}

func (options *ColumnOptions) ReadType(input string) bool {
	parts := strings.FieldsFunc(input, func(c rune) bool {
		return c == '(' || c == ')'
	})

	typeName := parts[0]
	length, isType := GetValidSqlType(typeName)

	if !isType {
		return false
	}

	if len(parts) > 1 {
		if plen, err := strconv.Atoi(parts[1]); err == nil {
			length = plen
		}
	}

	options.Type = typeName
	options.Length = length

	return true
}

func (options *ColumnOptions) ReadDefaultValue(input string) bool {
	value, ok := options.ReadAttr(input, "default")
	if !ok {
		return false
	}

	options.DefaultValue = value
	return true
}

func (options *ColumnOptions) ReadName(input string) bool {
	value, ok := options.ReadAttr(input, "name")
	if !ok {
		return false
	}

	options.Name = value
	return true
}

func (options *ColumnOptions) ReadTableName(input string) bool {
	value, ok := options.ReadAttr(input, "table-name")
	if !ok {
		return false
	}

	options.TableName = value
	return true
}

func (options *ColumnOptions) ReadAutoIncrement(input string) bool {
	value, ok := options.ReadAttr(input, "auto-increment", "auto_increment", "autoincrement")
	if !ok {
		return false
	}

	options.IsAutoIncrementing = true

	if len(value) > 0 {
		if val, err := strconv.Atoi(value); err == nil {
			options.AutoIncrement = val
		} else {
			options.AutoIncrement = 1
		}
	} else {
		options.AutoIncrement = 1
	}

	return true
}
