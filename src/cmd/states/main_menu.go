package states

import (
	"ADPwn/database/project/model"
	"fmt"
)

type MainMenuState struct {
	Project model.Project
}

func (s *MainMenuState) Execute(context *Context) {
	fmt.Println("\nWelcome to main menu for project: " + s.Project.Name)

	context.SetState(nil)
}
