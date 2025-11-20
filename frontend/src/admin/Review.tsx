import { Title } from "@solidjs/meta"
import { useNavigate, useParams } from "@solidjs/router"
import AdminRoute from "./AdminRoute"
import { createResource, Match, Switch } from "solid-js"
import { fetchPreviousSubmissions, postReview } from "../api"

const SubmissionReviewPage = () => {
  const navigate = useNavigate()
  const params = useParams()

  const [submissions] = createResource(
    { title: params.title, id: params.id },
    fetchPreviousSubmissions,
  )

  const statuses = ["unreviewed", "rejected", "accepted"]

  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>{decodeURI(params.title)} Review | DuckHunt Admin</Title>
        <h1>
          {decodeURI(params.title)} (Participant {params.id}) - Review
        </h1>

        <Switch>
          <Match when={submissions()}>
            <ol class="list-inside list-decimal overflow-y-scroll">
              {submissions()!.map((submission) => (
                <li>
                  <div class="flex gap-2">
                    {statuses.map((status) => (
                      <button
                        class={
                          "hover:bg-gray-500 cursor-pointer" +
                          (status === submission.status ? " bg-yellow-300" : "")
                        }
                        onClick={() => {
                          postReview(submission.id, status)
                        }}
                      >
                        {status}
                      </button>
                    ))}
                  </div>
                  <img
                    src={`data:image/jpeg;base64,${submission.image}`}
                    class="max-h-[40vh]"
                  />
                </li>
              ))}
            </ol>
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
