// Copyright (C) 2021 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpplink

import (
	"fmt"

	"git.fd.io/govpp.git/binapi/vpe"
	"github.com/pkg/errors"
)

func (v *VppLink) GetVPPVersion() (version string, err error) {
	v.Lock()
	defer v.Unlock()

	response := &vpe.ShowVersionReply{}
	request := &vpe.ShowVersion{}

	err = v.GetChannel().SendRequest(request).ReceiveReply(response)
	if err != nil {
		return "", errors.Wrapf(err, "ShowVersion failed: req %+v reply %+v", request, response)
	} else if response.Retval != 0 {
		return "", fmt.Errorf("ShowVersion failed (retval %d). Request: %+v", response.Retval, request)
	}
	return response.Version, nil
}
