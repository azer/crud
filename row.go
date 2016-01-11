package crud

type RowValue struct {
	SQLColumn string
	Value     interface{}
}

type Row struct {
	SQLTableName string
	Values       []*RowValue
}

func (row *Row) SQLValues() map[string]interface{} {
	result := map[string]interface{}{}

	for _, v := range row.Values {
		result[v.SQLColumn] = v.Value
	}

	return result
}

func NewRow(st interface{}) (*Row, error) {
	values, err := GetRowValuesOf(st)
	if err != nil {
		return nil, err
	}

	return &Row{
		SQLTableName: SQLTableNameOf(st),
		Values:       values,
	}, nil
}

func GetRowValuesOf(st interface{}) ([]*RowValue, error) {
	values := []*RowValue{}

	iter := NewIteration(st)
	for iter.Next() {
		sqlOptions, err := iter.SQLOptions()

		if err != nil {
			return nil, err
		}

		if sqlOptions.Ignore {
			continue
		}

		value := iter.Value()

		if n, ok := value.(int); ok && sqlOptions.AutoIncrement > 0 && n == 0 {
			continue
		}

		values = append(values, &RowValue{
			SQLColumn: sqlOptions.Name,
			Value:     value,
		})
	}

	return values, nil
}
