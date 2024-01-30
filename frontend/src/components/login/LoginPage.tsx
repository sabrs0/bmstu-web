
import { LoginAPI } from "./LoginAPI";
import { LoginToken, LoginTransfer } from "./LoginTransfer";
import React, { SyntheticEvent, useState } from "react";
import { redirect } from "react-router";
import Logout from "./Logout";
/*interface LoginFormProps{
    setEntityRole:  React.Dispatch<React.SetStateAction<string>>
   // onSave: (Login: Login) => void;
    //onCancel: () =>void;
}*/
function LoginForm(){
    const [login, setLogin] = useState('gduggen0');
    const [password, setPassword] = useState('hbsvo9r');
    const [role, setRole] = useState('User');
    const [loginToken, setLoginToken] = useState<LoginToken | undefined>(undefined);
    const [error, setError] = useState<string | undefined>(undefined);
    localStorage.removeItem('token')
    const handleLoginChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setLogin(event.target.value);
    };

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
        
        <div>
            <label htmlFor="login">Login:</label>
            <input type="text" id="login" value={login} onChange={handleLoginChange} />
        </div>
        <div>
            <label htmlFor="password">Password:</label>
            <input type="password" id="password" value={password} onChange={handlePasswordChange} />
        </div>
        <div>
            <label htmlFor="role">Role:</label>
            <select id="role" value={role} onChange={handleRoleChange}>
            <option value="User">User</option>
            <option value="Foundation">Foundation</option>
            </select>
        </div>
        <button type="submit" >Log in</button>
        
        </form>
    );
}
export default LoginForm;