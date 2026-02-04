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
import SubmissionPage from "./Submission"
import ParticipantDashboard from "./admin/Participants"
import ReviewFeed from "./admin/ReviewFeed"
import ParticipantInfo from "./admin/Participant"

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
        <Route path="/activities/:title" component={SubmissionPage} />
        <Route path="/admin/participants" component={ParticipantDashboard} />
        <Route path="/admin/participants/:id" component={ParticipantInfo} />
        <Route path="/admin/review" component={ReviewFeed} />
        <Route path="*" component={NotFound} />
      </Router>
    </MetaProvider>
  ),
  root!,
)
