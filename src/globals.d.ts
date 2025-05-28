import type Alpine from "alpinejs";
import type htmx from "htmx.org";

declare global {
  interface Window {
    htmx: htmx;
    Alpine: Alpine;
  }
}
