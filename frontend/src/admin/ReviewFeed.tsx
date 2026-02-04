import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import AdminRoute from "./AdminRoute"
import { createResource, Match, Switch } from "solid-js"
import { fetchUnreviewedSubmissions } from "../api"
import ReviewList from "./ReviewList"

const ReviewFeed = () => {
  const [submissions, { refetch }] = createResource(fetchUnreviewedSubmissions)

  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Review | DuckHunt Admin</Title>
        <h1>Review Feed</h1>
        <Switch>
          <Match when={submissions.loading}>loading...</Match>
          <Match when={submissions.error}>Error</Match>
          <Match when={submissions()}>
            <ReviewList submissions={submissions()!} onReview={refetch} />
          </Match>
        </Switch>
        <div class="grow" />
        <A href="/admin/participants">Go to participants</A>
      </main>
    </AdminRoute>
  )
}

export default ReviewFeed
