package docker

import "testing"

func TestClear(t *testing.T) {
	t.Parallel()

	docker := &Docker{}
	docker.Containers = append(docker.Containers, Container{Name: "test"})
	docker.Clear()

	if len(docker.Containers) != 0 {
		t.Fatal("containers should be empty")
	}
}
func TestIniStatusFromString(t *testing.T) {
	t.Parallel()
	tt := map[string]struct {
		status         string
		expectedStatus Status
	}{
		"should detect up status": {
			status:         "Up",
			expectedStatus: StatusUp,
		},
		"should detect paused status": {
			status:         "Paused",
			expectedStatus: StatusPaused,
		},
		"should detect stopped status": {
			status:         "Exited",
			expectedStatus: StatusStopped,
		},
		"handle unwknown status": {
			status:         "",
			expectedStatus: StatusUnknown,
		},
	}

	for tname, tc := range tt {
		tc := tc
		t.Run(tname, func(t *testing.T) {
			t.Parallel()
			result := iniStatusFromString(tc.status)
			if result != tc.expectedStatus {
				t.Fatalf("invalid status received. Expecte %s but received %s", tc.expectedStatus, result)
			}
		})
	}
}
