export interface Participant {
  id: number
  name: string
  score: number
}

export interface Activity {
  title: string
  points: number
  description: string
}

export type SubmissionStatus = "unreviewed" | "rejected" | "accepted"

export interface Submission {
  id: number
  status: SubmissionStatus
  image: string
  participant_id: number
  activity_title: string
}
