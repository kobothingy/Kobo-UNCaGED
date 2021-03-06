// +build linux,!arm

// Copyright 2019 Sherman Perry

// This file is part of Kobo UNCaGED.

// Kobo UNCaGED is free software: you can redistribute it and/or modify
// it under the terms of the Affero GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Kobo UNCaGED is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with Kobo UNCaGED.  If not, see <https://www.gnu.org/licenses/>.

package main

import "fmt"

type kuUserPrint struct {
}

func newKuPrint(fontPath string) (*kuUserPrint, error) {
	kup := &kuUserPrint{}
	return kup, nil
}

func (kup *kuUserPrint) kuPrintln(section mboxSection, a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func (kup *kuUserPrint) kuClose() {
}
