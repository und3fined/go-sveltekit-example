import { Adapter } from '@sveltejs/kit';

interface AdapterOptions {
	out?: string;
	precompress?: boolean;
}

declare function plugin(options?: AdapterOptions): Adapter;
export = plugin;
