/**
 * File: path.go
 * Project: path
 * File Created: 06 Jan 2022 10:57:45
 * Author: und3fined (me@und3fined.com)
 * -----
 * Last Modified: 06 Jan 2022 11:49:25
 * Modified By: und3fined (me@und3fined.com)
 * -----
 * Copyright (c) 2022 und3fined.com
 */
package path

import (
	"os"
	"path/filepath"
	"strings"
)

func Root() string {
	ex, _ := os.Executable()

	if strings.Contains(ex, "go-build") {
		ex2, _ := os.Getwd()
		return filepath.Dir(ex2)
	}

	return filepath.Dir(ex)
}
