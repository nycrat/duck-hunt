import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"

const ParticipantInfo = () => {
  return (
    <main class="h-dvh p-10 flex flex-col gap-1">
      <Title>Participant | DuckHunt Admin</Title>
      <h1>Participant Dashboard</h1>
      admin page: participants
      <div class="grow" />
      <A href="/admin/participants">Go back</A>
    </main>
  )
}

export default ParticipantInfo
