/* @refresh reload */
import { render } from "solid-js/web"
import "solid-devtools"

import "./paquito.css"
import "./chillax.css"
import "./index.css"
import { Route, Router } from "@solidjs/router"
import Welcome from "./Welcome"
import NotFound from "./NotFound"
import { MetaProvider } from "@solidjs/meta"
import Leaderboard from "./Leaderboard"
import Activities from "./Activities"
import Submission from "./Submission"

const root = document.getElementById("root")

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    "Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?",
  )
}

render(
  () => (
    <MetaProvider>
      <Router>
        <Route path="/" component={Welcome} />
        <Route path="/leaderboard" component={Leaderboard} />
        <Route path="/activities" component={Activities} />
        <Route path="/activities/:title" component={Submission} />
        <Route path="*" component={NotFound} />
      </Router>
    </MetaProvider>
  ),
  root!,
)
