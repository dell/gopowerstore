/*
 *
 * Copyright © 2026 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package gopowerstore

import "testing"

// TestBuildMIPURL_G1_AsyncReplication validates that an IPv6 MIP produces a
// properly bracketed HTTPS URL — fixing the bare-concatenation bug in
// inttests/replication_async_test.go (G1).
func TestBuildMIPURL_G1_AsyncReplication(t *testing.T) {
	got := BuildMIPURL("2001:db8::1")
	want := "https://[2001:db8::1]/api/rest"
	if got != want {
		t.Errorf("BuildMIPURL(IPv6) = %q, want %q", got, want)
	}
}

// TestBuildMIPURL_G2_SyncReplication validates that an IPv6 MIP produces a
// properly bracketed HTTPS URL — fixing the bare-concatenation bug in
// inttests/replication_sync_test.go (G2).
func TestBuildMIPURL_G2_SyncReplication(t *testing.T) {
	got := BuildMIPURL("2607:f2b1:f1d0:768::158")
	want := "https://[2607:f2b1:f1d0:768::158]/api/rest"
	if got != want {
		t.Errorf("BuildMIPURL(IPv6 real addr) = %q, want %q", got, want)
	}
}

// TestBuildMIPURL_IPv4Unchanged confirms that IPv4 MIPs are unaffected (regression).
func TestBuildMIPURL_IPv4Unchanged(t *testing.T) {
	got := BuildMIPURL("192.168.1.100")
	want := "https://192.168.1.100/api/rest"
	if got != want {
		t.Errorf("BuildMIPURL(IPv4) = %q, want %q", got, want)
	}
}

// TestBuildMIPURL_HostnameUnchanged confirms FQDNs pass through unchanged.
func TestBuildMIPURL_HostnameUnchanged(t *testing.T) {
	got := BuildMIPURL("powerstore.example.com")
	want := "https://powerstore.example.com/api/rest"
	if got != want {
		t.Errorf("BuildMIPURL(FQDN) = %q, want %q", got, want)
	}
}

// TestBuildMIPURL_AlreadyBracketed guards against double-bracketing when a caller
// passes a pre-bracketed IPv6 MIP such as "[2001:db8::1]". Without the bracket-strip
// the function would emit "https://[[2001:db8::1]]/api/rest".
func TestBuildMIPURL_AlreadyBracketed(t *testing.T) {
	got := BuildMIPURL("[2001:db8::1]")
	want := "https://[2001:db8::1]/api/rest"
	if got != want {
		t.Errorf("BuildMIPURL(pre-bracketed IPv6) = %q, want %q", got, want)
	}
}
