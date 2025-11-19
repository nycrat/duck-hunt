import { Title } from "@solidjs/meta"
import { A, useParams } from "@solidjs/router"
import AdminRoute from "./AdminRoute"
import { createResource, Match, Switch } from "solid-js"
import { fetchParticipantInfo, fetchParticipantSubmissions } from "../api"

const ParticipantInfo = () => {
  const params = useParams()
  const id = parseInt(params.id)

  const [participant] = createResource(id, fetchParticipantInfo)
  const [activities] = createResource(id, fetchParticipantSubmissions)

  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Participant | DuckHunt Admin</Title>
        <h1>Participant Dashboard</h1>

        <Switch>
          <Match when={participant.loading || activities.loading}>
            loading...
          </Match>
          <Match when={participant.error}>Error {participant.error}</Match>
          <Match when={activities.error}>Error {activities.error}</Match>
          <Match when={participant() && activities()}>
            <h2>
              {participant()!.name} ({participant()!.id})
            </h2>
            Score: {participant()!.score}
            <ul>
              {activities()!.map((activity) => (
                <li>
                  <A
                    href={`/admin/review/${activity.title}/${participant()!.id}/0`}
                    class="text-xl"
                  >
                    {activity.title} ({activity.count} submissions)
                  </A>
                </li>
              ))}
            </ul>
          </Match>
        </Switch>

        <div class="grow" />
        <A href="/admin/participants">Go back</A>
      </main>
    </AdminRoute>
  )
}

export default ParticipantInfo
