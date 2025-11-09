import { Title } from "@solidjs/meta"
import { useNavigate } from "@solidjs/router"
import { createEffect, createResource, createSignal } from "solid-js"

const fetchSessionId = async (): Promise<number | null> => {
  const response = await fetch("http://localhost:8000/session", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("jwtToken")}`,
    },
  })
  if (response.status !== 200) {
    return null
  }
  return parseInt(await response.text())
}

const Welcome = () => {
  const [eventCode, setEventCode] = createSignal("")
  const [passCode, setPassCode] = createSignal("")
  const navigate = useNavigate()

  const [id] = createResource(fetchSessionId)

  createEffect(() => {
    if (id()) {
      navigate("/leaderboard")
    }
  })

  return (
    <main class="h-screen flex justify-center items-center text-center">
      <Title>Welcome | DuckHunt</Title>
      <form
        class="flex flex-col gap-1 items-center"
        onSubmit={async (e) => {
          e.preventDefault()

          const res = await fetch("http://localhost:8000/auth", {
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

          if (eventCode() === "test") {
            navigate("/leaderboard")
          } else {
            alert(`Event (${eventCode()}) not found.`)
          }
        }}
      >
        <h1>DuckHunt</h1>
        {/* <label for="code">Event code:</label> */}
        <input
          type="text"
          name="code"
          value={eventCode()}
          onInput={(e) => {
            setEventCode(e.target.value)
          }}
          placeholder="Enter event code"
          class="text-center"
        />
        {/* <label for="id">ID: </label> */}
        <input
          type="password"
          name="id"
          value={passCode()}
          onInput={(e) => {
            setPassCode(e.target.value)
          }}
          placeholder="Enter your ID"
          class="text-center"
        />
        <button
          type="submit"
          disabled={eventCode() === "" || passCode() === ""}
          class="disabled:text-gray-500 disabled:outline-gray-500 rounded outline-black outline-1
                 w-min px-2 not-disabled:hover:bg-gray-300 not-disabled:cursor-pointer"
        >
          Join
        </button>
      </form>
    </main>
  )
}

export default Welcome
