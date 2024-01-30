import { Link } from "react-router-dom";
import { FoundrisingTransfer } from "../../../foundrising/Transfer";
import { useState } from "react";
import UserDonateToFoundForm from "../DonateFoundForm";
import { UserDonate } from "../../Transfer";
import UserDonateToFrisingForm from "../DonateFingForm";

interface FoundrisingCardProps{
    user_id: number
    Found: FoundrisingTransfer;
}

function FoundrisingCard({user_id, Found}: FoundrisingCardProps){
    const  Foundrising  = Found;
    const[donate, setDonate] = useState(false);
    const handleDonate = () =>{
        setDonate(true)
        //window.location.reload()
    }
    return (
        <div className="card border-dark mb-3" style={{maxWidth: "350px", border: "2px solid"}}>
            {donate  && ( <UserDonateToFrisingForm user_id={user_id} 
                                initialDonateForm={{entity_type:true,
                                    entity_id: Foundrising.id ? Foundrising.id.toString(): '',
                                    sum_of_money: '',
                                    comment: ''
                                }
                                }/>)
                }
            {!donate && (Foundrising.closing_date === undefined || Foundrising.closing_date === '') &&(
            <div className="card-body">
                <section className="section dark">
                    <h5 className="strong">
                        <strong>ID: {Foundrising.id}</strong>
                    </h5>
                
                    <p>Found ID: {Foundrising.found_id.toLocaleString()}</p>
                    <p>Description: {Foundrising.description}</p>
                    <p>Required sum: {Foundrising.required_sum.toLocaleString()}</p>
                    <p>Current sum: {Foundrising.current_sum.toLocaleString()}</p>
                    <p>Philantrops amount: {Foundrising.philantrops_amount.toLocaleString()}</p>
                    <p>Creation Date: {Foundrising.creation_date}</p>
                    <p>Closing Date: {Foundrising.closing_date}</p>
                    <button onClick={handleDonate}>Donate</button>
                </section>
            </div>)}
        </div>
    );
}
export default FoundrisingCard;