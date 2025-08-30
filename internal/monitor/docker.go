package monitor

import (
	"os/exec"
	"strings"

	"tfnsysmonitor/internal/config"
	"tfnsysmonitor/internal/notify"
)

func CheckDocker(cfg *config.Config, d config.DockerMonitor) {
	cmd := exec.Command("docker", "inspect", "-f", "{{.State.Running}}", d.ContainerName)
	out, err := cmd.Output()
	if err != nil || !strings.Contains(string(out), "true") {
		notify.NotifyFailure(d.Name+" Container", d.Name+" Docker container not running: "+d.ContainerName, cfg)
	}
}
