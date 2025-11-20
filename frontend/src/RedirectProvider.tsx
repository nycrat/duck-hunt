import { useNavigate } from "@solidjs/router"
import { createEffect, JSX } from "solid-js"

const RedirectProvider = ({ children }: { children: JSX.Element }) => {
  const navigate = useNavigate()

  createEffect(() => {
    if (!localStorage.getItem("jwtToken")) {
      navigate("/")
    }
  })

  return <>{children}</>
}

export default RedirectProvider
