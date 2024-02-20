import { FoundrisingAPI } from "../../foundrising/API";
import { FoundrisingPut, FoundrisingTransfer } from "../../foundrising/Transfer";
import  {useState } from "react";
import { useNavigate } from "react-router-dom";
import { FoundationPut } from "../Transfer";
import { FoundationAPI } from "../API";

interface FoundationUpdateProps{
    foundation_put_info: FoundationPut
    found_id: number
}
//вставлять текущие значения в поля при update тут и , user
function FoundationUpdateForm({found_id, foundation_put_info}: FoundationUpdateProps){
    const [name, setName] = useState(foundation_put_info.name);
    const [login, setLogin] = useState(foundation_put_info.login);
    const [password, setPassword] = useState(foundation_put_info.password);
    const [country, setCountry] = useState(foundation_put_info.country);
    const [nameError, setNameError] = useState<string>('');
    const [loginError, setLoginError] = useState<string>('');
    const [passwordError, setPasswordError] = useState<string>('');
    const [countryError, setCountryError] = useState<string>('');
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
        return (nameError.length === 0 && loginError.length === 0 && passwordError.length === 0 && countryError.length === 0);
    }
    const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setName(event.target.value);
        validateString(event.target.value, setNameError);
    };

    const handleLoginChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setLogin(event.target.value);
        validateString(event.target.value, setLoginError);
    };
    const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
        validateString(event.target.value, setPasswordError);
    };
    const handleCountryChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setCountry(event.target.value);
        validateString(event.target.value, setCountryError);
    };
    const handleClose = () => {
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
        validateString(name, setNameError)
        validateString(login, setLoginError)
        validateString(password, setPasswordError)
        validateString(country, setCountryError)


        if (!isValid()) {
            // обработка невалидной формы
            setError('Invalid data. Try again');
            return;
        }
        try {
            const FoundationData: FoundationPut = {
                name: name,
                login: login,
                password: password,
                country: country,
            };
            const token = localStorage.getItem('token');
            if (token === null){
                setError("Foundation put Form: token is undefined");
                
            }else{
                try{
                    const data =  await FoundationAPI.update(token, found_id.toString(), JSON.stringify(FoundationData));
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
            <h1>Update Foundation</h1>
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
                <label htmlFor="name">Name</label>
                <input
                type="text"
                id="name"
                value={name}
                onChange={handleNameChange}
                />
                <span style={{ color: 'red' }}>{nameError}</span>
            </div>
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
            <div className="input-row">
                <label htmlFor="country">Country</label>
                <input
                type="text"
                id="country"
                value={country}
                onChange={handleCountryChange}
                />
                <span style={{ color: 'red' }}>{countryError}</span>
            </div>
            <button className="button-login" type="submit">Update</button>
            <button className="button-login" onClick={handleClose}>Close</button>
            
            </form>
        </div>
    );
};

export default FoundationUpdateForm;