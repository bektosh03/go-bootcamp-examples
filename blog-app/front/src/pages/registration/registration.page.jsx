import {useNavigate} from "react-router-dom";
import {useState} from "react";

export default function RegistrationPage() {
  const navigate = useNavigate()
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const handleSubmit = async (event) => {
    event.preventDefault()
    const res = await fetch("http://localhost:8080/sign-in", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        "username": username,
        "password": password,
      })
    })

    const body = await res.json()
    console.log(body)
    console.log(res.ok)
    if (res.ok && body.success) {
      console.log("redirecting...")
      navigate(`/welcome?username=${username}`)
    }
  }

  return (
    <form>
      <p>
        <label htmlFor="username">Username </label>
        <input type="text" name="username" id="username" onChange={(event) => {
          setUsername(event.target.value)
        }}/>
      </p>
      <p>
        <label htmlFor="password">Password </label>
        <input type="password" name="password" id="password" onChange={(event) => {
          setPassword(event.target.value)
        }}/>
      </p>
      <button type="submit" onClick={handleSubmit}>Submit</button>
    </form>
  )
}