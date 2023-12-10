package db

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"

  "storage/models"
  "storage/config"
)

func Get() {
  db, _ := sql.Open("sqlite3", config.AppConfig.DBName)
  defer db.Close()
  Purchases = []models.PucrhaseOut{}

  rows, _ := db.Query(`SELECT Purchases.*, Storers.name AS StorersName, Vendors.name AS VendorsName, Clients.name AS ClientsName
FROM Purchases
LEFT JOIN Storers ON Purchases.storer_id = Storers.id
LEFT JOIN Vendors ON Purchases.vendor_id = Vendors.id
LEFT JOIN Clients ON Purchases.client_id = Clients.id;
`)
  for rows.Next() {
    p := models.PucrhaseOut{}
    rows.Scan(&p.Id, &p.Name, &p.StorerId, &p.VendorId, &p.ClientId, &p.StorersName, &p.VendorsName, &p.ClientsName)
    Purchases = append(Purchases, p)
  }
}
