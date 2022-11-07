import {useRouteError} from "react-router-dom";

export default function ErrorPage() {
  const error = useRouteError()
  console.log(error)

  return (
    <div>
      <h2>Something went wrong:</h2>
      <h4>{error}</h4>
    </div>
  )
}