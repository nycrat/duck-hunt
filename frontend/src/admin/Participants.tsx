import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import AdminRoute from "./AdminRoute"

const ParticipantsDashboard = () => {
  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Participants | DuckHunt Admin</Title>
        <h1>Participants Dashboard</h1>
        admin page: participants
        <div class="grow" />
        <A href="/admin/activities">Go to activities</A>
      </main>
    </AdminRoute>
  )
}

export default ParticipantsDashboard
