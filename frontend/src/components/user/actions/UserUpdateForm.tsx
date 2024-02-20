
import  {useState } from "react";
import { useNavigate } from "react-router-dom";
import { UserPut } from "../Transfer";
import { UserAPI } from "../API";

interface UserUpdateProps{
    user_put_info: UserPut
    user_id: number
}
//вставлять текущие значения в поля при update тут и , user
function UserUpdateForm({user_id, user_put_info}: UserUpdateProps){
    const [login, setLogin] = useState(user_put_info.login);
    const [password, setPassword] = useState(user_put_info.password);
    const [loginError, setLoginError] = useState<string>('');
    const [passwordError, setPasswordError] = useState<string>('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    const validateString = (value: string, setError: React.Dispatch<React.SetStateAction<string>>) => {
        if (value.length < 1) {
            setError('Must be at least 1 characters long');
        } else {
            setError('');
        }
    };
    
    const isValid = ()=>{
        return (loginError.length === 0 && passwordError.length === 0 );
    }

    const handleLoginChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setLogin(event.target.value);
        validateString(event.target.value, setLoginError);
    };
    const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
        validateString(event.target.value, setPasswordError);
    };
    const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault()
        if (success !== ''){
            const currentUrl: string = window.location.href;

            const urlParts: string[] = currentUrl.split('/');
    
            urlParts[urlParts.length - 1] = login;
    
            const newUrl: string = urlParts.join('/');
    
            window.location.href = newUrl;
        }else{
            window.location.reload()
        }
    }
    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError('')
        setSuccess('')
        validateString(login, setLoginError)
        validateString(password, setPasswordError)


        if (!isValid()) {
            // обработка невалидной формы
            setError('Invalid data. Try again');
            return;
        }
        try {
            const UserData: UserPut = {
                login: login,
                password: password,
            };
            const token = localStorage.getItem('token');
            if (token === null){
                setError("User put Form: token is undefined");
                
            }else{
                try{
                    const data =  await UserAPI.update(token, user_id.toString(), JSON.stringify(UserData));
                    setSuccess('Успешно')
                }catch (e){
                    if (e instanceof Error){
                        setError(e.message);
                    }
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
            <form onSubmit={handleFormSubmit}>
            
            {error.length !== 0 && (
                    <div className="alert alert-error" role="alert">
                    {`Error: ${error}`}
                  </div>
            )}
            {success.length !== 0 && (
                    <div className="alert alert-success" role="alert">
                    Успешно
                  </div>
            )}
            <div className="input-row">
                <label htmlFor="login">Login</label>
                <input
                type="text"
                id="login"
                value={login}
                onChange={handleLoginChange}
                />
                <span style={{ color: 'red' }}>{loginError}</span>
            </div>
            <div className="input-row">
                <label htmlFor="password">Password</label>
                <input
                type="text"
                id="password"
                value={password}
                onChange={handlePasswordChange}
                />
                <span style={{ color: 'red' }}>{passwordError}</span>
            </div>
            <button className='button-login' type="submit">Update</button>
            <button className='button-login' onClick={handleClose}>Close</button>
            
            </form>
        </div>
    );
};

export default UserUpdateForm;