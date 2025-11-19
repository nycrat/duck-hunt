import { Activity, Participant } from "./types"
import { getServerURL } from "./utils"

export const fetchParticipants = async (): Promise<Participant[]> => {
  const response = await fetch(`${getServerURL()}/participants`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  return response.json()
}

export const fetchParticipantInfo = async (
  id: number,
): Promise<Participant> => {
  const response = await fetch(`${getServerURL()}/participants/${id}`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  return response.json()
}

export const fetchActivities = async (): Promise<Activity[]> => {
  const response = await fetch(`${getServerURL()}/activities`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  return response.json()
}
