import { Accessor } from "solid-js"
import { postReview } from "../api"
import { Submission, SubmissionStatus } from "../types"
import { toTitleCase } from "../utils"

type Props = {
  submissions: Accessor<Submission[] | undefined>
  onReview: (submission: Submission) => void
}

const ReviewList = (props: Props) => {
  const statuses: SubmissionStatus[] = ["unreviewed", "rejected", "accepted"]
  const statusColors = ["bg-yellow-300/80", "bg-red-300/80", "bg-green-300/80"]

  return (
    <ul class="overflow-y-scroll space-y-4">
      {props.submissions()!.map((submission) => (
        <li class="space-y-2">
          <strong>
            {submission.activity_title} - participant{" "}
            {submission.participant_id}
          </strong>
          <div class="flex gap-2">
            {statuses.map((status, i) => (
              <button
                class={
                  "border rounded-full px-2 hover:bg-gray-200/50 cursor-pointer " +
                  (status === submission.status ? statusColors[i] : "")
                }
                onClick={async () => {
                  await postReview({ ...submission, status: status })
                  props.onReview(submission)
                }}
              >
                {toTitleCase(status)}
              </button>
            ))}
          </div>
          <img
            src={`data:image/jpeg;base64,${submission.image}`}
            class="max-h-[40vh]"
          />
        </li>
      ))}
    </ul>
  )
}

export default ReviewList
