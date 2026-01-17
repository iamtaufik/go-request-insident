package main

import (
	"be-request-insident/internal/config"
	"be-request-insident/internal/db"

	"gorm.io/gen"
)

func init() {
	config.LoadEnv()
}

func main() {
	db := db.ConnectMysql()

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/models",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.GenerateAllTable()

	g.Execute()
}