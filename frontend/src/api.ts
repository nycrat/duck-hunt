import { Participant } from "./types"
import { getServerURL } from "./utils"

export const fetchParticipants = async (): Promise<Participant[]> => {
  const response = await fetch(`${getServerURL()}/participants`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  return response.json()
}
