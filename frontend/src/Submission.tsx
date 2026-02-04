import { Title } from "@solidjs/meta"
import { A, useParams } from "@solidjs/router"
import {
  createResource,
  createSignal,
  JSX,
  Match,
  Show,
  Switch,
} from "solid-js"
import { imageToBlob, imageToImageURL, toTitleCase } from "./utils"
import {
  fetchActivity,
  fetchActivitySubmissionList,
  postSubmission,
} from "./api"
import RedirectProvider from "./RedirectProvider"

const MAX_FILE_SIZE_BYTES = 10_000_000

const SubmissionPage = () => {
  const params = useParams()

  const [activity] = createResource(params.title, fetchActivity)
  const [submissions, { refetch: refetchSubmissions }] = createResource(
    { title: params.title },
    fetchActivitySubmissionList,
  )

  const [image, setImage] = createSignal<File | null>(null)

  const [imagePreview] = createResource(image, imageToImageURL)

  const handleUploadImage: JSX.EventHandler<HTMLInputElement, Event> = (ev) => {
    if (!ev.currentTarget.files || !ev.currentTarget.files[0]) return
    const image = ev.currentTarget.files[0]

    // TODO: implement reducing file size instead of just denying
    if (image.size > MAX_FILE_SIZE_BYTES) {
      alert("file too large")
      return
    }

    setImage(image)
  }

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
                for="photo-input"
                class="outline px-3 py-1 rounded-full hover:bg-gray-300"
              >
                Capture
              </label>
              <input
                id="photo-input"
                type="file"
                accept="image/*"
                capture="environment"
                class="hidden"
                onChange={handleUploadImage}
              />
              <label
                for="image-upload"
                class="outline px-3 py-1 rounded-full hover:bg-gray-300"
              >
                Upload
              </label>
              <input
                id="image-upload"
                type="file"
                accept="image/*"
                class="hidden"
                onChange={handleUploadImage}
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
