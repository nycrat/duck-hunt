import { Title } from "@solidjs/meta"
import { useNavigate } from "@solidjs/router"
import { createEffect, createResource, createSignal } from "solid-js"
import { fetchWithMiddleware } from "./api"
import RedirectProvider from "./RedirectProvider"
import logo from "./assets/favicon.svg"

const fetchSessionId = async (): Promise<number | null> => {
  const response = await fetchWithMiddleware(`/auth/session`, {
    method: "POST",
  })
  if (response.status !== 200) {
    return null
  }
  return parseInt(await response.text())
}

const Welcome = () => {
  const [passCode, setPassCode] = createSignal("")
  const navigate = useNavigate()

  const [id] = createResource(fetchSessionId)

  createEffect(() => {
    if (id()) {
      navigate("/leaderboard")
    }
  })

  return (
    <RedirectProvider>
      <main class="h-dvh flex justify-center items-center text-center">
        <Title>Welcome | DuckHunt</Title>
        <form
          class="flex flex-col gap-1 items-center"
          onSubmit={async (e) => {
            e.preventDefault()

            const res = await fetchWithMiddleware(`/auth/login`, {
              method: "POST",
              headers: {
                Authorization: `Basic ${passCode()}`,
              },
            })

            if (res.status !== 200) {
              alert("Issue with authorization request")
              return
            }

            localStorage.setItem("jwtToken", await res.text())

            navigate("/leaderboard")
          }}
        >
          <img src={logo} alt="duck hunt logo" width={150} />
          <div class="h-2" />
          <h1>DuckHunt</h1>
          <input
            type="password"
            name="id"
            value={passCode()}
            onInput={(e) => {
              setPassCode(e.target.value)
            }}
            placeholder="Enter your passcode"
            class="text-center"
          />
          <button
            type="submit"
            disabled={passCode() === ""}
            class="disabled:text-gray-500 disabled:outline-gray-500 rounded outline-black outline-1
                 w-min px-2 not-disabled:hover:bg-gray-300 not-disabled:cursor-pointer"
          >
            Join
          </button>
          <div class="h-16" />
        </form>
      </main>
    </RedirectProvider>
  )
}

export default Welcome
