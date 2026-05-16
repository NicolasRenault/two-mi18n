import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			title: "Two Mi18n",
			customCss: ["./src/styles/custom.css"],
			social: [
				{ icon: 'github', label: 'GitHub', href: 'https://github.com/NicolasRenault/two-mi18n' },
				{ icon: 'npm', label: 'NPM', href: 'https://www.npmjs.com/package/two-mi18n' }
			],
			sidebar: [
				{
					label: "Start Here",
					items: [
						{ label: "Getting Started", link: "/getting-started" },
					],
				},
				{
					label: "Concept",
					items: [
						{
							label: "Why Two Mi18n",
							link: "/concept/why-two-mi18n",
						},
						{ label: "Example", link: "/concept/example" },
					],
				},
				{
					label: "Usages",
					items: [
						{
							label: "Initialization",
							link: "/usages/initialization",
						},
						{ label: "translate()", link: "/usages/translate" },
						{
							label: "translateHTML()",
							link: "/usages/translate-html",
						},
					],
				},
				{
					label: "Behavior",
					items: [
						{ label: "Fallbacks", link: "/behavior/fallbacks" },
						{
							label: "Errors",
							link: "/behavior/errors",
						},
					],
				},
			],
		}),
	],
});
