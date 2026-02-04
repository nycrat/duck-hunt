import { Activity, Participant, Submission } from "./types"
import { getServerURL } from "./utils"

export const fetchWithMiddleware = async (
  path: string,
  init?: RequestInit,
): Promise<Response> => {
  const res = await fetch(`${getServerURL()}${path}`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
    ...init,
  })

  if (res.status === 401) {
    console.log("refresh session token")
    localStorage.removeItem("jwtToken")
  }

  return res
}

export const fetchParticipants = async (): Promise<Participant[]> => {
  const response = await fetchWithMiddleware(`/participants`)
  return response.json()
}

export const fetchParticipantInfo = async (
  id: number,
): Promise<Participant> => {
  const response = await fetchWithMiddleware(`/participants/${id}`)
  return response.json()
}

export const fetchActivities = async (): Promise<Activity[]> => {
  const response = await fetchWithMiddleware(`/activities`)
  return response.json()
}

export const fetchActivityInfo = async (
  title: string,
): Promise<Activity | null> => {
  const properlyEncodedTitle = title.replaceAll("'", "%27")
  const response = await fetchWithMiddleware(
    `/activities/${properlyEncodedTitle}`,
  )

  if (response.status !== 200) {
    return null
  }

  return response.json()
}

export const fetchPreviousSubmissions = async (params: {
  title: string
  id?: string
}): Promise<Submission[] | null> => {
  const response = await fetchWithMiddleware(
    params.id
      ? `/submissions/${params.title}/${params.id}`
      : `/submissions/${params.title}`,
  )

  if (response.status !== 200) {
    return null
  }

  return response.json()
}

export const postSubmission = async (
  title: string,
  image: Blob,
): Promise<boolean> => {
  console.log(title, image)
  const response = await fetchWithMiddleware(`/submissions/${title}`, {
    method: "POST",
    body: image,
  })

  return response.status === 200
}

export const postReview = async (
  submissionId: number,
  status: string,
): Promise<boolean> => {
  const response = await fetchWithMiddleware(
    `/submissions/review/${submissionId}`,
    {
      method: "POST",
      body: status,
    },
  )

  return response.status === 200
}

export const fetchUnreviewedSubmissions = async (): Promise<
  Submission[] | null
> => {
  const response = await fetchWithMiddleware(
    `/submissions/list/unreviewed/todo`,
  )

  if (response.status !== 200) {
    return null
  }

  return response.json()
}
