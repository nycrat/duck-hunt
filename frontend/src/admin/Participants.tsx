import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"

const ParticipantsDashboard = () => {
  return (
    <main class="h-dvh p-10 flex flex-col gap-1">
      <Title>Participants | DuckHunt Admin</Title>
      <h1>Participants Dashboard</h1>
      admin page: participants
      <div class="grow" />
      <A href="/admin/activities">Go to activities</A>
    </main>
  )
}

export default ParticipantsDashboard
