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

		g.ListContainers(d.Containers)
		g.Render()
		//read option from command line
		option := gui.ReadCommand("Please enter the container id:")

		if option == "r" {
			continue
		}
		if option == "q" {
			break
		}

		// Convert the selected container name to int
		index, err := strconv.ParseInt(option, 10, 0)
		if err != nil || index >= int64(len(d.Containers)) || (index < 0) {
			gui.DisplayErrorAndWaitForEnter("Invalid container provided.")
			continue
		}

		selectedContainer := d.Containers[index]
		g.ListActions(selectedContainer)
		g.Render()
		action := gui.ReadCommand("Please select your action:")

		switch action {
		case "S":
			fmt.Println(colors.WriteYellow("Starting container.."))
			err := selectedContainer.Start()
			if err != nil {
				gui.DisplayErrorAndWaitForEnter(err.Error())
			}
		case "s":
			fmt.Println(colors.WriteYellow("Stopping container.."))
			err := selectedContainer.Stop()
			if err != nil {
				gui.DisplayErrorAndWaitForEnter(err.Error())
			}
		case "r":
			fmt.Println(colors.WriteYellow("Restarting container.."))
			err := selectedContainer.Restart()
			if err != nil {
				gui.DisplayErrorAndWaitForEnter(err.Error())
			}
		case "p":
			fmt.Println(colors.WriteYellow("Pausing container.."))
			err := selectedContainer.Pause()
			if err != nil {
				gui.DisplayErrorAndWaitForEnter(err.Error())
			}
		case "u":
			fmt.Println(colors.WriteYellow("Unpausing container.."))
			err := selectedContainer.UnPause()
			if err != nil {
				gui.DisplayErrorAndWaitForEnter(err.Error())
			}
		case "c":
			shell := gui.ReadCommand("Please select the shell [default /bin/bash]:")

			fmt.Println("selected shell: ", shell)
			err := selectedContainer.ConnectToContainer(shell)
			if err != nil {
				gui.DisplayErrorAndWaitForEnter(err.Error())
			}
		case "l":
			logs, err := selectedContainer.ShowLogs()
			if err != nil {
				gui.DisplayErrorAndWaitForEnter(err.Error())

				continue
			}
			fmt.Println(string(*logs))
			gui.PressEnterToContinue("")
		case "d":
			confirm := gui.ReadCommand(
				"Are you sure you want to delete " +
					colors.WriteYellow(selectedContainer.Name) +
					" container? [y/n]",
			)
			if strings.Contains(confirm, "q") {
				break infiniteLoop
			}
			if !strings.Contains(confirm, "y") {
				continue
			}

			err = selectedContainer.Delete()
			if err != nil {
				gui.DisplayErrorAndWaitForEnter(err.Error())
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
			gui.DisplayErrorAndWaitForEnter("Invalid action provided")
		}
	}

	fmt.Println("Good bye!")
}
