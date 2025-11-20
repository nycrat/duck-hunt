import { Activity, ActivitySubmissions, Participant } from "./types"
import { getServerURL } from "./utils"

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

export const fetchParticipantSubmissions = async (
  id: number,
): Promise<ActivitySubmissions[]> => {
  const response = await fetchWithMiddleware(
    `${getServerURL()}/participants/${id}/submission_counts`,
    {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
      },
    },
  )
  return response.json()
}
