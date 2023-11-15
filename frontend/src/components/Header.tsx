import { FC, useState } from "react";
import Navbar from "./Navbar";

const Header:FC = () => {

  const [showNav, setShowNav] = useState(false);

  return (
    <header className="bg-gray-300">
      <div className="flex justify-between items-center p-4">
        <h1 className=" font-medium text-4xl">egonote</h1>
        <i
          className="las la-bars text-3xl"
          onClick={() => setShowNav(!showNav)}
        ></i>
      </div>
      {showNav && <Navbar/>}
    </header>
  );
}

export default Header;
