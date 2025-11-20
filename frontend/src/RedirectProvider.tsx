import { useNavigate } from "@solidjs/router"
import { createEffect, JSX } from "solid-js"
import { getSessionId } from "./utils"

const RedirectProvider = ({ children }: { children: JSX.Element }) => {
  const navigate = useNavigate()

  createEffect(() => {
    if (!getSessionId()) {
      navigate("/")
      localStorage.removeItem("jwtToken")
    }
  })

  return <>{children}</>
}

export default RedirectProvider
