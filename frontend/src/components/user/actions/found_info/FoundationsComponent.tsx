import React, { useEffect, useState } from 'react';
import { FoundationAPI } from '../../../foundation/API';
import { FoundationTransfer } from '../../../foundation/Transfer';
import FoundationListExt from './FoundationList';
import { Container } from "@mui/material";

interface FoundationListProps{
    user_id: number
}
function FoundationListForm({user_id}: FoundationListProps) {
  const [error, setError] = useState<string | null>(null);


  //FoundationS
  const [Foundations, setFoundations] = useState<FoundationTransfer[]>([]);
  const [foundationsloading, setFoundationLoading] = useState(false);
  const [FoundationError, setFoundationError] = useState<string | undefined>(undefined);
  const handleClose = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
    window.localStorage.setItem('showFoundationForm', '0')
    window.localStorage.setItem('showDash', '1')
    window.location.reload();
}
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
        <h1>Foundations</h1>
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
            <FoundationListExt user_id={user_id} foundations={Foundations} />
        </div>
        <button className="button-login" onClick={handleClose}>Close</button>
    
    </div>
  );
}
export default FoundationListForm;