import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import { createResource, Match, Show, Switch } from "solid-js"
import { getServerURL } from "./utils"

interface Participant {
  name: string
  score: number
}

const fetchParticipants = async (): Promise<Participant[]> => {
  const response = await fetch(`${getServerURL()}/participants`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  return response.json()
}

const Leaderboard = () => {
  const [participants] = createResource(fetchParticipants)

  return (
    <main class="h-screen p-10 flex flex-col gap-1">
      <Title>Leaderboard | DuckHunt</Title>

      <h1>Leaderboard</h1>

      <Show when={participants.loading}>loading...</Show>

      <Switch>
        <Match when={participants.error}>Error: {participants.error}</Match>
        <Match when={participants()}>
          <div>
            <ol class="list-decimal list-inside">
              {participants()!
                .toSorted((a, b) => b.score - a.score)
                .map((participant) => (
                  <li>{`${participant.name} (${participant.score}pts)`}</li>
                ))}
            </ol>
          </div>
        </Match>
      </Switch>

      <div class="grow" />

      <A href="/activities">Go to activities</A>
    </main>
  )
}

export default Leaderboard
