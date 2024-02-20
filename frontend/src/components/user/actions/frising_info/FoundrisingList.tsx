import { FoundrisingTransfer } from "../../../foundrising/Transfer";
import FoundrisingCardExt from "./FoundrisingCard";

interface FoundrisingListExtProps{
    user_id: number
    foundrisings: FoundrisingTransfer[];
}
function FoundrisingListExt({user_id, foundrisings}: FoundrisingListExtProps){
     
  return (
            <div className="row">

              {foundrisings.map((Foundrising) => (
                (!Foundrising.closing_date || Foundrising.closing_date.length === 0) &&
                <div key={Foundrising.id} className="cols-sm">
                  <FoundrisingCardExt user_id={user_id} Found={Foundrising}/>
                </div>
              ))}
            </div>

          );
          
}

export default FoundrisingListExt;