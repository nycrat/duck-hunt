import { createResource, JSX, Match, Switch } from "solid-js"
import { getServerURL } from "../utils"
import { fetchWithMiddleware } from "../api"
import RedirectProvider from "../RedirectProvider"

const authenticate = async (): Promise<boolean> => {
  const response = await fetchWithMiddleware(`${getServerURL()}/auth/admin`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  return response.status === 200
}

/**
 * NOTE: AdminRoute is not made to securely hide the admin interface since
 * everything is client-side anyways. Just here so that users don't see the
 * admin interface, need to authenticate for admin actions anyways on server
 */
const AdminRoute = ({ children }: { children: JSX.Element }) => {
  const [hasAdminAccess] = createResource(authenticate)

  return (
    <div class="bg-amber-50">
      <RedirectProvider>
        <Switch>
          <Match when={hasAdminAccess.loading}>loading...</Match>
          <Match when={!hasAdminAccess()}>
            You have no access to this page
          </Match>
          <Match when={hasAdminAccess()}>{children}</Match>
        </Switch>
      </RedirectProvider>
    </div>
  )
}

export default AdminRoute
