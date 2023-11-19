package docker

import (
	"encoding/json"
	"os/exec"
	"strings"
)

type Docker struct {
	Containers []Container
	Search     string
}

func (d *Docker) InitContainers() error {
	cmd := exec.Command("docker", "container", "ls", "--all", "--format", "{{json .}}")
	out, err := cmd.Output()

	if err != nil {
		return err
	}

	//Split the string into an array based on newline
	arr := strings.Split(string(out), "\n")
	for _, line := range arr {
		var c Container
		e := json.Unmarshal([]byte(line), &c)
		if e != nil {
			continue
		}
		c.Status = iniStatusFromString(c.InternalStatus)

		// Inspect container to get more details
		cmd := exec.Command("docker", "inspect", "-f", "{{json .}}", c.GetName())
		out, _ := cmd.Output()

		_ = json.Unmarshal(out, &c.Details)

		// Filter by name
		if !strings.Contains(strings.ToLower(c.GetName()), strings.ToLower(d.Search)) &&
			!strings.Contains(c.GetIp(), d.Search) {
			continue
		}

		d.Containers = append(d.Containers, c)
	}

	return nil
}

func (d *Docker) Clear() {
	d.Containers = make([]Container, 0)
}

func iniStatusFromString(status string) Status {
	if strings.Contains(status, "Pause") {
		return StatusPaused
	}

	if strings.Contains(status, "Up") {
		return StatusUp
	}

	if strings.Contains(status, "Exited") {
		return StatusStopped
	}
	if strings.Contains(status, "Restarting") {
		return StatusRestarting
	}

	return StatusUnknown
}
