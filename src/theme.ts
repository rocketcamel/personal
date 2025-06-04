import Alpine from "alpinejs";

type ThemeStore = {
  on: boolean;
  toggle: () => void;
};

let theme = localStorage.getItem("theme");

if (!theme) {
  theme = window.matchMedia("(prefers-color-scheme: dark)").matches
    ? "dark"
    : "light";
}

Alpine.store("darkMode", {
  on: theme === "dark" ? true : false,

  toggle() {
    this.on = !this.on;
    localStorage.setItem("theme", this.on ? "dark" : "light");
  },
} as ThemeStore);
