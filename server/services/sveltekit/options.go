/**
 * File: options.go
 * Project: sveltekit
 * File Created: 08 Jan 2022 14:09:04
 * Author: und3fined (me@und3fined.com)
 * -----
 * Last Modified: 08 Jan 2022 16:06:58
 * Modified By: und3fined (me@und3fined.com)
 * -----
 * Copyright (c) 2022 und3fined.com
 */
package sveltekit

var (
	defaultMainFile = "client/main.js"
)

type Options struct {
	MainFile *string
	SSR      *SSR
}

func newOptions(opt ...Option) Options {
	opts := Options{}

	for _, o := range opt {
		o(&opts)
	}

	if opts.MainFile == nil {
		opts.MainFile = &defaultMainFile
	}

	if opts.SSR == nil {
		opts.SSR = newSSR()
	}

	return opts
}

func UseMain(path string) Option {
	return func(o *Options) {
		o.MainFile = &path
	}
}

func UseSSR() Option {
	return func(o *Options) {
		o.SSR = newSSR()
	}
}
