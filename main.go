package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nicumicle/go-docker/internal/docker"
	"github.com/nicumicle/go-docker/internal/gui"
	"github.com/nicumicle/go-docker/pkg/colors"
)

func main() {

	g := &gui.Gui{}
	d := &docker.Docker{}

infiniteLoop:
	for {
		d.Clear()
		err := d.InitContainers()
		if err != nil {
			gui.DisplayErrorAndWaitForEnter(err.Error())
			break
		}

		gui.ClearScreen()
		g.ListContainers(d.Containers, d.Search)
		g.Render()

		option := gui.ReadCommand("Please enter your choice:")

		if option == gui.ActionClearSearch {
			d.Search = ""
			continue
		}
		if option == gui.ActionSearch {
			d.Search = gui.ReadCommand("What are you searching for:")

			continue
		}
		if option == gui.ActionRefresh {
			continue
		}
		if option == gui.ActionQuit {
			break
		}

		// Convert the selected container name to int
		index, err := strconv.ParseInt(option, 10, 0)
		if err != nil || index >= int64(len(d.Containers)) || (index < 0) {
			g.Error = fmt.Errorf("invalid container provided. Index %d is invalid", index)
			continue
		}

		selectedContainer := d.Containers[index]
		gui.ClearScreen()
		g.ListActions(selectedContainer)
		g.Render()
		action := gui.ReadCommand("Please select your action:")

		switch action {
		case "S":
			fmt.Println(colors.WriteYellow("Starting container..."))
			g.Error = selectedContainer.Start()
		case "s":
			fmt.Println(colors.WriteYellow("Stopping container..."))
			g.Error = selectedContainer.Stop()
		case "r":
			fmt.Println(colors.WriteYellow("Restarting container..."))
			g.Error = selectedContainer.Restart()
		case "p":
			fmt.Println(colors.WriteYellow("Pausing container..."))
			g.Error = selectedContainer.Pause()
		case "u":
			fmt.Println(colors.WriteYellow("Un-pausing container..."))
			g.Error = selectedContainer.UnPause()
		case "c":
			shell := gui.ReadCommand("Please select the shell [default /bin/bash]:")
			g.Error = selectedContainer.ConnectToContainer(shell)
		case "l":
			logs, err := selectedContainer.ShowLogs()
			if err != nil {
				g.Error = err

				continue
			}
			fmt.Println(string(*logs))
			gui.PressEnterToContinue("")
		case "d":
			confirm := gui.ReadCommand(
				fmt.Sprintf(
					"Are you sure you want to delete the %s container? [y/n]",
					colors.WriteYellow(selectedContainer.Name),
				),
			)
			if strings.Contains(strings.ToLower(confirm), "q") {
				break infiniteLoop
			}
			if !strings.Contains(strings.ToLower(confirm), "y") {
				continue
			}

			if err = selectedContainer.Delete(); err != nil {
				g.Error = err
			} else {
				fmt.Println(colors.WriteGreen("Container has been deleted."))
			}
			gui.PressEnterToContinue("")

		//--------
		case "b":
			// do nothing, just re-render everything
		case "q":
			break infiniteLoop
		default:
			g.Error = fmt.Errorf("%s is not a valid action", action)
		}
	}

	fmt.Println("Good bye!")
}
