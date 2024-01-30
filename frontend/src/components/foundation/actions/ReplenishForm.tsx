import { FoundrisingAPI } from "../../foundrising/API";
import { FoundrisingPut, FoundrisingTransfer } from "../../foundrising/Transfer";
import  {useState } from "react";
import { useNavigate } from "react-router-dom";
import { FoundationDonate, FoundationReplenish } from "../Transfer";
import { FoundationAPI } from "../API";

interface FoundationReplenishFormProps{
    found_id: number;
   // setDashError:  React.Dispatch<React.SetStateAction<string | null>>
  //  setDashSuccess:  React.Dispatch<React.SetStateAction<string | null>>
}

function FoundationReplenishForm({found_id}: FoundationReplenishFormProps){
    
    const [sumOfMoney, setSumOfMoney] = useState('');
    const [sumOfMoneyError, setSumOfMoneyError] = useState<string>('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    
    const isValid = ()=>{
        return (sumOfMoneyError.length === 0);
    }
    const validateSumOfMoney = (value: string) => {
        if (!/^\d+(\.\d{1,2})?$/.test(value)) {
            setSumOfMoneyError('Invalid required sum format. Example: 1000.00');
        }else if (parseFloat(value) <= 10.00) {
            setSumOfMoneyError('Required sum should be greater than 10.00');
          } 
        else {
            setSumOfMoneyError('');
        }
    };


    const handleSumOfMoneyChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSumOfMoney(event.target.value);
        validateSumOfMoney(event.target.value);
    };
    const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        window.location.reload();
    }
    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError('')
        setSuccess('')
        validateSumOfMoney(sumOfMoney)
        

        if (!isValid()) {
            // обработка невалидной формы
            setError('Invalid data. Try again');
            return;
        }
        try {
            const DonateData: FoundationReplenish = {
                sum_of_money: parseFloat(sumOfMoney).toFixed(2),
            };
            const token = localStorage.getItem('token');
            if (token === null){
                setError("Foundrising add Form: token is undefined");
                
            }else{
                try{
                    const data =  await FoundationAPI.replenish(token, found_id.toString(), JSON.stringify(DonateData))
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
            <div>
                <label htmlFor="sumOfMoney">Sum of money</label>
                <input
                type="text"
                id="sumOfMoney"
                value={sumOfMoney}
                onChange={handleSumOfMoneyChange}
                />
                <span style={{ color: 'red' }}>{sumOfMoneyError}</span>
            </div>
            <button type="submit">Replenish</button>
            <button onClick={handleClose}>Close</button>
            </form>
        </div>
    );
    };

export default FoundationReplenishForm;