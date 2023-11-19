package docker

import (
	"strings"
	"time"
)

type Status string

const (
	StatusUp         Status = "UP"
	StatusPaused     Status = "PAUSED"
	StatusStopped    Status = "STOPPED"
	StatusRestarting Status = "RESTARTING"
	StatusUnknown    Status = "-"
)

type Container struct {
	Command        string `json:"Command"`
	CreatedAt      string `json:"CreatedAt"`
	ID             string `json:"ID"`
	Image          string `json:"Image"`
	Labels         string `json:"Labels"`
	LocalVolumes   string `json:"LocalVolumes"`
	Mounts         string `json:"Mounts"`
	Names          string `json:"Names"`
	Networks       string `json:"Networks"`
	Ports          string `json:"Ports"`
	RunningFor     string `json:"RunningFor"`
	Size           string `json:"Size"`
	State          string `json:"State"`
	InternalStatus string `json:"Status"`

	Name    string
	IP      string `json:"IP"`
	Status  Status
	Details Details
}

type Details struct {
	ID      string    `json:"Id"`
	Created time.Time `json:"Created"`
	Path    string    `json:"Path"`
	Args    []string  `json:"Args"`
	State   struct {
		Status     string    `json:"Status"`
		Running    bool      `json:"Running"`
		Paused     bool      `json:"Paused"`
		Restarting bool      `json:"Restarting"`
		OOMKilled  bool      `json:"OOMKilled"`
		Dead       bool      `json:"Dead"`
		Pid        int       `json:"Pid"`
		ExitCode   int       `json:"ExitCode"`
		Error      string    `json:"Error"`
		StartedAt  time.Time `json:"StartedAt"`
		FinishedAt time.Time `json:"FinishedAt"`
	} `json:"State"`
	Image           string `json:"Image"`
	ResolvConfPath  string `json:"ResolvConfPath"`
	HostnamePath    string `json:"HostnamePath"`
	HostsPath       string `json:"HostsPath"`
	LogPath         string `json:"LogPath"`
	Name            string `json:"Name"`
	RestartCount    int    `json:"RestartCount"`
	Driver          string `json:"Driver"`
	Platform        string `json:"Platform"`
	MountLabel      string `json:"MountLabel"`
	ProcessLabel    string `json:"ProcessLabel"`
	AppArmorProfile string `json:"AppArmorProfile"`
	ExecIDs         any    `json:"ExecIDs"`
	HostConfig      struct {
		Binds           []string `json:"Binds"`
		ContainerIDFile string   `json:"ContainerIDFile"`
		LogConfig       struct {
			Type   string `json:"Type"`
			Config struct {
			} `json:"Config"`
		} `json:"LogConfig"`
		NetworkMode  string `json:"NetworkMode"`
		PortBindings map[string]struct {
			HostIP   string `json:"HostIp"`
			HostPort string `json:"HostPort"`
		} `json:"PortBindings"`
		RestartPolicy struct {
			Name              string `json:"Name"`
			MaximumRetryCount int    `json:"MaximumRetryCount"`
		} `json:"RestartPolicy"`
		AutoRemove           bool     `json:"AutoRemove"`
		VolumeDriver         string   `json:"VolumeDriver"`
		VolumesFrom          []any    `json:"VolumesFrom"`
		ConsoleSize          []int    `json:"ConsoleSize"`
		CapAdd               any      `json:"CapAdd"`
		CapDrop              any      `json:"CapDrop"`
		CgroupnsMode         string   `json:"CgroupnsMode"`
		DNS                  any      `json:"Dns"`
		DNSOptions           any      `json:"DnsOptions"`
		DNSSearch            any      `json:"DnsSearch"`
		ExtraHosts           []string `json:"ExtraHosts"`
		GroupAdd             any      `json:"GroupAdd"`
		IpcMode              string   `json:"IpcMode"`
		Cgroup               string   `json:"Cgroup"`
		Links                any      `json:"Links"`
		OomScoreAdj          int      `json:"OomScoreAdj"`
		PidMode              string   `json:"PidMode"`
		Privileged           bool     `json:"Privileged"`
		PublishAllPorts      bool     `json:"PublishAllPorts"`
		ReadonlyRootfs       bool     `json:"ReadonlyRootfs"`
		SecurityOpt          any      `json:"SecurityOpt"`
		UTSMode              string   `json:"UTSMode"`
		UsernsMode           string   `json:"UsernsMode"`
		ShmSize              int      `json:"ShmSize"`
		Runtime              string   `json:"Runtime"`
		Isolation            string   `json:"Isolation"`
		CPUShares            int      `json:"CpuShares"`
		Memory               int      `json:"Memory"`
		NanoCpus             int      `json:"NanoCpus"`
		CgroupParent         string   `json:"CgroupParent"`
		BlkioWeight          int      `json:"BlkioWeight"`
		BlkioWeightDevice    any      `json:"BlkioWeightDevice"`
		BlkioDeviceReadBps   any      `json:"BlkioDeviceReadBps"`
		BlkioDeviceWriteBps  any      `json:"BlkioDeviceWriteBps"`
		BlkioDeviceReadIOps  any      `json:"BlkioDeviceReadIOps"`
		BlkioDeviceWriteIOps any      `json:"BlkioDeviceWriteIOps"`
		CPUPeriod            int      `json:"CpuPeriod"`
		CPUQuota             int      `json:"CpuQuota"`
		CPURealtimePeriod    int      `json:"CpuRealtimePeriod"`
		CPURealtimeRuntime   int      `json:"CpuRealtimeRuntime"`
		CpusetCpus           string   `json:"CpusetCpus"`
		CpusetMems           string   `json:"CpusetMems"`
		Devices              any      `json:"Devices"`
		DeviceCgroupRules    any      `json:"DeviceCgroupRules"`
		DeviceRequests       any      `json:"DeviceRequests"`
		MemoryReservation    int      `json:"MemoryReservation"`
		MemorySwap           int      `json:"MemorySwap"`
		MemorySwappiness     any      `json:"MemorySwappiness"`
		OomKillDisable       any      `json:"OomKillDisable"`
		PidsLimit            any      `json:"PidsLimit"`
		Ulimits              any      `json:"Ulimits"`
		CPUCount             int      `json:"CpuCount"`
		CPUPercent           int      `json:"CpuPercent"`
		IOMaximumIOps        int      `json:"IOMaximumIOps"`
		IOMaximumBandwidth   int      `json:"IOMaximumBandwidth"`
		MaskedPaths          []string `json:"MaskedPaths"`
		ReadonlyPaths        []string `json:"ReadonlyPaths"`
	} `json:"HostConfig"`
	GraphDriver struct {
		Data struct {
			LowerDir  string `json:"LowerDir"`
			MergedDir string `json:"MergedDir"`
			UpperDir  string `json:"UpperDir"`
			WorkDir   string `json:"WorkDir"`
		} `json:"Data"`
		Name string `json:"Name"`
	} `json:"GraphDriver"`
	Mounts []struct {
		Type        string `json:"Type"`
		Source      string `json:"Source"`
		Destination string `json:"Destination"`
		Mode        string `json:"Mode"`
		Rw          bool   `json:"RW"`
		Propagation string `json:"Propagation"`
	} `json:"Mounts"`
	Config struct {
		Hostname     string            `json:"Hostname"`
		Domainname   string            `json:"Domainname"`
		User         string            `json:"User"`
		AttachStdin  bool              `json:"AttachStdin"`
		AttachStdout bool              `json:"AttachStdout"`
		AttachStderr bool              `json:"AttachStderr"`
		ExposedPorts map[string]string `json:"ExposedPorts"`
		Tty          bool              `json:"Tty"`
		OpenStdin    bool              `json:"OpenStdin"`
		StdinOnce    bool              `json:"StdinOnce"`
		Env          []string          `json:"Env"`
		Cmd          []string          `json:"Cmd"`
		Image        string            `json:"Image"`
		Volumes      map[string]any    `json:"Volumes"`
		WorkingDir   string            `json:"WorkingDir"`
		Entrypoint   []string          `json:"Entrypoint"`
		OnBuild      any               `json:"OnBuild"`
		Labels       map[string]string `json:"Labels"`
	} `json:"Config"`
	NetworkSettings struct {
		Bridge                 string `json:"Bridge"`
		SandboxID              string `json:"SandboxID"`
		HairpinMode            bool   `json:"HairpinMode"`
		LinkLocalIPv6Address   string `json:"LinkLocalIPv6Address"`
		LinkLocalIPv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen"`
		Ports                  map[string][]struct {
			HostIP   string `json:"HostIp"`
			HostPort string `json:"HostPort"`
		} `json:"Ports"`
		SandboxKey             string             `json:"SandboxKey"`
		SecondaryIPAddresses   any                `json:"SecondaryIPAddresses"`
		SecondaryIPv6Addresses any                `json:"SecondaryIPv6Addresses"`
		EndpointID             string             `json:"EndpointID"`
		Gateway                string             `json:"Gateway"`
		GlobalIPv6Address      string             `json:"GlobalIPv6Address"`
		GlobalIPv6PrefixLen    int                `json:"GlobalIPv6PrefixLen"`
		IPAddress              string             `json:"IPAddress"`
		IPPrefixLen            int                `json:"IPPrefixLen"`
		IPv6Gateway            string             `json:"IPv6Gateway"`
		MacAddress             string             `json:"MacAddress"`
		Networks               map[string]Network `json:"Networks"`
	} `json:"NetworkSettings"`
}

type Network struct {
	IPAMConfig          any      `json:"IPAMConfig"`
	Links               any      `json:"Links"`
	Aliases             []string `json:"Aliases"`
	NetworkID           string   `json:"NetworkID"`
	EndpointID          string   `json:"EndpointID"`
	Gateway             string   `json:"Gateway"`
	IPAddress           string   `json:"IPAddress"`
	IPPrefixLen         int      `json:"IPPrefixLen"`
	IPv6Gateway         string   `json:"IPv6Gateway"`
	GlobalIPv6Address   string   `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen int      `json:"GlobalIPv6PrefixLen"`
	MacAddress          string   `json:"MacAddress"`
	DriverOpts          any      `json:"DriverOpts"`
}

func (c *Container) GetName() string {
	return c.Names
}

func (c *Container) GetStatus() string {
	return string(c.Status)
}

func (c *Container) GetIp() string {
	var ip string

	if len(c.Details.NetworkSettings.Networks) == 0 {
		return ip
	}
	for name, network := range c.Details.NetworkSettings.Networks {
		ip += name + " : " + network.IPAddress + "; "
	}

	return strings.TrimRight(ip, "; ")
}

func (c *Container) GetPortDetails() string {
	var output string
	if len(c.Details.NetworkSettings.Ports) == 0 {
		return output
	}

	for _, port := range c.Details.NetworkSettings.Ports {
		for _, x := range port {
			output += x.HostIP + ":" + x.HostPort + "; "
		}
	}

	return strings.TrimRight(output, "; ")
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
