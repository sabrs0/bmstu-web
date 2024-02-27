import { Link } from "react-router-dom";
import { FoundrisingTransfer } from "../../../foundrising/Transfer";
import { useState } from "react";
import UserDonateToFoundForm from "../DonateFoundForm";
import { UserDonate } from "../../Transfer";
import UserDonateToFrisingForm from "../DonateFingForm";
import { ModalExpand } from "../../../foundation/descriptionList";
import ModalDonate from "../ModalDonate";

interface FoundrisingCardProps{
    user_id: number
    Found: FoundrisingTransfer;
}

function FoundrisingCardExt({user_id, Found}: FoundrisingCardProps){
    const  Foundrising  = Found;
    const[donate, setDonate] = useState(false);
    const handleDonate = () =>{
        setDonate(true)
        //window.location.reload()
    }
    return (
            <div className="card-basic">
                <div className="card-body">
                    <section className="section dark">
                        <h5 className="strong">
                            <strong>Description: {Foundrising.description.slice(0, 10)+(Foundrising.description.length > 10 ? '...' : '')}</strong>
                        </h5>
                    
                        <ModalExpand name={Foundrising.description} />
                        <p>Required sum: {Foundrising.required_sum.toLocaleString()}</p>
                        <p>Current sum: {Foundrising.current_sum.toLocaleString()}</p>
                        <p>Philantrops amount: {Foundrising.philantrops_amount.toLocaleString()}</p>
                        <p>Creation Date: {Foundrising.creation_date}</p>
                        <p>Closing Date: {Foundrising.closing_date}</p>
                        
                    </section>
                </div>
                <ModalDonate user_id={user_id} entity_id={Found.id ? Found.id : -1} entity_type={true}/>
            </div>
    );
}
export default FoundrisingCardExt;