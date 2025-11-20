import { Title } from "@solidjs/meta"
import { A, useParams } from "@solidjs/router"
import { createResource, createSignal, Match, Show, Switch } from "solid-js"
import {
  getServerURL,
  imageToBlob,
  imageToImageURL,
  toTitleCase,
} from "./utils"
import { Activity, Submission } from "./types"
import { fetchWithMiddleware } from "./api"
import RedirectProvider from "./RedirectProvider"

const fetchActivityInfo = async (title: string): Promise<Activity | null> => {
  const properlyEncodedTitle = title.replaceAll("'", "%27")
  const response = await fetchWithMiddleware(
    `${getServerURL()}/activities/${properlyEncodedTitle}`,
    {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
      },
    },
  )

  if (response.status !== 200) {
    return null
  }

  return response.json()
}

const fetchPreviousSubmissions = async (
  title: string,
): Promise<Submission[] | null> => {
  const response = await fetchWithMiddleware(
    `${getServerURL()}/submissions/${title}`,
    {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
      },
    },
  )

  if (response.status !== 200) {
    return null
  }

  return response.json()
}

const postSubmission = async (title: string, image: Blob): Promise<boolean> => {
  console.log(title, image)
  const response = await fetchWithMiddleware(
    `${getServerURL()}/submissions/${title}`,
    {
      method: "POST",
      body: image,
      headers: {
        Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
      },
    },
  )

  return response.status === 200
}

const MAX_FILE_SIZE_BYTES = 10_000_000

const SubmissionPage = () => {
  const params = useParams()

  const [activity] = createResource(params.title, fetchActivityInfo)
  const [submissions, { refetch: refetchSubmissions }] = createResource(
    params.title,
    fetchPreviousSubmissions,
  )

  const [image, setImage] = createSignal<File | null>(null)

  const [imagePreview] = createResource(image, imageToImageURL)

  return (
    <RedirectProvider>
      <main class="h-dvh p-10 flex flex-col gap-1">
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
              class="flex gap-2"
              onSubmit={async (ev) => {
                ev.preventDefault()

                const imageFile = image()
                if (!imageFile) return

                const imageBlob = await imageToBlob(imageFile)
                if (imageBlob) await postSubmission(params.title, imageBlob)

                setImage(null)
                refetchSubmissions()
              }}
            >
              <label
                for="image-upload"
                class="outline px-3 py-1 rounded-full hover:bg-gray-300"
              >
                Select photo
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
              <label
                for="submit"
                class="outline px-3 py-1 rounded-full hover:bg-gray-300"
              >
                Submit
              </label>
              <input id="submit" type="submit" class="hidden" />
            </form>

            <Show when={imagePreview() && image()}>
              <img src={imagePreview()} />
            </Show>

            <div class="h-px my-2 bg-black" />

            <h2>Submissions</h2>

            <Show when={submissions()}>
              <ul class="list-inside list-decimal overflow-y-scroll">
                {submissions()!.map((submission) => (
                  <li>
                    {toTitleCase(submission.status)}
                    <img src={`data:image/jpeg;base64,${submission.image}`} />
                  </li>
                ))}
              </ul>
            </Show>
          </Match>
        </Switch>

        <div class="grow" />

        <A href="/activities">Return to activities</A>
      </main>
    </RedirectProvider>
  )
}

export default SubmissionPage
