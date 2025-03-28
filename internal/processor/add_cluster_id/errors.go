// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package add_cluster_id

import (
	"fmt"
)

type (
	configUnpackError struct{ cause error }
	computeIDError    struct{ cause error }
)

func makeErrConfigUnpack(cause error) configUnpackError {
	return configUnpackError{cause}
}
func (e configUnpackError) Error() string {
	return fmt.Sprintf("failed to unpack %v processor configuration: %v", processorName, e.cause)
}
func (e configUnpackError) Unwrap() error {
	return e.cause
}

func makeErrComputeID(cause error) computeIDError {
	return computeIDError{cause}
}
func (e computeIDError) Error() string {
	return fmt.Sprintf("failed to compute ID: %v", e.cause)
}
func (e computeIDError) Unwrap() error {
	return e.cause
}
