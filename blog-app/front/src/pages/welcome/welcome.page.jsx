import {useSearchParams} from "react-router-dom";

export default function WelcomePage() {
  const [params] = useSearchParams()

  return (
    <h1>Welcome {params.get("username")}!</h1>
  )
}