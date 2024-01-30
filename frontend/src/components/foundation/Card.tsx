import { Link } from "react-router-dom";
import { FoundationTransfer } from "./Transfer";

interface FoundationCardProps{
    Foundation: FoundationTransfer;
}

function FoundationCard(props: FoundationCardProps){
    const { Foundation } = props;
  return (
    <div className="card border-dark mb-3" style={{maxWidth: "350px", border: "2px solid"}}>
      <div className="card-body">
        <section className="section dark">
        <Link to={"/foundations/" + Foundation.id}>
        <h5 className="strong">
            <strong>id: {Foundation.id}</strong>
            </h5>
        </Link>
            <p>name : {Foundation.name}</p>
            <p>Current foundrising amount : {Foundation.cur_foudrising_amount.toLocaleString()}</p>
            <p>Closed foundrising amount : {Foundation.closed_foundrising_amount.toLocaleString()}</p>
            <p>Volunteer amount : {Foundation.volunteer_amount.toLocaleString()}</p>
            <p>Country : {Foundation.country}</p>
        
        </section>
      </div>
      </div>
  );
}
export default FoundationCard;