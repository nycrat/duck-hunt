import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"

const Activities = () => {
  const activities = [
    { title: "Attend SUS's Workshop", points: 50, link: "/activities/tech" },
    {
      title: "Speed typing competition",
      points: 100,
      link: "/activities/tech",
    },
    { title: "SUS Jeopardy Event", points: 100, link: "/activities/tech" },
    { title: "Visit 5 Club Booths", points: 200, link: "/activities/tech" },
  ]

  return (
    <main class="h-screen p-10 flex flex-col gap-1">
      <Title>Activities | DuckHunt</Title>

      <h1>Activities</h1>

      {/* TODO implement sorting options */}
      <div>
        Sort by: <button>Name</button> <button class="underline">Points</button>{" "}
        <button>Todo</button>
      </div>

      <div>
        <ol>
          {activities
            .toSorted((a, b) => b.points - a.points)
            .map((activity) => (
              <li>
                <A href={activity.link}>
                  {`${activity.title} (${activity.points}pts)`}
                </A>
              </li>
            ))}
        </ol>
      </div>

      <div class="grow" />

      <A href="/leaderboard">Go to leaderboard</A>
    </main>
  )
}

export default Activities
