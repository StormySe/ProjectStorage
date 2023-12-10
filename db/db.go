package db

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"

  "storage/models"
  "storage/config"
)

var Purchases []models.PucrhaseOut

func Init() {
  db, _ := sql.Open("sqlite3", config.AppConfig.DBName)
  defer db.Close()

  if config.AppConfig.DemoMode {
    db.Exec(`DROP TABLE IF EXISTS Clients`)
    db.Exec(`DROP TABLE IF EXISTS Purchases`)
    db.Exec(`DROP TABLE IF EXISTS Vendors`)
    db.Exec(`DROP TABLE IF EXISTS Storers`)
  }

  clientCreate, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS Clients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
  );`)
  vendorCreate, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS Vendors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
  );`)
  storerCreate, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS Storers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
  );`)
  purchasesCreate, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS Purchases (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    client_id INTEGER,
    storer_id INTEGER,
    vendor_id INTEGER,
    FOREIGN KEY (client_id) REFERENCES "Clients"("id"),
    FOREIGN KEY (storer_id) REFERENCES "Storers"("id"),
    FOREIGN KEY (vendor_id) REFERENCES "Vendors"("id")
  );`)
  clientCreate.Exec()
  vendorCreate.Exec()
  storerCreate.Exec()
  purchasesCreate.Exec()
  Get()
}

