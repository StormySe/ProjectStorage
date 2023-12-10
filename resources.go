package main

import (
  	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
    "storage/db"

)

// resources
var addIcon, _ = fyne.LoadResourceFromPath("./icons/add.png")
var settingsIcon, _ = fyne.LoadResourceFromPath("./icons/settings.png")
var searchIcon, _ = fyne.LoadResourceFromPath("./icons/search.png")
var backIcon, _ = fyne.LoadResourceFromPath("./icons/back.png")
var editIcon, _ = fyne.LoadResourceFromPath("./icons/edit.png")
var deleteIcon, _ = fyne.LoadResourceFromPath("./icons/delete.png")

// entries
var purchaseEntry = widget.NewEntry()
var vendorEntry = widget.NewEntry()
var clientEntry = widget.NewEntry()
var storerEntry = widget.NewEntry()

var purchasesTable = widget.NewList(
  func() int {
    return len(db.Purchases)
  },
  func() fyne.CanvasObject {
    return widget.NewLabel("Default")
  },
  func(lii widget.ListItemID, co fyne.CanvasObject) {
    co.(*widget.Label).SetText(db.Purchases[lii].Name)
  },
)

var form = widget.NewForm(
  widget.NewFormItem("Purchase: ", purchaseEntry),
  widget.NewFormItem("Vendor: ", vendorEntry),
  widget.NewFormItem("Client: ", clientEntry),
  widget.NewFormItem("Storer: ", storerEntry),
)
