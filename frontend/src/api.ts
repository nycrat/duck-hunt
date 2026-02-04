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

export const fetchParticipantList = async (): Promise<Participant[]> => {
  const response = await fetchWithMiddleware(`/participants`)
  return response.json()
}

export const fetchParticipant = async (id: number): Promise<Participant> => {
  const response = await fetchWithMiddleware(`/participants/${id}`)
  return response.json()
}

export const fetchActivityList = async (): Promise<Activity[]> => {
  const response = await fetchWithMiddleware(`/activities`)
  return response.json()
}

export const fetchActivity = async (
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

export const fetchActivitySubmissionList = async (params: {
  title: string
}): Promise<Submission[] | null> => {
  const response = await fetchWithMiddleware(
    `/activities/${params.title}/submissions`,
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
  const response = await fetchWithMiddleware(
    `/activities/${title}/submissions/`,
    {
      method: "POST",
      body: image,
    },
  )

  return response.status === 200
}

export const postReview = async (
  updatedSubmission: Submission,
): Promise<boolean> => {
  const response = await fetchWithMiddleware(`/admin/submissions`, {
    method: "PATCH",
    body: JSON.stringify({
      id: updatedSubmission.id,
      status: updatedSubmission.status,
      participant_id: updatedSubmission.participant_id,
    }),
  })

  return response.status === 200
}

export const fetchUnreviewedSubmissions = async (): Promise<
  Submission[] | null
> => {
  const response = await fetchWithMiddleware(`/admin/submissions`)

  if (response.status !== 200) {
    return null
  }

  return response.json()
}
