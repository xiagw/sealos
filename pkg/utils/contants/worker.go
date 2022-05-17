/*
Copyright 2022 cuisongliu@qq.com.

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

package contants

import "path/filepath"

const (
	DefaultClusterFileName = "Clusterfile"
)

func Workdir() string {
	return filepath.Join(GetHomeDir(), ".sealos")
}

func ClusterDir(clusterName string) string {
	return filepath.Join(Workdir(), clusterName)
}

func Clusterfile(clusterName string) string {
	return filepath.Join(Workdir(), clusterName, DefaultClusterFileName)
}
