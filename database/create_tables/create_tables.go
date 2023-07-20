package create_tables

func CreateAllTablesIfNotExists() (err error) {
	err = CreateUsersTable()
	if err != nil {
		return err
	}
	err = CreateWebsitesTable()
	if err != nil {
		return err
	}
	err = CreateVisitorsTable()
	if err != nil {
		return err
	}
	err = CreateSessionsTable()
	if err != nil {
		return err
	}

	return nil
}
