import { Title } from "@solidjs/meta"
import { A, useParams } from "@solidjs/router"
import AdminRoute from "./AdminRoute"
import { createResource, Match, Show, Switch } from "solid-js"
import { fetchParticipantInfo } from "../api"

const ParticipantInfo = () => {
  const params = useParams()
  const id = parseInt(params.id)

  const [participant] = createResource(id, fetchParticipantInfo)

  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Participant | DuckHunt Admin</Title>
        <h1>
          Participant:{" "}
          <Show when={participant()}>
            {participant()!.name} ({participant()!.id})
          </Show>
        </h1>

        <Switch>
          <Match when={participant.loading}>loading...</Match>
          <Match when={participant.error}>Error {participant.error}</Match>
          <Match when={participant()}>Score: {participant()!.score}</Match>
        </Switch>

        <div class="grow" />
        <A href="/admin/participants">Go back</A>
      </main>
    </AdminRoute>
  )
}

export default ParticipantInfo
