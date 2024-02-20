import { Link } from "react-router-dom";
import { FoundrisingTransfer } from "./Transfer";

interface FoundrisingCardProps{
    Foundrising: FoundrisingTransfer;
}

function FoundrisingCard({Foundrising}: FoundrisingCardProps){
  console.log("In Foundrising card. Foundrising is - ", Foundrising.description)
  return (
    <div className="card-basic">
      <div className="card-body">
        <section className="section dark">
        <h5 className="strong">
            <strong>Description: {Foundrising.description}</strong>
            </h5>
            <p>Required sum: {Foundrising.required_sum.toLocaleString()}</p>
            <p>Current sum: {Foundrising.current_sum.toLocaleString()}</p>
            <p>Philantrops amount: {Foundrising.philantrops_amount.toLocaleString()}</p>
            <p>Creation Date: {Foundrising.creation_date}</p>
            <p>Closing Date: {Foundrising.closing_date}</p>
        
        </section>
      </div>
      </div>
  );
}
export default FoundrisingCard;