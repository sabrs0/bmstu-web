import { Link } from "react-router-dom";
import { FoundationTransfer } from "../../../foundation/Transfer";
import { useState } from "react";
import UserDonateToFoundForm from "../DonateFoundForm";
import { UserDonate } from "../../Transfer";
import ModalDonate from "../ModalDonate";


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
        <div className="card-basic">
            <div className="card-body">
                <section className="section dark">
                    <h5 className="strong">
                        <strong>name : {Foundation.name}</strong>
                    </h5>
                    <p>Current foundrising amount : {Foundation.cur_foudrising_amount.toLocaleString()}</p>
                    <p>Closed foundrising amount : {Foundation.closed_foundrising_amount.toLocaleString()}</p>
                    <p>Volunteer amount : {Foundation.volunteer_amount.toLocaleString()}</p>
                    <p>Country : {Foundation.country}</p>
                </section>
            </div>
            <ModalDonate user_id={user_id} entity_id={Found.id ? Found.id : -1} entity_type={false}/>
        </div>
            
    );
}
export default FoundationCard;