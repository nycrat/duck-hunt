import { Title } from "@solidjs/meta"
import { A } from "@solidjs/router"

const NotFound = () => {
  return (
    <div class="min-h-screen p-10">
      <Title>Page Not Found | DuckHunt</Title>
      <main class="max-w-[1200px] m-auto">
        <h1>404 page not found</h1>
        <A href="/">click here to return</A>
      </main>
    </div>
  )
}

export default NotFound
