package main

import "github.com/eduardomarini16/api_os/os_manager/config"

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.RunMigrations(db)
}
