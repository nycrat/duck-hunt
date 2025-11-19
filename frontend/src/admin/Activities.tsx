import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"
import AdminRoute from "./AdminRoute"

const ActivitiesDashboard = () => {
  return (
    <AdminRoute>
      <main class="h-dvh p-10 flex flex-col gap-1">
        <Title>Activities | DuckHunt Admin</Title>
        <h1>Activities Dashboard</h1>
        admin page: activities
        <div class="grow" />
        <A href="/admin/participants">Go to participants</A>
      </main>
    </AdminRoute>
  )
}

export default ActivitiesDashboard
