package store

type sqlCrudStore struct {
	SQLStore
	Table string
}

// NewSQLCrudStore create new schema
func NewSQLCrudStore(sqlStore SQLStore) CrudStore {
	cr := &sqlCrudStore{
		SQLStore: sqlStore,
		Table:    "Crud",
	}

	// cr.SQLStore.GetConn().AddTable()

	return cr
}

func (sql sqlCrudStore) Get(id string) Channel {
	return Do(func(result *Result) {

	})
}
