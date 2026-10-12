import { defineConfig } from 'vite';

const base = process.env.GITHUB_PAGES === 'true' ? '/madmaid/' : '/';

export default defineConfig({ base });
