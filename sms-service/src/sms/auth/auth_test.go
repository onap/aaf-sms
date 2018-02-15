/*
* Copyright 2018 TechMahindra
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package auth

import (
	"crypto/tls"
	"testing"
)

//Unit test to varify GetTLSconfig func and varify the tls config min version to be 771
//Assuming cert file name as auth_test.cert
func TestGetTLSConfig(t *testing.T) {
	_, err := GetTLSConfig("filedoesnotexist.cert")
	if err == nil {
		t.Errorf("Test failed due to expected behaviour")
	}
	tlsConfig, err := GetTLSConfig("../test/auth_test.cert")
	expected := tls.VersionTLS12
	actual := tlsConfig.MinVersion
	if tlsConfig != nil {
		if int(actual) != expected {
			t.Errorf("Test Failed due to version mismatch")
		}
	if tlsConfig == nil {
		t.Errorf("Test Failed due to GetTLSConfig returned nil")
	        }
	}
}
