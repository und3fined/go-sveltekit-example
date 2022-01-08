/**
 * File: index.js
 * Project: @sveltejs/adapter-node
 * File Created: 06 Jan 2022 14:49:27
 * Author: und3fined (me@und3fined.com)
 * -----
 * Last Modified: 08 Jan 2022 17:24:17
 * Modified By: und3fined (me@und3fined.com)
 * -----
 * Copyright (c) 2022 und3fined.com
 */
import { createReadStream, createWriteStream, existsSync, statSync, writeFileSync } from 'fs';
import { pipeline } from 'stream';
import glob from 'tiny-glob';
import { fileURLToPath } from 'url';
import { promisify } from 'util';
import zlib from 'zlib';
import { resolve, posix } from 'path';
import esbuild from 'esbuild';

const pipe = promisify(pipeline);

const files = fileURLToPath(new URL('./files', import.meta.url));

/** @type {import('.')} */
export default function ({
	out = 'build',
	precompress
} = {}) {
	return {
		name: '@sveltejs/adapter-go',

		async adapt(builder) {
      const tmp = builder.getBuildDirectory('go-tmp');

			builder.rimraf(out);
      builder.rimraf(tmp);

      builder.log.minor('Copying assets');
			builder.writeClient(`${out}`);
			builder.writeStatic(`${out}/static`);

			builder.log.minor('Prerendering static pages');
			await builder.prerender({
				dest: `${out}/prerendered`
			});

      builder.log.minor('Generating for go');
      const relativePath = posix.relative(tmp, builder.getServerDirectory());

			builder.copy(files, tmp, {
				replace: {
					APP: `${relativePath}/app.js`,
					MANIFEST: `${relativePath}/manifest.js`
				}
			});

      writeFileSync(
				`${tmp}/manifest.js`,
				`export const manifest = ${builder.generateManifest({
					relativePath
				})};\n`
			);

      await esbuild.build({
				entryPoints: [`${tmp}/main.js`],
				outfile: `${out}/main.js`,
				target: 'node14',
				bundle: true,
				platform: 'node',
			});

			if (precompress) {
				builder.log.minor('Compressing assets');
				await compress(`${out}/app`);
				await compress(`${out}/static`);
				await compress(`${out}/prerendered`);
			}

			builder.rimraf(tmp);
		}
	};
}

/**
 * @param {string} directory
 */
async function compress(directory) {
	if (!existsSync(directory)) {
		return;
	}

	const files = await glob('**/*.{html,js,json,css,svg,xml}', {
		cwd: directory,
		dot: true,
		absolute: true,
		filesOnly: true
	});

	await Promise.all(
		files.map((file) => Promise.all([compress_file(file, 'gz'), compress_file(file, 'br')]))
	);
}

/**
 * @param {string} file
 * @param {'gz' | 'br'} format
 */
async function compress_file(file, format = 'gz') {
	const compress =
		format == 'br'
			? zlib.createBrotliCompress({
					params: {
						[zlib.constants.BROTLI_PARAM_MODE]: zlib.constants.BROTLI_MODE_TEXT,
						[zlib.constants.BROTLI_PARAM_QUALITY]: zlib.constants.BROTLI_MAX_QUALITY,
						[zlib.constants.BROTLI_PARAM_SIZE_HINT]: statSync(file).size
					}
			  })
			: zlib.createGzip({ level: zlib.constants.Z_BEST_COMPRESSION });

	const source = createReadStream(file);
	const destination = createWriteStream(`${file}.${format}`);

	await pipe(source, compress, destination);
}