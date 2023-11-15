import { FC } from "react";
import { fetchNote } from "../api/notes";
import { Login } from "../api/auth";

const handleClick = async () => {
    let note = await fetchNote(2)
    if(typeof note === undefined){
        console.log("error")
    } else{
        console.log(note)
    }
    
}

const handleClickLogin = async () => {
    Login("username", "password")
}

const Index:FC = () => {

    return (
        <div>
            <h1>This is the index</h1>
            <button className="bg-red-500 p-4" onClick={handleClick}>CLICK THIS</button>
            <button className="bg-green-500 p-4" onClick={handleClickLogin}>Login</button>
            
        </div>
    )
}

export default Index