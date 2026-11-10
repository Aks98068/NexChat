package database




func Connect(databaseURL string) (*gorm.DB, error){
	db, err := gorm.Open(postgres.Open(databaseURL), gorm.Config{})
	if err != nil{
		return nil,fmt.Errorf("failed to connect to database: %v", err)
	}
	sqlDB , err := db.DB()
	if err != nil{
		return nil, fmt.Errorf("failed to get underlying database connection: %v", err)
	}

	// connections pool settings 
	sqlDB.SetMaxOpenConns(25) // maximum number of open connections to the database
	sqlDB.SetMaxIdleConns(25) // maximum number of idle connections in the pool
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // maximum amount of time a connection may be reused
	sqlDB.SetConnMaxIdleTime(5 * time.Minute) // maximum amount of time a connection may be idle

	// verify the connection of the dataabse
	if err := sqlDB.Ping(); err != nil{
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil

}