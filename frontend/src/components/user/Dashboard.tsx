import React,{ Fragment, useEffect, useState } from "react";
import { useParams } from 'react-router-dom';
import FoundationReplenishForm from "./actions/ReplenishForm";
import { UserDonate, UserPut, UserReplenish, UserTransferExt } from "../user/Transfer";
import { UserAPI } from "../user/API";
import UserDonateToFoundForm from "./actions/DonateFoundForm";
import UserDonateToFrisingForm from "./actions/DonateFingForm";
import UserDeleteForm from "./actions/UserDelete";
import UserReplenishForm from "./actions/ReplenishForm";
import UserUpdateForm from "./actions/UserUpdateForm";
import FoundationListForm from "./actions/found_info/FoundationsComponent";
import FoundrisingListForm from "./actions/frising_info/FoundrisingsComponent";
import { IconButton } from "@mui/material";
import UpgradeIcon from "@mui/icons-material/Upgrade"
import PaidIcon from '@mui/icons-material/Paid';
import DeleteIcon from '@mui/icons-material/Delete';
import LogoutIcon from '@mui/icons-material/Logout';

//props: any
function UserDashboard(){
    const [loading, setLoading] = useState(false);
    const [user, setUser] = useState<UserTransferExt | null>(null);
    const [error, setError] = useState<string | null>(null);

    const[showFoundationsForm, setShowFoundationsForm] = useState(false)
    const[showFoundrisingsForm, setShowFoundrisingsForm] = useState(false)
    const[showUpdateForm, setShowUpdateForm] = useState(false);
    const[showDeleteForm, setShowDeleteForm] = useState(false);
    const[showDonateFoundForm, setShowDonateFoundForm] = useState(false);
    const[showDonateFrisingForm, setShowDonateFrisingForm] = useState(false);
    const[showReplenishForm, setShowReplenishForm] = useState(false);
    const [userPut, setUserPut] = useState<UserPut | undefined>(undefined)
    const[showDash, setShowDash] = useState(true);
    let setsArr: React.Dispatch<React.SetStateAction<boolean>>[] = [
        setShowUpdateForm,setShowDeleteForm,
        setShowDonateFoundForm, setShowDonateFrisingForm, 
        setShowReplenishForm, setShowFoundationsForm,
        setShowFoundrisingsForm, setShowDash ]
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
    
    const handleUpd = () =>{
        falseAll();
        setShowUpdateForm(true);
        setUserPut(new UserPut(user))
    }
    const handleDel = () =>{
        falseAll();
        setShowDeleteForm(true);
    }
    const handleDonFound = () =>{
        falseAll();
        setShowDonateFoundForm(true);
    }
    const handleDonFrising = () =>{
        falseAll();
        setShowDonateFrisingForm(true);
    }
    const handleRep = () =>{
        falseAll();
        setShowReplenishForm(true);
    }
    const handleListFounds = () =>{
        falseAll();
        setShowFoundationsForm(true);
    }
    const handleListFrisings = () =>{
        falseAll();
        setShowFoundrisingsForm(true);
    }

    const params = useParams();
    const login = String(params.login);
    const token = localStorage.getItem('token');
    useEffect(() => {
        
        if (token === null){
            setError("Invalid token. Try again");
        }else{
            setLoading(true);
            UserAPI
            .findExtLogin(login, token)
            .then((data) => {
                if (data === null){
                    setError("Data is null. Try again");
                }else{
                    setUser(data);
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

            {error && ( 
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
            
            {!error && !loading && user &&(
                 <div className="dash">
                 {showDash && (<div className="dash">
                    <div className="panel-container">
                        <div className="user-info">
                            <strong>
                            <p>Login: {user.login}</p>
                            </strong>
                            <p>Balance: {user.balance}</p>
                            <p>Charity sum: {user.charity_sum}</p>
                            <p>Closed Foundrising Amount: {user.closed_fing_amount}</p>
                            
                        </div>
                        <div className="icon-operations">    
                            <button className="button-login" onClick={handleRep}> <PaidIcon/></button>
                            <button className="button-login" onClick={handleUpd}><UpgradeIcon/></button>
                            <button className="button-login" onClick={handleDel}><DeleteIcon/></button>
        
                            <button className="button-login" onClick={handleExit}><LogoutIcon/></button>
                        </div>
                </div>
                <div className="operations">
                    <button className="button-login" onClick={handleDonFound}>Donate to Foundation</button>
                    <button className="button-login" onClick={handleDonFrising}>Donate to Foundrising</button>
                    <button className="button-login"  onClick={handleListFounds}>List Foundations</button>
                    <button className="button-login" onClick={handleListFrisings}>List Foundrisings</button>

                </div>
                </div>)}

                {showDonateFoundForm && (<UserDonateToFoundForm user_id={user.id === undefined ? 0 :user.id} initialDonateForm={new UserDonate} /> ) }
                {showDonateFrisingForm && (<UserDonateToFrisingForm user_id={user.id === undefined ? 0 :user.id} initialDonateForm={new UserDonate}/> ) }
                {showReplenishForm && (<UserReplenishForm user_id={user.id === undefined ? 0 :user.id} /> ) }
                {showUpdateForm && (<UserUpdateForm 
                                user_id={user.id === undefined ? 0 :user.id} 
                                user_put_info={user === undefined ? new UserPut(undefined): user}/> ) }
                {showDeleteForm && (<UserDeleteForm 
                    user_id={user.id === undefined ? 0 :user.id}/> ) }
                {showFoundationsForm && (<FoundationListForm 
                    user_id={user.id === undefined ? 0 :user.id}/> ) }
                {showFoundrisingsForm && (<FoundrisingListForm 
                    user_id={user.id === undefined ? 0 :user.id}/> ) }
            </div>
            )}
        </div>
      );
}
export default UserDashboard;