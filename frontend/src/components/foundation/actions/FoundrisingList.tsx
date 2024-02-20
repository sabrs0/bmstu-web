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
  const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
    window.location.reload();
}
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
    <div className='form-container'>
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
        <button className="button-login" onClick={handleClose}>Close</button>
    </div>
  );
}
export default FoundrisingListForm;