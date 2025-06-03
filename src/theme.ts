import Alpine from "alpinejs";

Alpine.store("darkMode", {
  on: true,

  toggle() {
    this.on = !this.on;
  },
});
