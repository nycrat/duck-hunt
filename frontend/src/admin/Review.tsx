import { Title } from "@solidjs/meta"
import { useNavigate, useParams } from "@solidjs/router"
import AdminRoute from "./AdminRoute"
import { createResource, Match, Switch } from "solid-js"
import { fetchPreviousSubmissions, postReview } from "../api"
import { toTitleCase } from "../utils"

const SubmissionReviewPage = () => {
  const navigate = useNavigate()
  const params = useParams()

  const [submissions, { refetch: refetchSubmissions }] = createResource(
    { title: params.title, id: params.id },
    fetchPreviousSubmissions,
  )

  const statuses = ["unreviewed", "rejected", "accepted"]
  const statusColors = ["bg-yellow-300/80", "bg-red-300/80", "bg-green-300/80"]

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
            <ul class="overflow-y-scroll space-y-4">
              {submissions()!.map((submission) => (
                <li class="space-y-2">
                  <div class="flex gap-2">
                    {statuses.map((status, i) => (
                      <button
                        class={
                          "border rounded-full px-2 hover:bg-gray-200/50 cursor-pointer " +
                          (status === submission.status ? statusColors[i] : "")
                        }
                        onClick={async () => {
                          await postReview(submission.id, status)
                          refetchSubmissions()
                        }}
                      >
                        {toTitleCase(status)}
                      </button>
                    ))}
                  </div>
                  <img
                    src={`data:image/jpeg;base64,${submission.image}`}
                    class="max-h-[40vh]"
                  />
                </li>
              ))}
            </ul>
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
