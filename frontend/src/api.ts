import { Activity, Participant, Submission } from "./types"
import { getServerURL } from "./utils"

// TODO: refactor to reduce code repeating

export const fetchWithMiddleware = async (
  input: RequestInfo | URL,
  init?: RequestInit,
): Promise<Response> => {
  const res = await fetch(input, init)

  if (res.status === 401) {
    console.log("refresh session token")
    localStorage.removeItem("jwtToken")
  }

  return res
}

export const fetchParticipants = async (): Promise<Participant[]> => {
  const response = await fetchWithMiddleware(`${getServerURL()}/participants`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  return response.json()
}

export const fetchParticipantInfo = async (
  id: number,
): Promise<Participant> => {
  const response = await fetchWithMiddleware(
    `${getServerURL()}/participants/${id}`,
    {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
      },
    },
  )
  return response.json()
}

export const fetchActivities = async (): Promise<Activity[]> => {
  const response = await fetchWithMiddleware(`${getServerURL()}/activities`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  return response.json()
}

export const fetchActivityInfo = async (
  title: string,
): Promise<Activity | null> => {
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

export const fetchPreviousSubmissions = async (params: {
  title: string
  id?: string
}): Promise<Submission[] | null> => {
  const response = await fetchWithMiddleware(
    params.id
      ? `${getServerURL()}/submissions/${params.title}/${params.id}`
      : `${getServerURL()}/submissions/${params.title}`,
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

export const postSubmission = async (
  title: string,
  image: Blob,
): Promise<boolean> => {
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

export const postReview = async (
  submissionId: number,
  status: string,
): Promise<boolean> => {
  const response = await fetchWithMiddleware(
    `${getServerURL()}/submissions/review/${submissionId}`,
    {
      method: "POST",
      body: status,
      headers: {
        Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
      },
    },
  )

  return response.status === 200
}

export const fetchUnreviewedSubmissions = async (): Promise<
  Submission[] | null
> => {
  const response = await fetchWithMiddleware(
    `${getServerURL()}/submissions/list/unreviewed/todo`,
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
