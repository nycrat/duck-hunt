import { Title } from "@solidjs/meta"
import { createSignal } from "solid-js"

const Welcome = () => {
  const [code, setCode] = createSignal("")

  return (
    <main class="h-screen flex justify-center items-center text-center">
      <Title>Welcome | DuckHunt</Title>
      <form
        class="flex flex-col gap-1 items-center"
        onSubmit={(e) => {
          e.preventDefault()
          alert(code())
        }}
      >
        <h1>DuckHunt</h1>
        <input
          type="text"
          value={code()}
          onInput={(e) => {
            setCode(e.target.value)
          }}
          placeholder="Enter your event code"
          class="text-center"
        />
        <button
          type="submit"
          disabled={code() === ""}
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
