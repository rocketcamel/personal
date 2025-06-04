import * as esbuild from "esbuild";
import chokidar from "chokidar";

const prod = process.argv.includes("--prod");

const options = {
  entryPoints: ["src/index.ts"],
  bundle: true,
  outfile: "static/js/index.js",
  target: ["es2017"],
  minify: prod,
  platform: "browser",
};

async function watch() {
  const watcher = chokidar.watch("src", {
    ignoreInitial: true,
    persistent: true,
  });

  watcher.on("all", () => {
    build();
  });
}

async function build() {
  try {
    await esbuild.build(options);
    console.log("[ESBUILD] built bundle");
  } catch (e) {
    console.log("[ESBUILD] failed to build bundle:", e);
  }
}

build();

if (!prod) {
  watch();
}
