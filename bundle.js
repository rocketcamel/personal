import * as esbuild from "esbuild";

await esbuild.build({
  entryPoints: ["src/index.ts"],
  bundle: true,
  outfile: "static/js/index.js",
  target: ["es6"],
  minify: process.argv.includes("--prod"),
  platform: "browser",
});
