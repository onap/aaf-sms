/*
 * Copyright 2018 Intel Corporation, Inc
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

package backend

import (
	smsconfig "sms/config"
	"testing"
)

func TestInitSecretBackend(t *testing.T) {
	smsconfig.SMSConfig = &smsconfig.SMSConfiguration{
		VaultAddress: "http://localhost:8200",
	}
	sec, err := InitSecretBackend()
	// We expect an error to be returned as Init expects
	// backend to be running
	if err == nil {
		t.Fatal("InitSecretBackend : error creating")
	}
	if sec != nil {
		t.Fatal("InitSecretBackend: returned SecretBackend was *NOT* nil, expected nil")
	}
}
