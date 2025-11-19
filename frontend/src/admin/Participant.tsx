import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import AdminRoute from "./AdminRoute"

const ParticipantInfo = () => {
  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Participant | DuckHunt Admin</Title>
        <h1>Participant Dashboard</h1>
        admin page: participants
        <div class="grow" />
        <A href="/admin/participants">Go back</A>
      </main>
    </AdminRoute>
  )
}

export default ParticipantInfo
