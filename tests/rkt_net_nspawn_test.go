// Copyright 2016 The rkt Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build host coreos src

package main

import "testing"

func TestNetHost(t *testing.T) {
	NewNetHostTest().Execute(t)
}

func TestNetHostConnectivity(t *testing.T) {
	NewNetHostConnectivityTest().Execute(t)
}

func TestNetDefaultPortFwdConnectivity(t *testing.T) {
	NewNetDefaultPortFwdConnectivityTest(
		PortFwdCase{"172.16.28.1", "--net=default", true},
		PortFwdCase{"127.0.0.1", "--net=default", true},
	).Execute(t)
}

func TestNetNone(t *testing.T) {
	NewNetNoneTest().Execute(t)
}

func TestNetCustomMacvlan(t *testing.T) {
	NewNetCustomMacvlanTest().Execute(t)
}

func TestNetCustomBridge(t *testing.T) {
	NewNetCustomBridgeTest().Execute(t)
}

func TestNetOverride(t *testing.T) {
	NewNetOverrideTest().Execute(t)
}

func TestNetCustomPtp(t *testing.T) {
	// PTP means connection Point-To-Point. That is, connections to other pods/containers should be forbidden
	NewNetCustomPtpTest(true).Execute(t)
}

func TestNetDefaultConnectivity(t *testing.T) {
	NewNetDefaultConnectivityTest().Execute(t)
}

// TODO: fix this for kvm, see https://github.com/coreos/rkt/issues/2533
func TestNetLongName(t *testing.T) {
	nt := networkTemplateT{
		Name:   "thisnameiswaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaytoolong",
		Type:   "ptp",
		IpMasq: true,
		Ipam: ipamTemplateT{
			Type:   "host-local",
			Subnet: "11.11.6.0/24",
			Routes: []map[string]string{
				{"dst": "0.0.0.0/0"},
			},
		},
	}
	testNetCustomNatConnectivity(t, nt)
}
