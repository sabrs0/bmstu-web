import { FoundrisingAPI } from "../../foundrising/API";
import { FoundrisingPut, FoundrisingTransfer } from "../../foundrising/Transfer";
import  {useState } from "react";
import FoundrisingChoice from "../../foundrising/descriptionList";

interface FoundrisingUpdateFormProps{
    found_id: number;
}

function FoundrisingDeleteForm(){
    const [foundrisingID, setFoundrisingID] = useState('');
    const [foundrisingIDError, setFoundrisingIDError] = useState<string>('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    
    
    const validateFoundrisingID = (value: string)  => {
        const maxValue = Number.MAX_SAFE_INTEGER; // Максимальное значение для 64-битного целого числа
      
        if (!/^\d+$/.test(value)) {
           setFoundrisingIDError('FundraisingID должен быть целым числом');
        } else {
          const numericValue = parseInt(value, 10);
          if (numericValue < 0) {
            setFoundrisingIDError('FundraisingID должен быть неотрицательным числом');
          } else if (numericValue > maxValue) {
            setFoundrisingIDError(`FundraisingID не должен превышать ${maxValue}`)
          }
        }
    };
    const isValid = ()=>{
        return (foundrisingIDError.length === 0);
    }
    const handleFoundrisingIDChange = (id: string) => {
        
        setFoundrisingID(id);
        
    };
    const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        window.location.reload();
    }
    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError('')
        setSuccess('')
        validateFoundrisingID(foundrisingID)
        

        if (!isValid()) {
            // обработка невалидной формы
            setError('Invalid data. Try again');
            return;
        }
        try {
            const token = localStorage.getItem('token');
            if (token === null){
                setError("Foundrising add Form: token is undefined");
                
            }else{
                try{
                    const data =  await FoundrisingAPI.delete(token, foundrisingID);
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
            <h1>Delete Foundrising</h1>
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
                <label htmlFor="foundrising">Foundrising</label>
                <FoundrisingChoice handleSetID={handleFoundrisingIDChange} />
                {/*<span style={{ color: 'red' }}>{foundrisingIDError}</span>*/}
            </div>
            <button className="button-login" type="submit">Delete</button>
            <button className="button-login" onClick={handleClose}>Close</button>
            
            </form>
        </div>
    );
    };

export default FoundrisingDeleteForm;