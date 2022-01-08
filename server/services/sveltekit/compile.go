/**
 * File: compile.go
 * Project: kit
 * File Created: 06 Jan 2022 10:37:33
 * Author: und3fined (me@und3fined.com)
 * -----
 * Last Modified: 08 Jan 2022 19:01:05
 * Modified By: und3fined (me@und3fined.com)
 * -----
 * Copyright (c) 2022 und3fined.com
 */
package sveltekit

import (
	"encoding/json"
	"errors"
	"time"

	"rogchap.com/v8go"

	"gosvelte/services/sveltekit/polyfills/fetch"
	"gosvelte/services/sveltekit/polyfills/url2"
)

type Compiler struct {
	ctx     *v8go.Context
	isolate *v8go.Isolate
	global  *v8go.ObjectTemplate
}

func (c *Compiler) NewIsolate() {
	c.isolate = v8go.NewIsolate()
}

func (c *Compiler) NewGlobal() {
	c.global = v8go.NewObjectTemplate(c.isolate)
}

func (c *Compiler) NewRequest(url, method string, headers map[string]string, body []byte) error {
	if c.isolate == nil {
		c.NewIsolate()
	}

	if c.global == nil {
		c.NewGlobal()
	}

	if err := fetch.InjectTo(c.isolate, c.global); err != nil {
		return err
	}

	getUrl := v8go.NewFunctionTemplate(c.isolate, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		val, _ := v8go.NewValue(c.isolate, url)
		return val
	})

	getMethod := v8go.NewFunctionTemplate(c.isolate, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		val, _ := v8go.NewValue(c.isolate, method)
		return val
	})

	getHeaders := v8go.NewFunctionTemplate(c.isolate, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		jsonHeaders, _ := json.Marshal(headers)
		val, _ := v8go.NewValue(c.isolate, string(jsonHeaders))
		return val
	})

	getBody := v8go.NewFunctionTemplate(c.isolate, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		return nil
	})

	if err := c.global.Set("getUrl", getUrl, v8go.ReadOnly); err != nil {
		return err
	}

	if err := c.global.Set("getMethod", getMethod, v8go.ReadOnly); err != nil {
		return err
	}

	if err := c.global.Set("getHeaders", getHeaders, v8go.ReadOnly); err != nil {
		return err
	}
	if err := c.global.Set("getBody", getBody, v8go.ReadOnly); err != nil {
		return err
	}

	return nil
}

func (c *Compiler) Exec(script string) (string, error) {
	c.ctx = v8go.NewContext(c.isolate, c.global)

	if err := url2.InjectTo(c.ctx); err != nil {
		return "", err
	}

	var val *v8go.Value
	var promise *v8go.Promise
	var err error

	if val, err = c.ctx.RunScript(script, "kit_main.js"); err != nil {
		if jsErr, ok := err.(*v8go.JSError); ok {
			return "", errors.New(jsErr.StackTrace)
		}

		return "", err
	}

	if promise, err = val.AsPromise(); err != nil {
		return "", err
	}

	done := make(chan bool, 1)
	defer close(done)

	go func() {
		for promise.State() == v8go.Pending {
			continue
		}
		done <- true
	}()

	select {
	case <-time.After(time.Second * 10):
		return "", errors.New("compile timeout")
	case <-done:
		return v8go.JSONStringify(c.ctx, promise.Result())
	}
}

func newCompiler() *Compiler {
	compiler := Compiler{}

	return &compiler
}
