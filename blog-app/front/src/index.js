import React from 'react';
import ReactDOM from 'react-dom/client';
import reportWebVitals from './reportWebVitals';
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import RegistrationPage from "./pages/registration/registration.page";
import ErrorPage from "./pages/error/error.page";
import WelcomePage from "./pages/welcome/welcome.page";
import BlogsPage from "./pages/blogs/blogs.page";

const router = createBrowserRouter([
  {
    path: "/",
    element: <RegistrationPage/>,
    errorElement: <ErrorPage/>,
  },
  {
    path: "/welcome",
    element: <WelcomePage/>
  },
  {
    path: "/blogs",
    element: <BlogsPage/>
  }
])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
