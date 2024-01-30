import { FoundrisingTransfer } from "../../../foundrising/Transfer";
import FoundrisingCardExt from "./FoundrisingCard";

interface FoundrisingListExtProps{
    user_id: number
    Foundrisings: FoundrisingTransfer[];
}
function FoundrisingListExt({user_id, Foundrisings}: FoundrisingListExtProps){
     
  return (
            <div className="row">

              {Foundrisings.map((Foundrising) => (
                <div key={Foundrising.id} className="cols-sm">
                  <FoundrisingCardExt user_id={user_id} Found={Foundrising}/>
                </div>
              ))}
            </div>

          );
          
}

export default FoundrisingListExt;