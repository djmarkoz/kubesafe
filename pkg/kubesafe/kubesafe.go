/*
Copyright © 2021 Mark Freriks <djmarkoz@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package kubesafe

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const ExpectedClusterFilename = ".kubesafe-expected-cluster"

var NotFound = errors.New("no expected cluster found")

func CurrentCluster() (string, error) {
	currentClusterBytes, err := exec.Command("kubectl", "config", "view", "-o=jsonpath={.current-context}").Output()
	if err != nil {
		return "", err
	}
	currentCluster := string(currentClusterBytes)
	currentCluster = strings.ReplaceAll(currentCluster, "\n", "")
	return currentCluster, nil
}

func Exec(cmd string, args []string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
	return command.Run()
}

func ExpectedCluster() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return expectedCluster(dir)
}

func expectedCluster(dir string) (string, error) {
	path := dir + "/" + ExpectedClusterFilename
	if _, err := os.Stat(path); err == nil {
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			return "", err
		}
		expectedCluster := string(dat)
		expectedCluster = strings.ReplaceAll(expectedCluster, "\n", "")
		return expectedCluster, nil
	} else if os.IsNotExist(err) {
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", NotFound
		}
		return expectedCluster(parent)
	} else {
		return "", err
	}
}

func ExpectedClusterList() ([]ExpectedClusterDefinition, error) {
	dir, err := os.Getwd()
	if err != nil {
		return []ExpectedClusterDefinition{}, err
	}
	return expectedClusterList(dir, []ExpectedClusterDefinition{})
}

type ExpectedClusterDefinition struct {
	Dir             string
	ExpectedCluster string
}

func expectedClusterList(dir string, list []ExpectedClusterDefinition) ([]ExpectedClusterDefinition, error) {
	path := dir + "/" + ExpectedClusterFilename
	if _, err := os.Stat(path); err == nil {
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		expectedCluster := string(dat)
		expectedCluster = strings.ReplaceAll(expectedCluster, "\n", "")
		list = append([]ExpectedClusterDefinition{{
			Dir:             dir,
			ExpectedCluster: expectedCluster,
		}}, list...)

		parent := filepath.Dir(dir)
		if parent == dir {
			return list, nil
		}
		return expectedClusterList(parent, list)
	} else if os.IsNotExist(err) {
		parent := filepath.Dir(dir)
		list = append([]ExpectedClusterDefinition{{
			Dir:             dir,
			ExpectedCluster: "",
		}}, list...)
		if parent == dir {
			if len(list) == 0 {
				return list, NotFound
			}
			return list, nil
		}
		return expectedClusterList(parent, list)
	} else {
		return list, err
	}
}
