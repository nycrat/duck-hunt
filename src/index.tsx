/* @refresh reload */
import { render } from "solid-js/web"
import "solid-devtools"

import "./paquito.css"
import "./chillax.css"
import "./index.css"
import { Route, Router } from "@solidjs/router"
import App from "./App"
import NotFound from "./NotFound"

const root = document.getElementById("root")

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    "Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?",
  )
}

render(
  () => (
    <Router>
      <Route path="/" component={App} />
      <Route path="*" component={NotFound} />
    </Router>
  ),
  root!,
)
