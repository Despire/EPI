// Copyright (c) 2019 Matúš Mrekaj.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License

package stacksandqueues

import "testing"

func TestNormalizePathname(t *testing.T) {
	tests := []struct {
		//in
		path string
		//out
		want string
		err  error
	}{
		{path: "././aloha", want: "aloha", err: nil},
		{path: "/usr/lib/../bin/gcc", want: "/usr/bin/gcc", err: nil},
		{path: "", want: "", err: ErrInvalidPath},
		{path: "123/456", want: "123/456", err: nil},
		{path: "/123/456", want: "/123/456", err: nil},
		{path: "usr/lib/../bin/gcc", want: "usr/bin/gcc", err: nil},
		{path: "./../", want: "..", err: nil},
		{path: "../../local", want: "../../local", err: nil},
		{path: "./.././../local", want: "../../local", err: nil},
		{path: "/foo/../foo/./../", want: "/", err: nil},
		{path: "/", want: "/", err: nil},
		{path: "///", want: "/", err: nil},
		{path: "/.", want: "/", err: nil},
		{path: "/./", want: "/", err: nil},
		{path: "..", want: "..", err: nil},
		{path: "../", want: "..", err: nil},
		{path: "/GCxwlF", want: "/GCxwlF", err: nil},
		{path: "/laomWCV/", want: "/laomWCV", err: nil},
		{path: "/BmhAXRf/DzsRpMOCq", want: "/BmhAXRf/DzsRpMOCq", err: nil},
		{path: "/xGoueEW", want: "/xGoueEW", err: nil},
		{path: "/yjXFKG/", want: "/yjXFKG", err: nil},
		{path: "/LPyY", want: "/LPyY", err: nil},
		{path: "/iThjBHJf/", want: "/iThjBHJf", err: nil},
		{path: "/lreBzkEmeq", want: "/lreBzkEmeq", err: nil},
	}

	for _, test := range tests {
		got, err := NormalizePath(test.path)

		if err != test.err {
			t.Errorf("result mismatch. got: %v; want: %v", err, test.want)
		}

		if got != test.want {
			t.Errorf("result mismatch. got: %v; want: %v", got, test.want)
		}
	}
}
