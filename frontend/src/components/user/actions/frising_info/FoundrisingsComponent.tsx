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
  const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
    window.localStorage.setItem('showFoundrisingsForm', '0')
    window.localStorage.setItem('showDash', '1')
    window.location.reload();
}
  useEffect(() =>{
      async function loadFoundrisings() {
        setFoundrisingLoading(true);
          try{
              const data = await FoundrisingAPI.get();
                  setFoundrisings(data);
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
    <div  className='form-container'>
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
        <h1>Foundrisings</h1>
        <div className='wrapper'>
            
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
        <button className="button-login" onClick={handleClose}>Close</button>
    </div>
  );
}
export default FoundrisingListForm;