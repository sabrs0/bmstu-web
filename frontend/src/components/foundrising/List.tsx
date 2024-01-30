import { FoundrisingTransfer } from "./Transfer";
import FoundrisingCard from "./Card";

interface FoundrisingListProps{
    Foundrisings: FoundrisingTransfer[];
}
function FoundrisingList({Foundrisings}: FoundrisingListProps){
     
  return (
            <div className="row">

              {Foundrisings.map((Foundrising) => (
                <div key={Foundrising.id} className="cols-sm">
                  <FoundrisingCard Foundrising={Foundrising}/>
                </div>
              ))}
            </div>

          );
          
}

export default FoundrisingList;