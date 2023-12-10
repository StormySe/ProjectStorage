package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"storage/models"
    "storage/config"
)

func Edit(e *models.PucrhaseOut, purchaseName, vendorName, clientName, storerName string) {
  db, _ := sql.Open("sqlite3", config.AppConfig.DBName)
  defer db.Close()
  updQueryFmt := `UPDATE %s SET name = ? WHERE name = ?;`
  dbNames := map [string]string {
    "Purchases": purchaseName,
    "Vendors": vendorName,
    "Clients": clientName,
    "Storers": storerName,
  }
  oldValues := map [string]string {
    "Purchases": e.Name,
    "Vendors": e.VendorsName,
    "Clients": e.ClientsName,
    "Storers": e.StorersName,
  }
  for dbName := range dbNames {
    query, _ := db.Prepare(fmt.Sprintf(updQueryFmt, dbName))
    query.Exec(dbNames[dbName], oldValues[dbName])
  }
  e.Name = purchaseName
  e.VendorsName = vendorName
  e.ClientsName = clientName
  e.StorersName = storerName
}
