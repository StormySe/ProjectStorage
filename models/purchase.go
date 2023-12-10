package models


type Pucrhase struct {
  Id uint
  Name string
  VendorId uint
  ClientId uint
  StorerId uint
}

type PucrhaseOut struct {
  Id uint   `json:"Id"`
  Name string   `json:"Name"`
  VendorId uint `json:"VendorId"`
  ClientId uint `json:"ClientId"`
  StorerId uint `json:"StorerId"`
  VendorsName string `json:"VendorsName"`
  ClientsName string `json:"ClientsName"`
  StorersName string  `json:"StorersName"`
}
