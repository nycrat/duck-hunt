import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import { createResource, Match, Show, Switch } from "solid-js"

interface Activity {
  title: string
  points: number
  description: string
}

const fetchActivities = async (): Promise<Activity[]> => {
  const response = await fetch("http://localhost:8000/activities", {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  return response.json()
}

const Activities = () => {
  const [activities] = createResource(fetchActivities)

  return (
    <main class="h-screen p-10 flex flex-col gap-1">
      <Title>Activities | DuckHunt</Title>

      <h1>Activities</h1>

      {/* TODO implement sorting options */}
      <div>
        Sort by: <button>Name</button> <button class="underline">Points</button>{" "}
        <button>Todo</button>
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
  )
}

export default Activities
