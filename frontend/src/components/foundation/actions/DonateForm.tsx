import { FoundrisingAPI } from "../../foundrising/API";
import { FoundrisingPut, FoundrisingTransfer } from "../../foundrising/Transfer";
import  {useState } from "react";
import { useNavigate } from "react-router-dom";
import { FoundationDonate } from "../Transfer";
import { FoundationAPI } from "../API";
import DescriptionList from "../../foundrising/descriptionListCurrent";
import SelectedComponent from "../../foundrising/descriptionListCurrent";
import FoundrisingChoice from "../../foundrising/descriptionListCurrent";

interface FoundationDonateFormProps{
    found_id: number;
   // setDashError:  React.Dispatch<React.SetStateAction<string | null>>
  //  setDashSuccess:  React.Dispatch<React.SetStateAction<string | null>>
}

function FoundationDonateForm({found_id}: FoundationDonateFormProps){
    const [comment, setComment] = useState('');
    const [sumOfMoney, setSumOfMoney] = useState('');
    const [foundrisingID, setFoundrisingID] = useState('');
    //const [commentError, setCommentError] = useState<string>('');
    const [sumOfMoneyError, setSumOfMoneyError] = useState<string>('');
    const [foundrisingIDError, setFoundrisingIDError] = useState<string>('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    const [donateClicked, setDonateClicked] = useState(false);
    

    const isValid = ()=>{
        return (sumOfMoneyError.length === 0 && foundrisingIDError.length === 0);
    }
    const validateSumOfMoney = (value: string) => {
        if (value.length === 0){
            setSumOfMoneyError('Empty sum');
            return;
        }
        if (!/^\d+(\.\d{1,2})?$/.test(value)) {
            setSumOfMoneyError('Invalid required sum format. Example: 1000.00');
        }else if (parseFloat(value) <= 10.00) {
            setSumOfMoneyError('Required sum should be greater than 10.00');
          } 
        else {
            setSumOfMoneyError('');
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

    const handleCommentChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setComment(event.target.value);
        
    };

    const handleSumOfMoneyChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSumOfMoney(event.target.value);
        validateSumOfMoney(event.target.value);
    };
    const handleFoundrisingIDChange = (id: string) => {
        
        setFoundrisingID(id);
        
    };
    const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        window.location.reload();
    }
    
    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        
        
        setError('')
        setDonateClicked(true)
        setSuccess('')
        validateFoundrisingID(foundrisingID)
        validateSumOfMoney(sumOfMoney)
        
        if (!isValid()) {
            // обработка невалидной формы
            setError('Invalid data. Try again');
            return;
        }
        event.preventDefault();
        try {
            const DonateData: FoundationDonate = {
                foundrising_id: foundrisingID,
                comment: comment,
                sum_of_money: parseFloat(sumOfMoney).toFixed(2),
            };
            const token = localStorage.getItem('token');
            if (token === null){
                setError("Foundrising add Form: token is undefined");
                
            }else{
                try{
                    const data =  await FoundationAPI.donate(token, found_id.toString(), JSON.stringify(DonateData))
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
            <h1>Donate to Foundrising</h1>
            <form onSubmit={handleFormSubmit}>
            
            {error.length !== 0 && donateClicked && (
                    <div className="alert alert-error" role="alert">
                    {`Error: ${error}`}
                  </div>
            )}
            {success.length !== 0 && donateClicked && isValid() &&(
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
                <label htmlFor="comment">Comment</label>
                <textarea
                id="comment"
                value={comment}
                onChange={handleCommentChange}
                />
            </div>
            <div className="input-row">
                <label htmlFor="sumOfMoney">Sum of money</label>
                <input
                type="text"
                id="sumOfMoney"
                value={sumOfMoney}
                onChange={handleSumOfMoneyChange}
                />
                <span style={{ color: 'red' }}>{sumOfMoneyError}</span>
            </div>
            <button className="button-login" type="submit">Donate</button>
            <button className="button-login" onClick={handleClose}>Close</button>
            </form>
        </div>
    );
    };

export default FoundationDonateForm;