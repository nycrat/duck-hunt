import { Title } from "@solidjs/meta"
import { A, useParams } from "@solidjs/router"
import { createResource, createSignal, Match, Show, Switch } from "solid-js"
import { imageToBlob, imageToImageURL } from "./utils"

interface Activity {
  title: string
  points: number
  description: string
}

interface Submission {
  status: "unreviewed" | "rejected" | "accepted"
  image: string
}

const fetchActivityInfo = async (title: string): Promise<Activity | null> => {
  const response = await fetch(`http://localhost:8000/activities/${title}`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })

  if (response.status !== 200) {
    return null
  }

  return response.json()
}

const fetchPreviousSubmissions = async (
  title: string,
): Promise<Submission[] | null> => {
  const response = await fetch(`http://localhost:8000/submissions/${title}`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })

  if (response.status !== 200) {
    return null
  }

  return response.json()
}

const postSubmission = async (title: string, image: Blob): Promise<boolean> => {
  console.log(title, image)
  const response = await fetch(`http://localhost:8000/submissions/${title}`, {
    method: "POST",
    body: image,
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })

  return response.status === 200
}

const MAX_FILE_SIZE_BYTES = 10_000_000

const SubmissionPage = () => {
  const params = useParams()

  const [activity] = createResource(params.title, fetchActivityInfo)
  const [submissions] = createResource(params.title, fetchPreviousSubmissions)

  const [image, setImage] = createSignal<File | null>(null)

  const [imagePreview] = createResource(image, imageToImageURL)

  return (
    <main class="h-screen p-10 flex flex-col gap-1">
      <Title>{decodeURI(params.title)} | DuckHunt</Title>

      <Show when={activity.loading}>loading...</Show>

      <Switch>
        <Match when={activity.error}>Error: {activity.error}</Match>
        <Match when={activity() === null}>
          <h1>404 activity not found</h1>
        </Match>
        <Match when={activity()}>
          <h1>{decodeURI(params.title)}</h1>
          <b>{activity()!.points}pts</b>
          <p>{activity()!.description}</p>

          <div class="h-px my-2 bg-black" />

          <form
            onSubmit={async (ev) => {
              ev.preventDefault()

              const imageFile = image()
              if (!imageFile) return

              const imageBlob = await imageToBlob(imageFile)
              if (imageBlob) postSubmission(params.title, imageBlob)
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
              onChange={(ev) => {
                if (!ev.target.files || !ev.target.files[0]) return
                const image = ev.target.files[0]

                // TODO: implement reducing file size instead of just denying
                if (image.size > MAX_FILE_SIZE_BYTES) {
                  alert("file too large")
                  return
                }

                setImage(image)
              }}
            />
            <input type="submit" />

            <Show when={imagePreview()}>
              <img src={imagePreview()} />
            </Show>
          </form>

          <Show when={submissions()}>
            {submissions()!.map((submission) => (
              <div>
                {submission.status}{" "}
                <img src={`data:image/jpeg;base64,${submission.image}`} />
              </div>
            ))}
          </Show>
        </Match>
      </Switch>

      <div class="grow" />

      <A href="/activities">Return to activities</A>
    </main>
  )
}

export default SubmissionPage
