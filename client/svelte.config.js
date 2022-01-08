/**
 * File: svelte.config.js
 * Project: client
 * File Created: 05 Jan 2022 22:06:34
 * Author: und3fined (me@und3fined.com)
 * -----
 * Last Modified: 08 Jan 2022 19:20:01
 * Modified By: und3fined (me@und3fined.com)
 * -----
 * Copyright (c) 2022 und3fined.com
 */
import { resolve, dirname } from 'path';
import preprocess from 'svelte-preprocess';

import adapter from './adapter-go/index.js';

const __dirname = dirname(new URL(import.meta.url).pathname);
const outpath = resolve(__dirname, '../dest/client');

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: [preprocess({})],

	kit: {
		adapter: adapter({
			out: outpath
		}),

		appDir: 'app',

		files: {
			assets: 'static',
			hooks: 'src/hooks',
			lib: 'src/lib',
			routes: 'src/routes',
			template: 'src/app.html'
		},

		target: '#gosvelte'
	}
};

export default config;
