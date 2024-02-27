
import { LoginAPI } from "./LoginAPI";
import { LoginToken, LoginTransfer } from "./LoginTransfer";
import React, { SyntheticEvent, useEffect, useState } from "react";
import { redirect } from "react-router";
import Logout from "./Logout";
//USER:
//ecane1
//JEqZ7pO8xTb

function LoginForm(){
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('');
    const [loginToken, setLoginToken] = useState<LoginToken | undefined>(undefined);
    const [error, setError] = useState<string | undefined>(undefined);
    for (let i = 0; i < localStorage.length; i ++){
        const key = localStorage.key(i);
        if (key){
            if (key.length >= 4){
                if (key.slice(0, 4) === 'user' || 
                    key.slice(0,10) === 'Foundation' || 
                    key.slice(0, "selected".length) === "selected"){
                    localStorage.removeItem(key)
                }
            }
        }
    }
    localStorage.removeItem('token')
    // Сохранение состояния    

    useEffect(() => {
        const item = window.localStorage.getItem('login')
        if (item && item.length > 0){
            setLogin(item);
        }
        }, []);
    useEffect(() => {
        window.localStorage.setItem('login', login);
        }, [login]);

    useEffect(() => {
        const item = window.localStorage.getItem('password')
        if (item && item.length > 0){
            setPassword(item);
        }
        }, []);
    useEffect(() => {
        window.localStorage.setItem('password', password);
        }, [password]);
    
        
    useEffect(() => {
        const item = window.localStorage.getItem('role')
        if (item && item.length > 0){
            setRole(item);
        }
        }, []);
    useEffect(() => {
        window.localStorage.setItem('role', role);
        }, [role]);
    
    
    const handleLoginChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setLogin(event.target.value);
    };
    
    // Сохранение состояния
    const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
    };

    const handleRoleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setRole(event.target.value);
    };

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        try {
        const data = await LoginAPI.login(login, password, role);
        //const tokenData = new LoginToken(data);
        //setLoginToken();
        console.log("TOKEN IS: ", data.token);
        if (data !== undefined){
            localStorage.setItem('token', data.token);
            if (role === "Foundation"){
                window.location.href = `/foundation_dash/${login}`;
                
                //return;    
            }else if (role === "User"){
                window.location.href = `/user_dash/${login}`;    
                
            }
        }
        }catch (e){
            if (e instanceof Error){
                setError(e.message);
            }
        }
    };

    return (
        <div className="form-container">
            <h1>
                Login
            </h1>
            <form onSubmit={handleSubmit}>
            
            {error &&( 
                    <div className="row">
                        <div className="card large error">
                            <section>
                                <p>
                                    <span className="icon-alert inverse "></span>
                                    {error}
                                </p>
                            </section>
                        </div>
                    </div>
                )}
            
            <div className="input-row">
                <label htmlFor="login">Login:</label>
                <input type="text" id="login" value={login} onChange={handleLoginChange} />
            </div>
            <div className="input-row">
                <label htmlFor="password">Password:</label>
                <input type="password" id="password" value={password} onChange={handlePasswordChange} />
            </div>
            <div className="input-row">
                <label htmlFor="role">Role:</label>
                <select id="role" value={role} onChange={handleRoleChange}>
                <option value="User">User</option>
                <option value="Foundation">Foundation</option>
                </select>
            </div>
            <div className="input-row">
            <button type="submit" className="button-login" >Log in</button>
            </div>
            </form>
        </div>

    );
}
export default LoginForm;