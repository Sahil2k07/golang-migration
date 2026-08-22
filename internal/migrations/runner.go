package migrations

func RunMigrations() error {
	_, err := getMigrationHistory()
	if err != nil {
		return err
	}

	return nil
}
