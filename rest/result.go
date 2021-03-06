// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package rest

import (
	"gopkg.in/resty.v0"
)

func Result(resp *resty.Response, data interface{}) (err error) {
	var d *Model

	if resp.Result() != nil {
		d = resp.Result().(*Model)

		err = d.Data(data)
	}

	if err == nil {
		if (d != nil && !d.IsOK()) && resp.Error() != nil {
			err = resp.Error().(*Model).Error()
		}
	}

	return err
}
