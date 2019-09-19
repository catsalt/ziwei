package main

import (
	"fmt"
	"strings"

	"github.com/catsalt/ziwei/zwds"

	"fyne.io/fyne"
	"fyne.io/fyne/app"

	// "fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type Person struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Calendar string `json:"calendar"`
	Birthday
}
type Birthday struct {
	Year   string `json:"year"`
	Month  string `json:"month"`
	Day    string `json:"day"`
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

var Hero zwds.Hero
var infoOne = Info{姓名: "命例", 性别: "女",
	公年: 1955, 公月: 3, 公日: 7, 公时: 4, 农年: 1955, 农月: 2, 农日: 14, 农时: 3,
	年干: "乙", 年支: "未", 月干: "己", 月支: "卯", 日干: "丁", 日支: "卯", 时干: "壬", 时支: "寅"}

func newForm() (info []string) {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Personal info")
	labels := []string{"Gender", "Calendar", "Name", "Year", "Month", "Day", "Hour", "Minute"}
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	childrn := []*widget.Select{
		widget.NewSelect(nums[0:3], nil), //year
		widget.NewSelect(nums, nil),
		widget.NewSelect(nums, nil),
		widget.NewSelect(nums, nil),
		widget.NewSelect(nums[0:2], nil), //month
		widget.NewSelect(nums, nil),
		widget.NewSelect(nums[0:4], nil), //day
		widget.NewSelect(nums, nil),
		widget.NewSelect(nums[0:3], nil), //hour
		widget.NewSelect(nums, nil),
		widget.NewSelect(nums[0:6], nil), //minute
		widget.NewSelect(nums, nil)}
	gender := widget.NewRadio([]string{"Male", "Female"}, nil)
	calendar := widget.NewRadio([]string{"Solar", "Lunar"}, nil)
	name := &widget.Entry{PlaceHolder: "Enter your name."}
	g := widget.NewHBox(widget.NewLabel(labels[0]), gender)
	c := widget.NewHBox(widget.NewLabel(labels[1]), calendar)
	n := widget.NewHBox(widget.NewLabel(labels[2]), name)
	year := widget.NewHBox(
		widget.NewLabel(labels[3]), childrn[0], childrn[1], childrn[2], childrn[3])
	monthDay := widget.NewHBox(
		widget.NewLabel(labels[4]), childrn[4], childrn[5],
		widget.NewLabel(labels[5]), childrn[6], childrn[7])
	hourMinute := widget.NewHBox(
		widget.NewLabel(labels[6]), childrn[8], childrn[9],
		widget.NewLabel(labels[7]), childrn[10], childrn[11])
	output := widget.NewLabel("")
	submit := widget.NewButton("Submit", func() {
		info = append(info, gender.Selected, calendar.Selected, name.Text)
		for _, v := range childrn {
			info = append(info, v.Selected)
		}
		output.SetText(strings.Join(info, " "))
		fmt.Println(info)
		return
	})

	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), g),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), c),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), n),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), year),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), monthDay),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), hourMinute),
		submit, output,
	))

	w.ShowAndRun()
	fmt.Println("hello.")
	return
}
func lifa() {

}
func main() {
	newForm()
}
