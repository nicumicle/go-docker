package gui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nicumicle/go-docker/internal/docker"
	"github.com/nicumicle/go-docker/pkg/colors"
)

type Gui struct {
	output string
	Error  error
}

const (
	ActionSearch      = "s"
	ActionClearSearch = "cs"
	ActionRefresh     = "r"
	ActionQuit        = "q"
)

func (g *Gui) ListContainers(c []docker.Container, search string) {
	g.output = ""
	if g.Error != nil {
		g.output += colors.WriteRed("Error: "+g.Error.Error()) + "\n"
		g.Error = nil
	}

	g.output += colors.WriteBlue("Available Containers:") + "\n"
	if search != "" {
		g.output += colors.WriteGray(fmt.Sprintf("Searching by \"%s\":", colors.WriteBlue(search))) + "\n"
	}
	g.output += fmt.Sprintf("%3s | %8s | %50s | %10s\n", "ID", "STATUS", "CONTAINER NAME", "NETWORK : IP")
	g.output += fmt.Sprintln(strings.Repeat("-", 90))

	// Render the containers here
	for i, container := range c {
		g.output += fmt.Sprintf(
			"[%12s] %17s | %50s | %10s\n",
			colors.WriteYellow(strconv.FormatInt(int64(i), 10)),
			iniStatusColorFromString(container.GetStatus()),
			container.GetName(),
			colors.WriteBlue(container.GetIp()),
		)
	}

	g.output += fmt.Sprintln()
	g.output += fmt.Sprintf("Type a number betweet %s to select a container.\n",
		colors.WriteYellow(fmt.Sprintf("%d-%d", 0, len(c)-1)),
	)
	g.output += fmt.Sprintln("Type", colors.WriteYellow(ActionSearch), "to search.")
	if search != "" {
		g.output += fmt.Sprintln("Type", colors.WriteYellow(ActionClearSearch), "to clear the search.")
	}
	g.output += fmt.Sprintln("Type", colors.WriteYellow(ActionRefresh), "to refresh the list.")
	g.output += fmt.Sprintln("Type", colors.WriteYellow(ActionQuit), "to exit.")
}

// ListActions Display the screen where we can choose actions for our selected container
func (g *Gui) ListActions(c docker.Container) {
	g.output = fmt.Sprintln("You have entered selected: ", colors.WriteYellow(c.GetName()))
	g.output += fmt.Sprintln("Status", iniStatusColorFromString(c.GetStatus()))
	g.output += fmt.Sprintln("IP: ", colors.WriteGray(c.GetIp()))
	g.output += fmt.Sprintln("Ports:", colors.WriteGray(c.GetPortDetails()))
	g.output += "\n"
	g.output += fmt.Sprintln("Options:")
	switch c.GetStatus() {
	case string(docker.StatusUp):
		g.output += fmt.Sprintln("\t[", colors.WriteYellow("s"), "] Stop")
		g.output += fmt.Sprintln("\t[", colors.WriteYellow("r"), "] Restart")
		g.output += fmt.Sprintln("\t[", colors.WriteYellow("p"), "] Pause")
		g.output += fmt.Sprintln("\t", strings.Repeat("-", 16))
		g.output += fmt.Sprintln("\t[", colors.WriteYellow("c"), "] Connect")
		g.output += fmt.Sprintln("\t[", colors.WriteYellow("l"), "] Show logs")
	case string(docker.StatusStopped):
		g.output += fmt.Sprintln("\t[", colors.WriteYellow("S"), "] Start")
	case string(docker.StatusPaused):
		g.output += fmt.Sprintln("\t[", colors.WriteYellow("u"), "] Unpause")
	case string(docker.StatusRestarting):
	case string(docker.StatusUnknown):
	}

	g.output += fmt.Sprintln("\t[", colors.WriteYellow("d"), "] Delete")
	g.output += fmt.Sprintln("\t", strings.Repeat("-", 16))
	g.output += fmt.Sprintln("\t[", colors.WriteRed("b"), "] Go back")
}

func (g *Gui) Render() {
	fmt.Println(g.output)
}

func iniStatusColorFromString(status string) string {
	switch true {
	case strings.Contains(status, "UP"):
		return colors.WriteGreen(status)
	case strings.Contains(status, "STOPPED"):
		return colors.WriteRed(status)
	case strings.Contains(status, "PAUSED"):
		return colors.WriteYellow(status)
	default:
		return colors.WriteGray(status)
	}
}
