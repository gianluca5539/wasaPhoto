package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeleteAllData() (error) {
	_,err := db.c.Query("DELETE FROM user")
	return err
}
