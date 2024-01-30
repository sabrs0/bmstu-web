import React from "react";
import { NavLink } from "react-router-dom";
interface LogoutProps{
    setLoginToken: (value: any) => void;
}
function Logout({setLoginToken}: LogoutProps){
    const handleLogoout = ()=>{
        localStorage.removeItem('token');
        setLoginToken(undefined);
        window.location.href = '/';
    }
    return (
        <button className='nav-link'  onClick={handleLogoout}>
        </button>
    )
}
export default Logout;