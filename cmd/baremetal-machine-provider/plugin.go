/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package main

import (
	"os"
	"path/filepath"

	"tkestack.io/tke/cmd/baremetal-machine-provider/app"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	err = os.Chdir(filepath.Dir(exe) + "/..")
	if err != nil {
		panic(err)
	}
	app.NewApp("baremetal-machine-provider").Run()
}
