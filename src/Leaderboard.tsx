import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"

const Leaderboard = () => {
  const participants = [
    { name: "Kel", score: 1600 },
    { name: "Aly", score: 1000 },
    { name: "Ado", score: 1800 },
    { name: "Kiy", score: 1200 },
    { name: "Bec", score: 1300 },
    { name: "Kie", score: 2600 },
    { name: "Luc", score: 900 },
  ]

  return (
    <main class="h-screen p-10 flex flex-col gap-1">
      <Title>Leaderboard | DuckHunt</Title>

      <h1>Leaderboard</h1>

      <div>
        <ol class="list-decimal list-inside">
          {participants
            .toSorted((a, b) => b.score - a.score)
            .map((participant) => (
              <li>{`${participant.name} (${participant.score}pts)`}</li>
            ))}
        </ol>
      </div>

      <div class="grow" />

      <A href="/activities">Go to activities</A>
    </main>
  )
}

export default Leaderboard
