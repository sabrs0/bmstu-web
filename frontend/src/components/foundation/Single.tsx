import React, { useEffect, useState } from 'react';
import { FoundationAPI } from './API';
import FoundationDetail from './Detail';
import { FoundationTransfer } from './Transfer';
import { FoundrisingTransfer } from '../foundrising/Transfer';
import { useParams } from 'react-router-dom';
import { FoundrisingAPI } from '../foundrising/API';
import FoundrisingList from '../foundrising/List';

function FoundationSinglePage(props: any) {
  const [loading, setLoading] = useState(false);
  const [Foundation, setFoundation] = useState<FoundationTransfer | null>(null);
  const [error, setError] = useState<string | null>(null);
  const params = useParams();
  const id = Number(params.id);

  useEffect(() => {
    setLoading(true);
    FoundationAPI
      .find(id)
      .then((data) => {
        setFoundation(data);
        setLoading(false);
      })
      .catch((e) => {
        setError(e);
        setLoading(false);
      });
  }, [id]);


  //FOUNDRISINGS
  const [Foundrisings, setFoundrisings] = useState<FoundrisingTransfer[]>([]);
  const [foundrisingloading, setFoundrisingLoading] = useState(false);
  const [foundrisingError, setFoundrisingError] = useState<string | undefined>(undefined);
  //const [currentPage, setCurrentPage] = useState(1);

  
  /*const handleOnMoreClick = () => {
      setCurrentPage((currentPage) => currentPage + 1);
  };*/
  useEffect(() =>{
      async function loadFoundrisings() {
        setFoundrisingLoading(true);
          try{
              const data = await FoundationAPI.findFrisingsByID(id/*, currentPage*/);
              //if (currentPage === 1) {
                  
                  setFoundrisings(data);
              //} else {
              //    setFoundrisings((Foundrisings) => [...Foundrisings, ...data]);
              //}
          }catch (e){
              if (e instanceof Error){
                setFoundrisingError(e.message);
              }
          }finally{
            setFoundrisingLoading(false);
          }
      }
      loadFoundrisings();
  }, [id]/*, [currentPage]*/);
  //FOUNDRISINGS STOP
  
  return (
    <div>
      <>
        <h1>Foundation Details</h1>

        {loading && (
          <div className="center-page">
            <span className="spinner primary"></span>
            <p>Loading...</p>
          </div>
        )}

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

        {Foundation && (<FoundationDetail Foundation={Foundation} />)}
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
                {/*foundrisingloading && ( 
                    <div className="center-page">
                        <span className="spinning primary"></span>
                        <p>Loading...</p>
                    </div>)*/
                }
        </div>
      </>
    </div>
  );
}

export default FoundationSinglePage;