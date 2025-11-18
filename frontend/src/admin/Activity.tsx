import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"

const ActivityInfo = () => {
  return (
    <main class="h-dvh p-10 flex flex-col gap-1">
      <Title>Activity | DuckHunt Admin</Title>
      <h1>Activity Dashboard</h1>
      admin page: activities
      <div class="grow" />
      <A href="/admin/activities">Go back</A>
    </main>
  )
}

export default ActivityInfo
