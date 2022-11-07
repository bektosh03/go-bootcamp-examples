export default function Blog(props) {
  return (
    <p>
      <h3>{props.title} by {props.author}</h3>
      <hr/>
      <h4>{props.body}</h4>
    </p>
  )
}