package systemd

import (
	"fmt"
	"os/exec"
	"strings"
)

// Service ..
type Service struct {
	Name string // nginx.service or nginx
}

// Restart ..
func (s *Service) Restart() error {
	if err := exec.Command("systemctl", "restart", s.Name).Run(); err != nil {
		return err
	}
	return nil
}

// Start ..
func (s *Service) Start() error {
	if err := exec.Command("systemctl", "start", s.Name).Run(); err != nil {
		return err
	}
	return nil
}

// Stop ..
func (s *Service) Stop() error {
	if err := exec.Command("systemctl", "stop", s.Name).Run(); err != nil {
		return err
	}
	return nil
}

// Enable ..
func (s *Service) Enable() error {
	if err := exec.Command("systemctl", "enable", s.Name).Run(); err != nil {
		return err
	}
	return nil
}

// Disable ..
func (s *Service) Disable() error {
	if err := exec.Command("systemctl", "disable", s.Name).Run(); err != nil {
		return err
	}
	return nil
}

// IsActive ..
func (s *Service) IsActive() (bool, error) {
	out, err := exec.Command("systemctl", "is-active", s.Name).Output()
	if err != nil {
		return false, err
	}

	switch strings.TrimSpace(string(out)) {
	case "active":
		return true, nil
	case "inactive":
		return false, nil
	default:
		return false, fmt.Errorf("unknown status '%s'", string(out))
	}
}
