import { defineConfig } from 'vitepress';

export default defineConfig({
  title: 'AI Starter',
  description: 'AI-driven repository setup and management tool',
  
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Guide', link: '/development/setup' },
      { text: 'ADRs', link: '/adr/' },
    ],

    sidebar: [
      {
        text: 'Architecture',
        items: [
          { text: 'Overview', link: '/adr/' },
          { text: 'Use MkDocs', link: '/adr/001-use-mkdocs' },
          { text: 'Perfect Commits', link: '/adr/002-perfect-commits' },
          { text: 'Track AI Prompts', link: '/adr/003-track-ai-prompts' },
          { text: 'Use Bun & TypeScript', link: '/adr/004-use-bun-typescript' },
        ],
      },
      {
        text: 'Development',
        items: [
          { text: 'Setup', link: '/development/setup' },
          { text: 'Contributing', link: '/development/contributing' },
          { text: 'Testing', link: '/development/testing' },
        ],
      },
      {
        text: 'Guides',
        items: [
          { text: 'Overview', link: '/guides/' },
        ],
      },
      {
        text: 'AI Prompts',
        items: [
          { text: 'Overview', link: '/prompts/' },
        ],
      },
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/xrendan/ai-starter' },
    ],

    search: {
      provider: 'local',
    },
  },
});
