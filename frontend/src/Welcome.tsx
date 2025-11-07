import { Title } from "@solidjs/meta"
import { useNavigate } from "@solidjs/router"
import { createSignal } from "solid-js"

const Welcome = () => {
  const [code, setCode] = createSignal("")
  const [id, setId] = createSignal("")
  const navigate = useNavigate()

  return (
    <main class="h-screen flex justify-center items-center text-center">
      <Title>Welcome | DuckHunt</Title>
      <form
        class="flex flex-col gap-1 items-center"
        onSubmit={(e) => {
          e.preventDefault()
          if (code() === "test") {
            navigate("/leaderboard")
          } else {
            alert(`Event (${code()}) not found.`)
          }
        }}
      >
        <h1>DuckHunt</h1>
        {/* <label for="code">Event code:</label> */}
        <input
          type="text"
          name="code"
          value={code()}
          onInput={(e) => {
            setCode(e.target.value)
          }}
          placeholder="Enter event code"
          class="text-center"
        />
        {/* <label for="id">ID: </label> */}
        <input
          type="password"
          name="id"
          value={id()}
          onInput={(e) => {
            setId(e.target.value)
          }}
          placeholder="Enter your ID"
          class="text-center"
        />
        <button
          type="submit"
          disabled={code() === "" || id() === ""}
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
