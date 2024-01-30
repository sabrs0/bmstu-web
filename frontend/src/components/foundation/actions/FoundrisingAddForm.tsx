import { FoundrisingAPI } from "../../foundrising/API";
import { FoundrisingAdd, FoundrisingTransfer } from "../../foundrising/Transfer";
import  {useState } from "react";
import { useNavigate } from "react-router-dom";

interface FoundrisingAddFormProps{
    found_id: number;
   // setDashError:  React.Dispatch<React.SetStateAction<string | null>>
  //  setDashSuccess:  React.Dispatch<React.SetStateAction<string | null>>
}

function FoundrisingAddForm({found_id}: FoundrisingAddFormProps){
    const [description, setDescription] = useState('');
    const [requiredSum, setRequiredSum] = useState('');
    const [descriptionError, setDescriptionError] = useState<string>('');
    const [requiredSumError, setRequiredSumError] = useState<string>('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    const validateDescription = (value: string) => {
        if (value.length < 20) {
        setDescriptionError('Description must be at least 20 characters long');
        } else {
        setDescriptionError('');
        }
    };
    const isValid = ()=>{
        return (descriptionError.length === 0 && requiredSumError.length === 0);
    }
    const validateRequiredSum = (value: string) => {
        if (!/^\d+(\.\d{1,2})?$/.test(value)) {
        setRequiredSumError('Invalid required sum format. Example: 1000.00');
        }else if (parseFloat(value) <= 10.00) {
            setRequiredSumError('Required sum should be greater than 10.00');
          } 
        else {
        setRequiredSumError('');
        }
    };

    const handleDescriptionChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setDescription(event.target.value);
        validateDescription(event.target.value);
    };

    const handleRequiredSumChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setRequiredSum(event.target.value);
        validateRequiredSum(event.target.value);
    };
    const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        window.location.reload();
    }
    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError('')
        setSuccess('')
        validateDescription(description);
        validateRequiredSum(description);
        

        if (!isValid()) {
            // обработка невалидной формы
            setError('Try again');
            return;
        }
        try {
            const FoundrisingData: FoundrisingAdd = {
                found_id: found_id,
                description :    description,
                required_sum: parseFloat(requiredSum).toFixed(2),
                creation_date: new Date().toISOString()
            };
            const token = localStorage.getItem('token');
            if (token === null){
                setError("Foundrising add Form: token is undefined");
                
            }else{
                try{
                    const data =  await FoundrisingAPI.add(token, JSON.stringify(FoundrisingData));
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
        <div>
            <form onSubmit={handleFormSubmit}>
            
            {error.length != 0 && (
                    <div className="alert alert-error" role="alert">
                    {`Error: ${error}`}
                  </div>
            )}
            {success.length != 0 && (
                    <div className="alert alert-success" role="alert">
                    Успешно
                  </div>
            )}
            <div>
                <label htmlFor="description">Description</label>
                <textarea
                id="description"
                value={description}
                onChange={handleDescriptionChange}
                />
                <span style={{ color: 'red' }}>{descriptionError}</span>
            </div>
            <div>
                <label htmlFor="requiredSum">Required Sum</label>
                <input
                type="text"
                id="requiredSum"
                value={requiredSum}
                onChange={handleRequiredSumChange}
                />
                <span style={{ color: 'red' }}>{requiredSumError}</span>
            </div>
            <button type="submit">Create Foundrising</button>
            <button onClick={handleClose}>Close</button>
            
            </form>
        </div>
    );
    };

export default FoundrisingAddForm;
/*{error &&( 
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
            )}*/