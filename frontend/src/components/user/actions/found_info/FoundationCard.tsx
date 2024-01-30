import { Link } from "react-router-dom";
import { FoundationTransfer } from "../../../foundation/Transfer";
import { useState } from "react";
import UserDonateToFoundForm from "../DonateFoundForm";
import { UserDonate } from "../../Transfer";

interface FoundationCardProps{
    user_id: number
    Found: FoundationTransfer;
}

function FoundationCard({user_id, Found}: FoundationCardProps){
    const  Foundation  = Found;
    const[donate, setDonate] = useState(false);
    const handleDonate = () =>{
        setDonate(true)
        //window.location.reload()
    }
    return (
        <div className="card border-dark mb-3" style={{maxWidth: "350px", border: "2px solid"}}>
            {donate  && ( <UserDonateToFoundForm user_id={user_id} 
                                initialDonateForm={{entity_type:false,
                                    entity_id: Foundation.id ? Foundation.id.toString(): '',
                                    sum_of_money: '',
                                    comment: ''
                                }
                                }/>)
                }
            {!donate && (
            <div className="card-body">
                <section className="section dark">
                    <h5 className="strong">
                        <strong>id: {Foundation.id}</strong>
                    </h5>
                            
                    <p>name : {Foundation.name}</p>
                    <p>Current foundrising amount : {Foundation.cur_foudrising_amount.toLocaleString()}</p>
                    <p>Closed foundrising amount : {Foundation.closed_foundrising_amount.toLocaleString()}</p>
                    <p>Volunteer amount : {Foundation.volunteer_amount.toLocaleString()}</p>
                    <p>Country : {Foundation.country}</p>
                    <button onClick={handleDonate}>Donate</button>
                </section>
            </div>)}
        </div>
            
    );
}
export default FoundationCard;