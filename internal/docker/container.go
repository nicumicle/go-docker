package docker

type Status string

const (
	StatusUp      Status = "UP"
	StatusPaused  Status = "PAUSED"
	StatusStopped Status = "STOPPED"
	StatusUnknown Status = "-"
)

type Container struct {
	Name   string
	IP     string
	Status Status
}

func (c *Container) GetName() string {
	return c.Name
}

func (c *Container) GetStatus() string {
	return string(c.Status)
}

func (c *Container) GetIp() string {
	return c.IP
}

// ConnectToContainer will execute docker exec in a container
// shell can be: "/bin/bash", "/sh", ...
func (c *Container) ConnectToContainer(shell string) error {
	if shell == "" {
		shell = "/bin/bash"
	}

	err := run("exec", "-itu", "0", c.GetName(), shell)
	if err != nil {
		return err
	}

	return nil
}

// ShowLogs will display the logs of a container
func (c *Container) ShowLogs() (*[]byte, error) {
	out, err := getOutput("logs", c.GetName())
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Container) Start() error {
	return run("start", c.GetName())
}

func (c *Container) Stop() error {
	return run("stop", c.GetName())
}

func (c *Container) Restart() error {
	return run("restart", c.GetName())
}

func (c *Container) Pause() error {
	return run("pause", c.GetName())
}

func (c *Container) UnPause() error {
	return run("unpause", c.GetName())
}

func (c *Container) Delete() error {
	return run("rm", c.GetName(), "--force")
}
