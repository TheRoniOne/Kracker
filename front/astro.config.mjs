// @ts-check
import { defineConfig } from "astro/config";

import node from "@astrojs/node";

import svelte from "@astrojs/svelte";

import tailwind from "@astrojs/tailwind";

// https://astro.build/config
export default defineConfig({
	adapter: node({
		mode: "standalone",
	}),

	integrations: [svelte(), tailwind({ applyBaseStyles: false })],
});
