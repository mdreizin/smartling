// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package main

import (
	"github.com/Fitbit/smartling/di"
	"gopkg.in/urfave/cli.v1"
	"path/filepath"
)

func injectDiContainerAction(c *cli.Context) error {
	filename, err := filepath.Abs(c.GlobalString(projectFileFlagName))

	if err == nil {
		var container *di.Container

		opts := di.Options{
			Filename: filename,
			Verbose:  c.GlobalBool(verboseFlagName),
		}

		container, err = di.Setup(&opts)

		c.App.Metadata[containerMetadataKey] = container
	}

	return err
}
