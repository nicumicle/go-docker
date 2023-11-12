package docker

import (
	"os/exec"
	"strings"
)

type Docker struct {
	Containers []Container
}

func (d *Docker) InitContainers() error {
	cmd := exec.Command("docker", "container", "ls", "--all", "--format", "{{.Names}}###{{.Status}}###{{.}}")
	out, err := cmd.Output()

	if err != nil {
		return err
	}

	//Split the string into an array based on newline
	arr := strings.Split(string(out), "\n")
	for i := range arr {
		oneline := strings.Trim(arr[i], "'")

		line_arr := strings.Split(oneline, "###")
		if len(line_arr) < 2 {
			continue
		}
		cmd := exec.Command("docker", "inspect", "-f", "{{range.NetworkSettings.Networks}} {{.IPAddress}}{{end}}", line_arr[0])
		out, _ := cmd.Output()

		containerIp := strings.Trim(string(out), "\n")
		containerIp = strings.Trim(containerIp, " ")
		containerIp = strings.Replace(containerIp, " ", " - ", 1)

		d.Containers = append(d.Containers, Container{
			Name:   line_arr[0],
			Status: iniStatusFromString(line_arr[1]),
			IP:     containerIp,
		})
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

	return StatusUnknown
}
