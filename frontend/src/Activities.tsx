import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import { createResource, createSignal, Match, Show, Switch } from "solid-js"
import { fetchActivities } from "./api"
import RedirectProvider from "./RedirectProvider"
import Footer from "./components/footer"

const Activities = () => {
  const [activities] = createResource(fetchActivities)
  const [sorting, setSorting] = createSignal("Name")
  const collator = new Intl.Collator()
  const sortingOptions = ["Name", "Points"]

  return (
    <RedirectProvider>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Activities | DuckHunt</Title>

        <h1>Activities</h1>

        <div class="flex gap-1">
          Sort by:
          {sortingOptions.map((option) => (
            <button
              onClick={() => setSorting(option)}
              class={`${sorting() === option ? "underline" : ""} hover:underline`}
            >
              {option}
            </button>
          ))}
        </div>

        <Show when={activities.loading}>loading...</Show>

        <Switch>
          <Match when={activities.error}>Error: {activities.error}</Match>
          <Match when={activities()}>
            <ol class="overflow-y-auto">
              {activities()!
                .toSorted(
                  sorting() === "Name"
                    ? (a, b) => collator.compare(a.title, b.title)
                    : (a, b) => b.points - a.points,
                )
                .map((activity) => (
                  <li>
                    <A href={`/activities/${activity.title}`}>
                      {`${activity.title} (${activity.points}pts)`}
                    </A>
                  </li>
                ))}
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
