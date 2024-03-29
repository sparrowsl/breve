import { presetForms } from "@julr/unocss-preset-forms";
import extractorSvelte from "@unocss/extractor-svelte";
import {
  defineConfig,
  presetIcons,
  presetUno,
  presetWebFonts,
  transformerVariantGroup,
} from "unocss";

export default defineConfig({
  presets: [
    presetForms(),
    presetIcons({
      extraProperties: {
        display: "inline-block",
        "vertical-align": "middle",
      },
    }),
    presetUno(),
    presetWebFonts({
      fonts: {
        roboto: "Roboto",
        raleway: "Raleway",
      },
    }),
  ],
  extractors: [extractorSvelte()],
  transformers: [transformerVariantGroup()],
  shortcuts: { container: "max-w-2xl mx-auto" },
});
