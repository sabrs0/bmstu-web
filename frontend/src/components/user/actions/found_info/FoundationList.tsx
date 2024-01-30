import { FoundationTransfer } from "../../../foundation/Transfer";
import FoundationCardExt from "./FoundationCard";

interface FoundationListExtProps{
    user_id: number
    foundations: FoundationTransfer[];
}
function FoundationListExt({user_id, foundations}: FoundationListExtProps){
     
  return (
            <div className="row">

              {foundations.map((foundation) => (
                <div key={foundation.id} className="cols-sm">
                  <FoundationCardExt user_id={user_id} Found={foundation}/>
                </div>
              ))}
            </div>

          );
          
}

export default FoundationListExt;