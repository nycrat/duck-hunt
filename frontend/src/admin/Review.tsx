import { Title } from "@solidjs/meta"
import { useNavigate } from "@solidjs/router"
import AdminRoute from "./AdminRoute"

const SubmissionReviewPage = () => {
  const navigate = useNavigate()

  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Participants | DuckHunt Admin</Title>
        <h1>Participants Dashboard</h1>
        admin page: review
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
