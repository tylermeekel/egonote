import { Outlet } from "react-router-dom"
import Header from "../components/Header"
import { FC } from "react"

const Root:FC = () => {

  return (
    <main className=" bg-white dark:bg-gray-700 w-screen min-h-screen">
      <Header/>
      <Outlet />
    </main>
  )
}

export default Root
