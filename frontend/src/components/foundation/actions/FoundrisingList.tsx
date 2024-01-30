import React, { useEffect, useState } from 'react';
import { FoundationAPI } from '../API';
import { FoundrisingTransfer } from '../../foundrising/Transfer';
import FoundrisingList from '../../foundrising/List';


interface FoundrisingListProps{
    found_id: number
}
function FoundrisingListForm({found_id}: FoundrisingListProps) {
  const [error, setError] = useState<string | null>(null);


  //FOUNDRISINGS
  const [Foundrisings, setFoundrisings] = useState<FoundrisingTransfer[]>([]);
  const [foundrisingloading, setFoundrisingLoading] = useState(false);
  const [foundrisingError, setFoundrisingError] = useState<string | undefined>(undefined);
  
  useEffect(() =>{
      async function loadFoundrisings() {
        setFoundrisingLoading(true);
          try{
              const data = await FoundationAPI.findFrisingsByID(found_id);
                  
                  setFoundrisings(data);
              
          }catch (e){
              if (e instanceof Error){
                setFoundrisingError(e.message);
              }
          }finally{
            setFoundrisingLoading(false);
          }
      }
      loadFoundrisings();
  }, [found_id]);
  
  return (
    <div>
        {error && (
          <div className="row">
            <div className="card large error">
              <section>
                <p>
                  <span className="icon-alert inverse "></span> {error}
                </p>
              </section>
            </div>
          </div>
        )}
        <div>
            <h1>Foundrisings</h1>
            {error &&( 
                <div className="row">
                    <div className="card large error">
                        <section>
                            <p>
                                <span className="icon-alert inverse "></span>
                                {error}
                            </p>
                        </section>
                    </div>
                </div>
            )}
                <FoundrisingList Foundrisings={Foundrisings} />
        </div>
    </div>
  );
}
export default FoundrisingListForm;