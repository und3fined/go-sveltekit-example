/**
 * File: kit.go
 * Project: kit
 * File Created: 06 Jan 2022 10:37:37
 * Author: und3fined (me@und3fined.com)
 * -----
 * Last Modified: 08 Jan 2022 19:01:09
 * Modified By: und3fined (me@und3fined.com)
 * -----
 * Copyright (c) 2022 und3fined.com
 */
package sveltekit

import (
	"io/ioutil"
	"path/filepath"

	"gosvelte/utils/path"
)

type SSR struct {
	compiler *Compiler
	mainFile string
}

func (ssr *SSR) SetMainFile(filePath string) error {
	currentPath := path.Root()
	appPath := filepath.Join(currentPath, "./dest/client/main.js")
	fileContent, err := ioutil.ReadFile(appPath)

	if err != nil {
		return err
	}

	ssr.mainFile = string(fileContent)

	return nil
}

func (ssr *SSR) Render(url, method string, headers map[string]string, body []byte) (*Response, error) {
	if err := ssr.compiler.NewRequest(url, method, headers, body); err != nil {
		return nil, err
	}

	result, err := ssr.compiler.Exec(ssr.mainFile)
	if err != nil {
		return nil, err
	}

	return parseResponse(result)
}

func newSSR() *SSR {
	compiler := newCompiler()

	return &SSR{
		compiler: compiler,
	}
}
