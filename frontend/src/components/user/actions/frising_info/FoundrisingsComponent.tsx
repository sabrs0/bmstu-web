import React, { useEffect, useState } from 'react';
import { FoundrisingAPI } from '../../../foundrising/API';
import { FoundrisingTransfer } from '../../../foundrising/Transfer';
import FoundrisingListExt from './FoundrisingList';


interface FoundrisingListProps{
    user_id: number
}
function FoundrisingListForm({user_id}: FoundrisingListProps) {
  const [error, setError] = useState<string | null>(null);


  //FoundrisingS
  const [Foundrisings, setFoundrisings] = useState<FoundrisingTransfer[]>([]);
  const [Foundrisingsloading, setFoundrisingLoading] = useState(false);
  const [FoundrisingError, setFoundrisingError] = useState<string | undefined>(undefined);
  
  useEffect(() =>{
      async function loadFoundrisings() {
        setFoundrisingLoading(true);
          try{
              const data = await FoundrisingAPI.get();
                console.log('Foundrisings before cast are: ', data);
                  setFoundrisings(data);
                  console.log('Foundrisings after cast are: ', Foundrisings);
          }catch (e){
              if (e instanceof Error){
                setFoundrisingError(e.message);
                setError(e.message)
              }
          }finally{
            setFoundrisingLoading(false);
          }
      }
      loadFoundrisings();
  }, []);
  
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
                <FoundrisingListExt user_id={user_id} foundrisings={Foundrisings} />
        </div>
    </div>
  );
}
export default FoundrisingListForm;