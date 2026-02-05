import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import AdminRoute from "./AdminRoute"
import { For, createResource, Match, Switch } from "solid-js"
import { fetchParticipantList } from "../api"

const ParticipantsDashboard = () => {
  const [participants] = createResource(fetchParticipantList)

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
              <For each={participants()!}>
                {(participant) => (
                  <li>
                    <A
                      href={`/admin/participants/${participant.id}`}
                      class="text-xl"
                    >
                      {participant.name} ({participant.id})
                    </A>
                  </li>
                )}
              </For>
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
