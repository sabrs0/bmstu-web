
import  {useState } from "react";
import { UserAPI } from "../API";

interface UserDeleteProps{
    user_id: number
}

function UserDeleteForm({user_id}: UserDeleteProps){
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    
    
    
    
    const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        window.localStorage.setItem('showDeleteForm', '0')
        window.localStorage.setItem('showDash', '1')
        window.location.reload();
    }
    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError('')
        setSuccess('')
        

        try {
            const token = localStorage.getItem('token');
            if (token === null){
                setError("User Delete Form: token is undefined");
                
            }else{
                try{
                    const data =  await UserAPI.delete(token, user_id.toString());
                    setSuccess('Успешно')
                    localStorage.removeItem('token');
                    window.location.href = '/login';
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
            <button className='button-login' type="submit">Delete User</button>
            <button className='button-login' onClick={handleClose}>Close</button>
            
            </form>
        </div>
    );
    };

export default UserDeleteForm;