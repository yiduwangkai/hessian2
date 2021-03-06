// Copyright 2016-2019 tianying Pan
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

package java_exception

import "strconv"

type IllegalFormatWidthException struct {
	SerialVersionUID     int64
	W                    int
	DetailMessage        string
	SuppressedExceptions []Throwabler
	StackTrace           []StackTraceElement
	Cause                Throwabler
}

func (e IllegalFormatWidthException) Error() string {
	return strconv.Itoa(e.W)
}

func (IllegalFormatWidthException) JavaClassName() string {
	return "java.util.IllegalFormatWidthException"
}
func NewIllegalFormatWidthException(w int) *IllegalFormatWidthException {
	return &IllegalFormatWidthException{W: w}
}
