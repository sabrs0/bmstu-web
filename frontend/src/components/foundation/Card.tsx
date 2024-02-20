import { Link } from "react-router-dom";
import { FoundationTransfer } from "./Transfer";

interface FoundationCardProps{
    Foundation: FoundationTransfer;
}

function FoundationCard(props: FoundationCardProps){
    const { Foundation } = props;
  return (
    <div className="card-basic">
      <div className="card-body">
        <section className="section dark">
        <Link to={"/foundations/" + Foundation.id} className="link-basic">
            <p>name : {Foundation.name}</p>
            <p>Current foundrising amount : {Foundation.cur_foudrising_amount.toLocaleString()}</p>
            <p>Closed foundrising amount : {Foundation.closed_foundrising_amount.toLocaleString()}</p>
            <p>Volunteer amount : {Foundation.volunteer_amount.toLocaleString()}</p>
            <p>Country : {Foundation.country}</p>
        </Link>
        </section>
      </div>
      </div>
  );
}
export default FoundationCard;