import * as esbuild from "esbuild";
import chokidar from "chokidar";

const options = {
  entryPoints: ["src/index.ts"],
  bundle: true,
  outfile: "static/js/index.js",
  target: ["es2017"],
  minify: process.argv.includes("--prod"),
  platform: "browser",
};

const watcher = chokidar.watch("src", {
  ignoreInitial: true,
  persistent: true,
});

async function build() {
  try {
    await esbuild.build(options);
    console.log("[ESBUILD] rebuilt bundle");
  } catch (e) {
    console.log("[ESBUILD] failed to rebuild bundle:", e);
  }
}

build();

watcher.on("all", () => {
  build();
});
