import React,{ Fragment, useEffect, useState } from "react";
import { FoundationExt, FoundationPut } from "./Transfer";
import { useParams } from 'react-router-dom';
import { FoundationAPI } from "./API";
import FoundrisingAddForm from "./actions/FoundrisingAddForm";
import FoundrisingUpdateForm from "./actions/FoundrisingUpdateForm";
import FoundrisingListForm from "./actions/FoundrisingList";
import FoundrisingDeleteForm from "./actions/FoundrisingDelete";
import FoundationDonateForm from "./actions/DonateForm";
import FoundationReplenishForm from "./actions/ReplenishForm";
import FoundationUpdateForm from "./actions/FoundationUpdateForm";
import FoundationDeleteForm from "./actions/FoundationDelete";
import UpgradeIcon from "@mui/icons-material/Upgrade"
import PaidIcon from '@mui/icons-material/Paid';
import DeleteIcon from '@mui/icons-material/Delete';
import LogoutIcon from '@mui/icons-material/Logout';
//props: any
function FoundationDashboard(){
    const [loading, setLoading] = useState(false);
    const [Foundation, setFoundation] = useState<FoundationExt | null>(null);
    const [error, setError] = useState<string | null>(null);
    //const [success, setSuccess] = useState<string | null>(null);

    const[showAddForm, setShowAddForm] = useState(false);
    const[showUpdateFingForm, setShowUpdateFingForm] = useState(false);
    const[showUpdateForm, setShowUpdateForm] = useState(false);
    const[showDeleteForm, setShowDeleteForm] = useState(false);
    const[showListForm, setShowListForm] = useState(false);
    const[showDonateForm, setShowDonateForm] = useState(false);
    const[showReplenishForm, setShowReplenishForm] = useState(false);
    const[showDeleteFingForm, setShowDeleteFingForm] = useState(false);
    const[showDash, setShowDash] = useState(true);
    const [foundationPut, setFoundationPut] = useState<FoundationPut | undefined>(undefined)
    let setsArr: React.Dispatch<React.SetStateAction<boolean>>[] = [
        setShowAddForm, setShowUpdateFingForm, setShowUpdateForm,setShowDeleteForm,
        setShowListForm, setShowDonateForm, setShowReplenishForm, setShowDeleteFingForm,
        setShowDash
    ]
    function falseAll(){
        for (let i = 0; i < setsArr.length; i ++){
            setsArr[i](false);
        }
    }
    const handleExit = () =>{
        localStorage.removeItem('token');
        //setLoginToken(undefined);
        window.location.href = '/login';
    }
    const handleAdd = () =>{
        falseAll();
        setShowAddForm(true);
    }
    const handleUpdFing = () =>{
        falseAll();
        setShowUpdateFingForm(true);
    }
    const handleUpd = () =>{
        falseAll();
        setShowUpdateForm(true);
        setFoundationPut(new FoundationPut(Foundation))
    }
    const handleDelFing = () =>{
        falseAll();
        setShowDeleteFingForm(true);
    }
    const handleDel = () =>{
        falseAll();
        setShowDeleteForm(true);
    }
    const handleDon = () =>{
        falseAll();
        setShowDonateForm(true);
    }
    const handleRep = () =>{
        falseAll();
        setShowReplenishForm(true);
    }
    const handleList = () =>{
        falseAll();
        setShowListForm(true);
    }

    const params = useParams();
    const login = String(params.login);
    const token = localStorage.getItem('token');
    useEffect(() => {
        
        if (token === null){
            setError("Invalid token. Try again");
        }else{
            setLoading(true);
            FoundationAPI
            .findExtLogin(login, token)
            .then((data) => {
                if (data === null){
                    setError("Data is null. Try again");
                }else{
                    setFoundation(data);
                    setLoading(false);
                }
            })
            .catch((e) => {
                setError(e);
                setLoading(false);
            });
        }
    }, [login, token]);
    return (
        <div>
            {loading && (
                <div className="center-page">
                    <span className="spinner primary"></span>
                    <p>Loading...</p>
                </div>
            )}

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
            
            {!error && !loading && Foundation &&(
                <div className="dash">
                    {showDash && (<div className="dash">
                        <div className="panel-container">
                            <div className="user-info">
                                <strong>
                                <p>Name: {Foundation.name}</p>
                                </strong>
                                <p>Current Foundrising Amount: {Foundation.cur_foudrising_amount}</p>
                                <p>Closed Foundrising Amount: {Foundation.closed_foundrising_amount}</p>
                                <p>Volunteer Amount: {Foundation.volunteer_amount}</p>
                                <p>Country: {Foundation.country}</p>
                                <p>Balance: {Foundation.fund_balance}</p>
                                <p>Income History: {Foundation.income_history}</p>
                                <p>Outcome History: {Foundation.outcome_history}</p>
                            
                            </div>
                            <div>
                                <div className="icon-operations">    
                                <button className="button-login" onClick={handleRep}> <PaidIcon/></button>
                                <button className="button-login" onClick={handleUpd}><UpgradeIcon/></button>
                                <button className="button-login" onClick={handleDel}><DeleteIcon/></button>
            
                                <button className="button-login" onClick={handleExit}><LogoutIcon/></button>
                                </div>
                            </div>
                    </div>
                    <div className="operations">
                        <button className="button-login" onClick={handleList}>My Foundrisings</button>
                        <button className="button-login" onClick={handleAdd}>Add Foundrising</button>
                        <button className="button-login" onClick={handleUpdFing}>Update Foundrising</button>
                        <button className="button-login" onClick={handleDelFing}>Delete Foundrising</button>
                        <button className="button-login" onClick={handleDon}>Donate</button>
                    </div>
                </div>)}
                    
                {showAddForm && (<FoundrisingAddForm 
                    found_id={Foundation.id === undefined ? 0 :Foundation.id}/> ) }
                {showUpdateFingForm && (<FoundrisingUpdateForm 
                    /> ) }
                {showListForm && (<FoundrisingListForm found_id={Foundation.id === undefined ? 0 :Foundation.id} /> ) }
                {showDeleteFingForm && (<FoundrisingDeleteForm  /> ) }
                {showDonateForm && (<FoundationDonateForm found_id={Foundation.id === undefined ? 0 :Foundation.id} /> ) }
                {showReplenishForm && (<FoundationReplenishForm found_id={Foundation.id === undefined ? 0 :Foundation.id} /> ) }
                {showUpdateForm && (<FoundationUpdateForm 
                                found_id={Foundation.id === undefined ? 0 :Foundation.id} 
                                foundation_put_info={foundationPut === undefined ? new FoundationPut(undefined): foundationPut}/> ) }
                {showDeleteForm && (<FoundationDeleteForm 
                    found_id={Foundation.id === undefined ? 0 :Foundation.id}/> ) }
            </div>
            )}
        </div>
      );
}
export default FoundationDashboard;