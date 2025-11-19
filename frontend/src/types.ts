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

export interface Submission {
  status: "unreviewed" | "rejected" | "accepted"
  image: string
}

export interface ActivitySubmissions {
  title: string
  count: number
}
