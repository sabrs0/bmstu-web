import { FoundrisingAPI } from "../../foundrising/API";
import { FoundrisingPut, FoundrisingTransfer } from "../../foundrising/Transfer";
import  {useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import FoundrisingChoice from "../../foundrising/descriptionListCurrent";



function FoundrisingUpdateForm(){
    const [description, setDescription] = useState('');
    const [requiredSum, setRequiredSum] = useState('');
    const [foundrisingID, setFoundrisingID] = useState('');
    const [descriptionError, setDescriptionError] = useState<string>('');
    const [requiredSumError, setRequiredSumError] = useState<string>('');
    const [foundrisingIDError, setFoundrisingIDError] = useState<string>('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');



    useEffect(() => {
        const item = window.localStorage.getItem('FoundationUpdateFingDescription')
        if (item && item.length > 0){
            setDescription(item);
        }
        }, []);
    useEffect(() => {
        window.localStorage.setItem('FoundationUpdateFingDescription', description as string);
        }, [description]);

    useEffect(() => {
        const item = window.localStorage.getItem('FoundationUpdateFingRequiredSum')
        if (item && item.length > 0){
            setRequiredSum(item);
        }
        }, []);
    useEffect(() => {
        window.localStorage.setItem('FoundationUpdateFingRequiredSum', requiredSum as string);
        }, [requiredSum]);


        
    useEffect(() => {
        const item = window.localStorage.getItem('FoundationUpdateFingFoundrisingID')
        if (item && item.length > 0){
            setFoundrisingID(item);
        }
        }, []);
    useEffect(() => {
        window.localStorage.setItem('FoundationUpdateFingFoundrisingID', foundrisingID as string);
        }, [foundrisingID]);

        
    const validateDescription = (value: string) => {
        if (value.length < 20) {
        setDescriptionError('Description must be at least 20 characters long');
        } else {
        setDescriptionError('');
        }
    };
    const isValid = ()=>{
        return (descriptionError.length === 0 && requiredSumError.length === 0 && foundrisingIDError.length === 0);
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

    const handleDescriptionChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setDescription(event.target.value);
        validateDescription(event.target.value);
    };

    const handleRequiredSumChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setRequiredSum(event.target.value);
        validateRequiredSum(event.target.value);
    };
    const handleFoundrisingIDChange = (id: string) => {
        
        setFoundrisingID(id);
        
    };
    const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        window.localStorage.setItem('showUpdateFingForm', '0')
        window.localStorage.setItem('showDash', '1')
        window.location.reload();
    }
    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError('')
        setSuccess('')
        validateDescription(description)
        validateFoundrisingID(foundrisingID)
        validateRequiredSum(requiredSum)
       

        if (!isValid()) {
            // обработка невалидной формы
            setError('Invalid data. Try again');
            return;
        }
        try {
            const FoundrisingData: FoundrisingPut = {
                description :    description,
                required_sum: parseFloat(requiredSum).toFixed(2),
            };
            const token = localStorage.getItem('token');
            if (token === null){
                setError("Foundrising add Form: token is undefined");
                
            }else{
                try{
                    const data =  await FoundrisingAPI.update(token, foundrisingID,JSON.stringify(FoundrisingData));
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
            <h1>Update Foundrising</h1>
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
            <div className="input-row">
                <label htmlFor="description">Description</label>
                <textarea
                id="description"
                value={description}
                onChange={handleDescriptionChange}
                />
                <span style={{ color: 'red' }}>{descriptionError}</span>
            </div>
            <div className="input-row">
                <label htmlFor="requiredSum">Required Sum</label>
                <input
                type="text"
                id="requiredSum"
                value={requiredSum}
                onChange={handleRequiredSumChange}
                />
                <span style={{ color: 'red' }}>{requiredSumError}</span>
            </div>
            <button className="button-login" type="submit">Update</button>
            <button className="button-login" onClick={handleClose}>Close</button>
            
            </form>
        </div>
    );
    };

export default FoundrisingUpdateForm;