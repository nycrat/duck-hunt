import { JSX } from "solid-js"
import logo from "../assets/favicon.svg"

const Footer = ({ children }: { children: JSX.Element }) => {
  return (
    <div class="flex justify-between">
      <div class="flex flex-col justify-end">{children}</div>
      <div>
        <img src={logo} alt="duck hunt logo" width={100} />
      </div>
    </div>
  )
}

export default Footer
