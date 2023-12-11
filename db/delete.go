package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"storage/config"
	"storage/models"
)

func Delete(purchase models.PucrhaseOut) {
  db, _ := sql.Open("sqlite3", config.AppConfig.DBName)
  defer db.Close()

  var (
    vendors int
    storers int
    clients int
  )

  delPurchase, _ := db.Prepare(`DELETE FROM Purchases WHERE id = ?;`)
  findVendor, _ := db.Prepare(`SELECT count(vendor_id) FROM Purchases WHERE vendor_id = ?;`)
  findClient, _ := db.Prepare(`SELECT count(client_id) FROM Purchases WHERE client_id = ?;`)
  findStorer, _ := db.Prepare(`SELECT count(storer_id) FROM Purchases WHERE storer_id = ?;`)

  
  count := findVendor.QueryRow(purchase.VendorId)
  count.Scan(&vendors)

  count = findStorer.QueryRow(purchase.StorerId)
  count.Scan(&storers)

  count = findClient.QueryRow(purchase.ClientId)
  count.Scan(&clients)

  if vendors <= 1 {
    db.Exec(`DELETE FROM Vendors WHERE id = ?;`, purchase.VendorId)
  }
  if storers <= 1 {
    db.Exec(`DELETE FROM Storers WHERE id = ?;`, purchase.StorerId)
  }
  if clients <= 1 {
    db.Exec(`DELETE FROM Clients WHERE id = ?;`, purchase.ClientId)
  }
  delPurchase.Exec(purchase.Id)
}
