package states

import (
	"ADPwn/database/project/model"
	worker "ADPwn/tools"
	"os"

	"github.com/rivo/tview"
)

type MainMenuState struct {
	Project model.Project
	App     *tview.Application
}

func (s *MainMenuState) Execute(context *Context) {

	title := tview.NewTextView().
		SetText("ADPwn - Main Menu - " + s.Project.Name).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Eine einzige Liste für alle Optionen
	mainMenuList := tview.NewList()

	mainMenuList.AddItem("[green::b] 🎯 "+s.Project.Name+"[-:-:-]", "", 0, nil)

	// Konfigurationsoptionen mit Überschrift hinzufügen
	mainMenuList.AddItem("[yellow::b] ⚙️ Configuration Options[-:-:-]", "", 0, nil)

	mainMenuList.AddItem("Add Single Host", "", '1', func() {
		context.SetState(&StartMenuState{App: s.App})
	})
	mainMenuList.AddItem("Add Host Range", "", '2', func() {
		context.SetState(&MainMenuAddHostRange{App: s.App, Project: s.Project})
	})
	mainMenuList.AddItem("Add User", "", '3', func() {
		context.SetState(&StartMenuState{App: s.App})
	})

	// Trennüberschrift hinzufügen
	mainMenuList.AddItem("[red::b] ⚔️ Execution Options[-:-:-]", "", 0, nil)

	// Ausführungsoptionen hinzufügen
	mainMenuList.AddItem("Run Enumeration", "", '4', func() {
		context.SetState(nil)
		s.App.Stop()
		enumerator := worker.Enumerator{}
		enumerator.Run(s.Project)
	})

	mainMenuList.AddItem("Exit", "", 'q', func() {
		context.SetState(nil)
		s.App.Stop()
		os.Exit(0)
	})

	mainMenuList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		// Wenn die Überschrift ausgewählt wird, springe zur nächsten Zeile
		if shortcut == 0 {
			if index < mainMenuList.GetItemCount()-1 {
				mainMenuList.SetCurrentItem(index + 1)
			} else {
				mainMenuList.SetCurrentItem(0)
			}
		}
	})

	// Layout mit der Liste erstellen
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 3, 0, false).      // Titel oben
		AddItem(mainMenuList, 0, 1, true) // Eine Liste, die alles enthält

	// Fokus auf die Liste setzen
	s.App.SetRoot(flex, true).SetFocus(mainMenuList)
}

func (s *MainMenuState) addUser(context *Context) {

}

func (s *MainMenuState) addHosts(context *Context) {

}
