import Blog from "../../components/blog/blog.component";
import {useState} from "react";

async function fetchBlogs() {
  const response = await fetch("http://localhost:8080/blogs", {
    method: "GET",
  })

  return await response.json()
}

export default function BlogsPage() {
  const [blogs, setBlogs] = useState([])
  console.log(blogs)
  if (blogs.length < 1) {
    fetchBlogs().then(fetchedBlogs => setBlogs(fetchedBlogs))
  }

  return blogs.map(blog => <Blog key={blogs.title} title={blog.title} author={blog.author}
                                 body={blog.body}/>)
}
