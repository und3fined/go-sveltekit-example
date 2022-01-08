/**
 * File: kit.go
 * Project: sveltekit
 * File Created: 08 Jan 2022 14:15:25
 * Author: und3fined (me@und3fined.com)
 * -----
 * Last Modified: 08 Jan 2022 16:42:03
 * Modified By: und3fined (me@und3fined.com)
 * -----
 * Copyright (c) 2022 und3fined.com
 */
package sveltekit

type SvelteKit struct {
	opts Options
}

type Option func(*Options)

func (kit *SvelteKit) Render(url, method string, headers map[string]string, body []byte) (*Response, error) {
	if err := kit.opts.SSR.SetMainFile(*kit.opts.MainFile); err != nil {
		return nil, err
	}

	return kit.opts.SSR.Render(url, method, headers, body)
}

func NewKit(opt ...Option) *SvelteKit {
	opts := newOptions(opt...)

	return &SvelteKit{opts: opts}
}
