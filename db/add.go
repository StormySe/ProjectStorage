package db

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"

  "storage/models"
  "storage/config"
)

func Add(purchase models.Pucrhase, client models.Client, vendor models.Vendor, storer models.Storer) {
  db, _ := sql.Open("sqlite3", config.AppConfig.DBName)
  defer db.Close()

  var (
    clientId int;
    vendorId int;
    storerId int;
  )
  // for each of database
  clientRows := db.QueryRow("SELECT id FROM Clients WHERE name = ?;", client.Name)
  if err := clientRows.Scan(&clientId); err == sql.ErrNoRows{
      add, _ := db.Prepare("INSERT INTO Clients (name) VALUES (?)")
      add.Exec(client.Name)
      repeatQuery := db.QueryRow("SELECT id FROM Clients WHERE name = ?;", client.Name)
      if err = repeatQuery.Scan(&clientId); err != nil {
        panic(err)
      }
  }

  vendorRows := db.QueryRow("SELECT id FROM Vendors WHERE name = ?;", vendor.Name)
  if err := vendorRows.Scan(&vendorId); err == sql.ErrNoRows{
      add, _ := db.Prepare("INSERT INTO Vendors (name) VALUES (?)")
      add.Exec(vendor.Name)
      repeatQuery := db.QueryRow("SELECT id FROM Vendors WHERE name = ?;", vendor.Name)
      if err = repeatQuery.Scan(&vendorId); err != nil {
        panic(err)
      }
  }

  storerRows := db.QueryRow("SELECT id FROM Storers WHERE name = ?;", storer.Name)
  if err := clientRows.Scan(&storerRows); err == sql.ErrNoRows{
      add, _ := db.Prepare("INSERT INTO Storers (name) VALUES (?)")
      add.Exec(storer.Name)
      repeatQuery := db.QueryRow("SELECT id FROM Storers WHERE name = ?;", storer.Name)
      if err = repeatQuery.Scan(&storerId); err != nil {
        panic(err)
      }
  }

  purchase.ClientId = uint(clientId)
  purchase.VendorId = uint(vendorId)
  purchase.StorerId = uint(storerId)

  purchaseAdd, _ := db.Prepare("INSERT INTO Purchases (name, client_id, vendor_id, storer_id) VALUES (?, ?, ?, ?)")
  purchaseAdd.Exec(purchase.Name, purchase.ClientId, purchase.VendorId, purchase.StorerId)
}
