import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import {
  For,
  createResource,
  createSignal,
  Match,
  Show,
  Switch,
} from "solid-js"
import { fetchActivityList } from "./api"
import RedirectProvider from "./RedirectProvider"
import Footer from "./components/footer"

const Activities = () => {
  const [activities] = createResource(fetchActivityList)
  const [sorting, setSorting] = createSignal("Name")
  const collator = new Intl.Collator()
  const sortingOptions = ["Name", "Points"]

  const sortedActivities = () => {
    if (!activities()) return []
    return activities()!.toSorted(
      sorting() === "Name"
        ? (a, b) => collator.compare(a.title, b.title)
        : (a, b) => b.points - a.points,
    )
  }

  return (
    <RedirectProvider>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Activities | DuckHunt</Title>

        <h1>Activities</h1>

        <div class="flex gap-1">
          Sort by:
          <For each={sortingOptions}>
            {(option) => (
              <button
                onClick={() => setSorting(option)}
                class={`${sorting() === option ? "underline" : ""} hover:underline`}
              >
                {option}
              </button>
            )}
          </For>
        </div>

        <Show when={activities.loading}>loading...</Show>

        <Switch>
          <Match when={activities.error}>Error: {activities.error}</Match>
          <Match when={activities()}>
            <ol class="overflow-y-auto">
              <For each={sortedActivities()}>
                {(activity) => (
                  <li>
                    <A href={`/activities/${activity.title}`}>
                      {`${activity.title} (${activity.points}pts)`}
                    </A>
                  </li>
                )}
              </For>
            </ol>
          </Match>
        </Switch>

        <div class="grow" />

        <Footer>
          <A href="/leaderboard">Go to leaderboard</A>
        </Footer>
      </main>
    </RedirectProvider>
  )
}

export default Activities
