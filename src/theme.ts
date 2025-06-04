import Alpine from "alpinejs";

type ThemeStore = {
  on: boolean;
  toggle: () => void;
};

Alpine.store("darkMode", {
  on: true,

  toggle() {
    this.on = !this.on;
  },
} as ThemeStore);
