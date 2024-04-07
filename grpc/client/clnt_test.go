package client

import (
	"os"
	"testing"
)

func TestReadSettings(t *testing.T) {
	file, _ := os.Create("settings_test.txt")
	file.WriteString("{\"ip\":\"localhost\",\"port\":5000}\n{\"ip\":\"localhost\",\"port\":5001}")
	file.Seek(0, 0)
	defer os.Remove(file.Name())

	readSettings("settings_test.txt")
	if orchestrator.Ip != "localhost" || orchestrator.Port != 5000 {
		t.Error("not valid settings orchestrator!")
	}
	if agent.Id != 0 || agent.Ip != "localhost" || agent.Port != 5001 {
		t.Error("not valid settings agent!")
	}
}
