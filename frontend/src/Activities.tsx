import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import { createResource, Match, Show, Switch } from "solid-js"
import { fetchActivities } from "./api"
import RedirectProvider from "./RedirectProvider"

const Activities = () => {
  const [activities] = createResource(fetchActivities)

  return (
    <RedirectProvider>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Activities | DuckHunt</Title>

        <h1>Activities</h1>

        {/* TODO implement sorting options */}
        <div>
          Sort by: <button>Name</button>{" "}
          <button class="underline">Points</button> <button>Todo</button>
        </div>

        <Show when={activities.loading}>loading...</Show>

        <Switch>
          <Match when={activities.error}>Error: {activities.error}</Match>
          <Match when={activities()}>
            <div>
              <ol>
                {activities()!
                  .toSorted((a, b) => b.points - a.points)
                  .map((activity) => (
                    <li>
                      <A href={`/activities/${activity.title}`}>
                        {`${activity.title} (${activity.points}pts)`}
                      </A>
                    </li>
                  ))}
              </ol>
            </div>
          </Match>
        </Switch>

        <div class="grow" />

        <A href="/leaderboard">Go to leaderboard</A>
      </main>
    </RedirectProvider>
  )
}

export default Activities
