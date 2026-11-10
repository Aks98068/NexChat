package database


func Migratedb(db *gorm.DB) error{

  if db == nil{
	return fmt.Errorf("database connection is nil")
  }

  err := db.AutoMigrate(
	
	&models.User{}, 
	&models.Message{},
   )
  if err != nil {
		return fmt.Errorf(
			"failed to migrate the database schema: %w",
			err,
		)
	}


	return nil


}