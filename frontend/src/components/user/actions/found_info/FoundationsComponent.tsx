import React, { useEffect, useState } from 'react';
import { FoundationAPI } from '../../../foundation/API';
import { FoundationTransfer } from '../../../foundation/Transfer';
import FoundationListExt from './FoundationList';


interface FoundationListProps{
    user_id: number
}
function FoundationListForm({user_id}: FoundationListProps) {
  const [error, setError] = useState<string | null>(null);


  //FoundationS
  const [Foundations, setFoundations] = useState<FoundationTransfer[]>([]);
  const [foundationsloading, setFoundationLoading] = useState(false);
  const [FoundationError, setFoundationError] = useState<string | undefined>(undefined);
  
  useEffect(() =>{
      async function loadFoundations() {
        setFoundationLoading(true);
          try{
              const data = await FoundationAPI.get();
                  setFoundations(data);
          }catch (e){
              if (e instanceof Error){
                setFoundationError(e.message);
              }
          }finally{
            setFoundationLoading(false);
          }
      }
      loadFoundations();
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
            <h1>Foundations</h1>
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
                <FoundationListExt user_id={user_id} foundations={Foundations} />
        </div>
    </div>
  );
}
export default FoundationListForm;