import { Title } from "@solidjs/meta"
import { A, useNavigate } from "@solidjs/router"
import { createMemo, createResource, For, Match, Show, Switch } from "solid-js"
import { fetchParticipants } from "./api"
import RedirectProvider from "./RedirectProvider"
import { getSessionId } from "./utils"
import Footer from "./components/footer"

const Leaderboard = () => {
  const [participants] = createResource(fetchParticipants)
  const id = getSessionId()
  const navigate = useNavigate()

  const sortedParticipants = createMemo(() => {
    let prevScore = -1
    let prevRanking = 0
    return (participants() ?? [])
      .toSorted((a, b) => b.score - a.score)
      .map((participant, index) => {
        if (participant.score !== prevScore) {
          prevRanking = index + 1
          prevScore = participant.score
        }

        return { ...participant, ranking: prevRanking }
      })
  })

  const you = createMemo(() => sortedParticipants().find((p) => p.id === id))

  return (
    <RedirectProvider>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Leaderboard | DuckHunt</Title>

        <h1>Leaderboard</h1>

        <Show when={participants.loading}>loading...</Show>

        <Switch>
          <Match when={participants.error}>Error: {participants.error}</Match>
          <Match when={participants()}>
            <ul class="overflow-y-auto">
              <em>You currently have {you()?.score} points</em>
              <For each={sortedParticipants().slice(0, 20)}>
                {(participant) => (
                  <li>
                    {participant.ranking}. {participant.name} (
                    {participant.score}pts){" "}
                    {participant.id === id && <em> &larr; this is you</em>}
                  </li>
                )}
              </For>
              {sortedParticipants().findIndex((p) => p.id === id) >= 20 && (
                <li>
                  {you()?.ranking}. {you()?.name} ({you()?.score}pts){" "}
                  <em> &larr; this is you</em>
                </li>
              )}
            </ul>
          </Match>
        </Switch>

        <div class="grow" />

        <Footer>
          <>
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
          </>
        </Footer>
      </main>
    </RedirectProvider>
  )
}

export default Leaderboard
