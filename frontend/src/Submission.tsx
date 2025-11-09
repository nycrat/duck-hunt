import { Title } from "@solidjs/meta"
import { A, useParams } from "@solidjs/router"
import { createResource, Match, Show, Switch } from "solid-js"

interface Activity {
  title: string
  points: number
  description: string
}

const fetchActivityInfo = async (id: string): Promise<Activity | null> => {
  const response = await fetch(`http://localhost:8000/activities/${id}`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })

  if (response.status !== 200) {
    return null
  }

  return response.json()
}

const Submission = () => {
  const params = useParams()

  const [info] = createResource(params.title, fetchActivityInfo)

  return (
    <main class="h-screen p-10 flex flex-col gap-1">
      <Title>{decodeURI(params.title)} | DuckHunt</Title>

      <Show when={info.loading}>loading...</Show>

      <Switch>
        <Match when={info.error}>Error: {info.error}</Match>
        <Match when={info() === null}>
          <h1>404 activity not found</h1>
        </Match>
        <Match when={info()}>
          <h1>{decodeURI(params.title)}</h1>
          <b>{info()!.points}pts</b>
          <p>{info()!.description}</p>

          <div class="h-px my-2 bg-black" />

          <form
            onSubmit={(ev) => {
              ev.preventDefault()
            }}
          >
            <label
              for="image-upload"
              class="outline px-3 py-1 rounded-full hover:bg-gray-300"
            >
              Upload photo
            </label>
            <input
              id="image-upload"
              type="file"
              accept="image/*"
              capture="environment"
              class="hidden"
            />
          </form>
        </Match>
      </Switch>

      <div class="grow" />

      <A href="/activities">Return to activities</A>
    </main>
  )
}

export default Submission
