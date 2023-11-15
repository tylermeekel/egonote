import { FC } from "react"
import { Link } from "react-router-dom"

const Navbar:FC = () => {
    return(
        <nav className="p-4 fixed w-full bg-gray-300 border-t border-gray-300 z-50 -mt-1">
            <ul className="text-right text-2xl last flex flex-col gap-4">
                <li className=""><Link to={'notebook'}>Notebook</Link></li>
                <li className="">Settings</li>
            </ul>
        </nav>
    )
}

export default Navbar