package cmd

import "github.com/manifoldco/promptui"

func NewSelectTemplates() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Help: Printer.Sprintf("select help"),
	}
}
