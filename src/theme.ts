import Alpine from "alpinejs";

type ThemeStore = {
  on: boolean;
  toggle: () => void;
};

let theme = localStorage.getItem("theme");

const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;

if (!theme) {
  theme = prefersDark ? "dark" : "light";
}

document.cookie = "theme=" + theme;

Alpine.store("darkMode", {
  on: theme === "dark" ? true : false,

  toggle() {
    this.on = !this.on;
    localStorage.setItem("theme", this.on ? "dark" : "light");
  },
} as ThemeStore);
