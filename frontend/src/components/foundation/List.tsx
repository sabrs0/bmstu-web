import { FoundationTransfer } from "./Transfer";
import FoundationCard from "./Card";

interface FoundationListProps{
    foundations: FoundationTransfer[];
}
function FoundationList({foundations}: FoundationListProps){
     
  return (
            <div className="row">

              {foundations.map((foundation) => (
                <div key={foundation.id} className="cols-sm">
                  <FoundationCard Foundation={foundation}/>
                </div>
              ))}
            </div>

          );
          
}

export default FoundationList;