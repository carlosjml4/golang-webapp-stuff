// Copyright 2013 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"
)

// A SampleValue is a representation of a value for a given sample at a given
// time.
type SampleValue float64

func (v SampleValue) Equal(o SampleValue) bool {
	return v == o
}

func (v SampleValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%f"`, v)), nil
}

func (v SampleValue) String() string {
	return fmt.Sprint(float64(v))
}
