import { Title } from "@solidjs/meta"
import { A, useNavigate } from "@solidjs/router"
import { createResource, Match, Show, Switch } from "solid-js"
import { fetchParticipants } from "./api"
import RedirectProvider from "./RedirectProvider"
import { getSessionId } from "./utils"

const Leaderboard = () => {
  const [participants] = createResource(fetchParticipants)
  const id = getSessionId()
  const navigate = useNavigate()

  return (
    <RedirectProvider>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Leaderboard | DuckHunt</Title>

        <h1>Leaderboard</h1>

        <Show when={participants.loading}>loading...</Show>

        <Switch>
          <Match when={participants.error}>Error: {participants.error}</Match>
          <Match when={participants()}>
            <div>
              <em>
                You currently have{" "}
                {participants()!.find((p) => p.id === id)?.score} points
              </em>
              <ol class="list-decimal list-inside">
                {participants()!
                  .toSorted((a, b) => b.score - a.score)
                  .map((participant) => (
                    <li>
                      {`${participant.name} (${participant.score}pts)`}
                      {participant.id === id && <em> &larr; this is you</em>}
                    </li>
                  ))}
              </ol>
            </div>
          </Match>
        </Switch>

        <div class="grow" />

        <button
          class="text-left underline text-blue-600 cursor-pointer"
          onClick={() => {
            if (confirm("Are you sure you want to log out?")) {
              localStorage.removeItem("jwtToken")
              navigate("/")
            }
          }}
        >
          Log out
        </button>

        <A href="/activities">Go to activities</A>
      </main>
    </RedirectProvider>
  )
}

export default Leaderboard
