
import  {useState } from "react";
import { FoundationAPI } from "../API";

interface FoundationDeleteProps{
    found_id: number
}

function FoundationDeleteForm({found_id}: FoundationDeleteProps){
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    
    
    
    
    const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        window.location.reload();
    }
    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError('')
        setSuccess('')
        

        try {
            const token = localStorage.getItem('token');
            if (token === null){
                setError("Foundation Delete Form: token is undefined");
                
            }else{
                try{
                    const data =  await FoundationAPI.delete(token, found_id.toString());
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
        <div>
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
            <button type="submit">Delete Foundation</button>
            <button onClick={handleClose}>Close</button>
            
            </form>
        </div>
    );
    };

export default FoundationDeleteForm;