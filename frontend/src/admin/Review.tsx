import { Title } from "@solidjs/meta"
import { useNavigate, useParams } from "@solidjs/router"
import AdminRoute from "./AdminRoute"
import { createResource, Match, Switch } from "solid-js"
import { fetchActivitySubmissionList } from "../api"
import ReviewList from "./ReviewList"

const SubmissionReviewPage = () => {
  const navigate = useNavigate()
  const params = useParams()

  const [submissions, { refetch: refetchSubmissions }] = createResource(
    { title: params.title, id: params.id },
    fetchActivitySubmissionList,
  )

  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>{decodeURI(params.title)} Review | DuckHunt Admin</Title>
        <h1>
          {decodeURI(params.title)} (Participant {params.id}) - Review
        </h1>

        <div class="h-px my-2 bg-black" />

        <Switch>
          <Match when={submissions() && submissions()!.length === 0}>
            No submissions for this activity
          </Match>
          <Match when={submissions() && submissions()!.length > 0}>
            <ReviewList
              submissions={submissions()!}
              onReview={refetchSubmissions}
            />
          </Match>
        </Switch>

        <div class="grow" />
        <button
          onClick={() => navigate(-1)}
          class="cursor-pointer underline text-blue-600 text-left"
        >
          Go back
        </button>
      </main>
    </AdminRoute>
  )
}

export default SubmissionReviewPage
