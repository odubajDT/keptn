package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/keptn/keptn/cli/pkg/credentialmanager"
	"github.com/keptn/keptn/cli/pkg/logging"
)

func init() {
	logging.InitLoggers(os.Stdout, os.Stdout, os.Stderr)
}

func CreateTmpShipyardFile(fileContent string) (string, error) {
	file, err := ioutil.TempFile("", "shipyard-*.yaml")
	if err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(file.Name(), []byte(fileContent), os.ModeAppend); err != nil {
		err = os.Remove(file.Name())
		if err != nil {
			return "", err
		}
		return "", err
	}
	return file.Name(), nil
}

const testingShipyard = `apiVersion: "spec.keptn.sh/0.2.3"
kind: "Shipyard"
metadata:
  name: "shipyard-podtato-ohead"
spec:
  stages:
    - name: "dev"
      sequences:
        - name: "delivery"
          tasks:
            - name: "deployment"
              properties:
                deploymentstrategy: "direct"
            - name: "release"

    - name: "staging"
      sequences:
        - name: "delivery"
          triggeredOn:
            - event: "dev.delivery.finished"
          tasks:
            - name: "deployment"
              properties:
                deploymentstrategy: "direct"
            - name: "release"

    - name: "production"
      sequences:
        - name: "delivery"
          triggeredOn:
            - event: "staging.delivery.finished"
          tasks:
            - name: "deployment"
              properties:
                deploymentstrategy: "direct"
            - name: "release"
`

// TestCreateProjectCmd tests the default use of the update project command
func TestUpdateProjectCmd(t *testing.T) {
	credentialmanager.MockAuthCreds = true
	checkEndPointStatusMock = true

	shipyardFilePath, err := CreateTmpShipyardFile(testingShipyard)
	cmd := fmt.Sprintf("update project sockshop -t token -u user -r https:// --mock --shipyard=%s", shipyardFilePath)
	_, err = executeActionCommandC(cmd)
	if err != nil {
		t.Errorf(unexpectedErrMsg, err)
	}
}

// TestUpdateProjectIncorrectProjectNameCmd tests whether the update project command aborts
// due to a project name with upper case character
func TestUpdateProjectIncorrectProjectNameCmd(t *testing.T) {
	credentialmanager.MockAuthCreds = true
	checkEndPointStatusMock = true

	shipyardFilePath, err := CreateTmpShipyardFile(testingShipyard)
	cmd := fmt.Sprintf("update project Sockshop2 -t token -u user -r https://github.com/user/upstream.git --mock --shipyard=%s", shipyardFilePath)
	_, err = executeActionCommandC(cmd)

	if !errorContains(err, "contains upper case letter(s) or special character(s)") {
		t.Errorf("missing expected error, but got %v", err)
	}
}
