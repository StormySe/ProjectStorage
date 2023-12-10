package main

import (
	"fmt"
	"image/color"
	"storage/db"
	"storage/models"
    "storage/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
)


func main() {
  config.ConfigInit()
  db.Init()

  a := app.New()
  a.Settings().SetTheme(theme.DarkTheme())
  w := a.NewWindow("Sklad")
  w.Resize(fyne.NewSize(700, 700))

  var createContent *fyne.Container
  var createBar *fyne.Container

  purchasesTable.OnSelected = func(id widget.ListItemID) {
    purchaseName := widget.NewLabel(db.Purchases[id].Name)
    purchaseInfo := widget.NewLabel(
      fmt.Sprintf(
        "Client: %s\nVendor: %s\nStorer: %s",
        db.Purchases[id].ClientsName,
        db.Purchases[id].VendorsName,
        db.Purchases[id].StorersName,
      ),
    )

    item := container.NewBorder(
      container.NewHBox(
        widget.NewButtonWithIcon("", backIcon, func() {
          w.SetContent(createContent)
          purchasesTable.UnselectAll()
        }),
        widget.NewButtonWithIcon("", editIcon, func() {
          purchase := db.Purchases[id]
          prevContent := createContent

          purName := widget.NewEntry()
          purName.SetText(purchase.Name)
          vendorName := widget.NewEntry()
          vendorName.SetText(purchase.VendorsName)
          clientName := widget.NewEntry()
          clientName.SetText(purchase.ClientsName)
          storerName := widget.NewEntry()
          storerName.SetText(purchase.StorersName)

          editForm := widget.NewForm(
            widget.NewFormItem("Purchase Name:", purName),
            widget.NewFormItem("Vendor Name:", vendorName),
            widget.NewFormItem("Client Name:", clientName),
            widget.NewFormItem("Storer Name:", storerName),
            )
          editForm.OnCancel = func() {
            purName.SetText("")
            purName.Refresh()
            vendorName.SetText("")
            vendorName.Refresh()
            clientName.SetText("")
            clientName.Refresh()
            storerName.SetText("")
            storerName.Refresh()
            createContent = prevContent
            purchasesTable.UnselectAll()
            w.SetContent(createContent)
          }
          editForm.OnSubmit = func() {
            db.Edit(&purchase, purName.Text, vendorName.Text, clientName.Text, storerName.Text)
            createContent = prevContent
            db.Get()
            purchasesTable.UnselectAll()
            w.SetContent(createContent)
          }
          createContent = container.NewBorder(
            nil,
            nil,
            nil,
            nil,
            editForm,
          )
          w.SetContent(createContent)
        }),
        widget.NewButtonWithIcon("", deleteIcon, func() {
         
          dialog.ShowConfirm(
            "Confirmation deletion", 
            "Are you sure you want to delete this purchase?",
            func(b bool) {
              if b {
                db.Delete(db.Purchases[id])
                db.Get()
                purchasesTable.UnselectAll()
                w.SetContent(createContent)
              } 
            },
            w,
          )
        }),
      ),
      nil,
      nil,
      nil,
      container.NewVBox(purchaseName, purchaseInfo),
    )
    w.SetContent(item)
  }

  bar := container.NewVBox(container.NewHBox(
    widget.NewButtonWithIcon("", addIcon, func() {
      previousContent := createContent
      createContent = container.NewBorder(nil, nil, nil, nil, form)
      form.OnCancel = func() {
        purchaseEntry.Text = ""
        purchaseEntry.Refresh()
        vendorEntry.Text = ""
        vendorEntry.Refresh()
        storerEntry.Text = ""
        storerEntry.Refresh()
        clientEntry.Text = ""
        clientEntry.Refresh()
        createContent = previousContent
        db.Get()
        purchasesTable.UnselectAll()
        w.SetContent(createContent)
      } 
      form.OnSubmit = func() {
        purchase := models.Pucrhase {
          Name: purchaseEntry.Text,
        }
        storer := models.Storer{
          Name: storerEntry.Text,
        }
        vendor := models.Vendor{
          Name: vendorEntry.Text,
        }
        client := models.Client{
          Name: clientEntry.Text,
        }
        db.Add(purchase, client, vendor, storer)
        purchaseEntry.Text = ""
        purchaseEntry.Refresh()
        vendorEntry.Text = ""
        vendorEntry.Refresh()
        storerEntry.Text = ""
        storerEntry.Refresh()
        clientEntry.Text = ""
        clientEntry.Refresh()
        db.Get()
        createContent = previousContent
        purchasesTable.UnselectAll()
        w.SetContent(createContent)
      }
      w.SetContent(createContent)
    }),
    widget.NewSeparator(),
    widget.NewButtonWithIcon("", settingsIcon, func() {
      prevContent := createContent
      
      dbname := widget.NewEntry()
      dbname.SetText(config.AppConfig.DBName)
      
      var demo bool
      demoModeSwitch := widget.NewCheck("Demo mode", func(b bool) {
        demo = b
      })
      demoModeSwitch.Checked = config.AppConfig.DemoMode

      settingsForm := widget.NewForm(
        widget.NewFormItem("Database name: ", dbname),
        widget.NewFormItem("Turn on demo mode: ", demoModeSwitch),
        widget.NewFormItem("", canvas.NewText("This will erase all data from database", color.RGBA{190, 37, 40, 255})),
      )
      
      settingsForm.OnCancel = func() {
        createContent = prevContent
        w.SetContent(createContent)
      }
      settingsForm.OnSubmit = func() {
        if config.AppConfig.DBName != dbname.Text {
          config.AppConfig.DBName = dbname.Text
          db.Init()
        }
        config.AppConfig.DemoMode = demo
        config.AppConfig.Update()

        createContent = prevContent
        w.SetContent(createContent)
      }

      createContent = container.NewBorder(nil, nil, nil, nil, settingsForm)
      w.SetContent(createContent)
    }),
    widget.NewSeparator(),
    widget.NewButtonWithIcon("", searchIcon, func()  {
      previousContent := createContent
      searchItems := []models.PucrhaseOut{}
      purchaseName := widget.NewEntry()

      backBtn := widget.NewButtonWithIcon("", backIcon, func() {
        purchasesTable.UnselectAll()
        createContent = previousContent
        w.SetContent(createContent)
      })
      searchResults := widget.NewList(
        func() int {
          return len(searchItems)
        },
        func() fyne.CanvasObject {
          return widget.NewLabel("Default")
        },
        func(lii widget.ListItemID, co fyne.CanvasObject) {
          co.(*widget.Label).SetText(searchItems[lii].Name)
        },
      )

      searchResults.OnSelected = func(id widget.ListItemID) {
        innerPreviousContent := createContent 

        purName := widget.NewLabel(searchItems[id].Name)
        purInfo := widget.NewLabel(
          fmt.Sprintf(
            "Client: %s\nVendor: %s\nStorer: %s",
            searchItems[id].ClientsName,
            searchItems[id].VendorsName,
            searchItems[id].StorersName,
          ),
        )
        info := container.NewVBox(purName, purInfo)
        back := widget.NewButtonWithIcon("", backIcon, func() {
          createContent = innerPreviousContent
          searchResults.UnselectAll()
          w.SetContent(createContent)
        })

        edit := widget.NewButtonWithIcon("", editIcon, func() {
          purchase := db.Purchases[id]
          prevContent := createContent

          purName := widget.NewEntry()
          purName.SetText(purchase.Name)
          vendorName := widget.NewEntry()
          vendorName.SetText(purchase.VendorsName)
          clientName := widget.NewEntry()
          clientName.SetText(purchase.ClientsName)
          storerName := widget.NewEntry()
          storerName.SetText(purchase.StorersName)

          editForm := widget.NewForm(
            widget.NewFormItem("Purchase Name:", purName),
            widget.NewFormItem("Vendor Name:", vendorName),
            widget.NewFormItem("Client Name:", clientName),
            widget.NewFormItem("Storer Name:", storerName),
            )
          editForm.OnCancel = func() {
            purName.SetText("")
            purName.Refresh()
            vendorName.SetText("")
            vendorName.Refresh()
            clientName.SetText("")
            clientName.Refresh()
            storerName.SetText("")
            storerName.Refresh()
            createContent = prevContent
            purchasesTable.UnselectAll()
            w.SetContent(createContent)
          }
          editForm.OnSubmit = func() {
            db.Edit(&purchase, purName.Text, vendorName.Text, clientName.Text, storerName.Text)
            searchItems[id] = purchase
            createContent = innerPreviousContent
            db.Get()
            searchResults.UnselectAll()
            w.SetContent(createContent)
          }
          createContent = container.NewBorder(
            nil,
            nil,
            nil,
            nil,
            editForm,
          )
          w.SetContent(createContent)

        })
        remove := widget.NewButtonWithIcon("", deleteIcon, func() {
                   
          dialog.ShowConfirm(
            "Confirmation deletion", 
            "Are you sure you want to delete this purchase?",
            func(b bool) {
              if b {
                db.Delete(db.Purchases[id])
                db.Get()
                searchResults.UnselectAll()
                createContent = innerPreviousContent
                searchItems = append(searchItems[:id], searchItems[id+1:]...)
                w.SetContent(createContent)
              } 
            },
            w,
          )

        })
        
        createContent = container.NewBorder(container.NewHBox(back, edit, remove), nil, nil, nil, info)
        w.SetContent(createContent)
      }

      searchBtn := widget.NewButtonWithIcon("", searchIcon, func() {
        searchItems = db.Search(purchaseName.Text)
        searchResults.Refresh()
      })
      bar := container.NewVBox(container.NewHBox(searchBtn, backBtn), purchaseName)
      createContent = container.NewBorder(bar, nil, nil, nil, searchResults)

      w.SetContent(createContent)
    }),
  ),
  canvas.NewLine(color.Black))

  createBar = bar
  createContent = container.NewBorder(createBar, nil, nil, nil, purchasesTable)

  w.SetContent(createContent)
  w.ShowAndRun()
}
