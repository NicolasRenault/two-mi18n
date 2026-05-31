import { defineConfig } from "astro/config";
import icon from "astro-icon";

// https://astro.build/config
export default defineConfig({
  scopedStyleStrategy: "class",
  integrations: [icon()],
  build: {
    inlineStylesheets: "never",
  },
  site: "https://links.nicolasrenault.com",
});
