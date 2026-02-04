import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import AdminRoute from "./AdminRoute"
import { createResource, Match, Switch } from "solid-js"
import { fetchParticipants } from "../api"

const ParticipantsDashboard = () => {
  const [participants] = createResource(fetchParticipants)

  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Participants | DuckHunt Admin</Title>
        <h1>Participants List</h1>
        <Switch>
          <Match when={participants.loading}>loading...</Match>
          <Match when={participants.error}>Error {participants.error}</Match>
          <Match when={participants()}>
            <ul class="overflow-y-scroll">
              {participants()!.map((participant) => (
                <li>
                  <A
                    href={`/admin/participants/${participant.id}`}
                    class="text-xl"
                  >
                    {participant.name} ({participant.id})
                  </A>
                </li>
              ))}
            </ul>
          </Match>
        </Switch>
        <div class="grow" />
        <A href="/admin/review">Go to review</A>
      </main>
    </AdminRoute>
  )
}

export default ParticipantsDashboard
