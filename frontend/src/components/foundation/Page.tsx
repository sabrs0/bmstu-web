import React,{ Fragment, useEffect, useState } from "react";
import FoundationList from "./List";
import { FoundationTransfer } from "./Transfer";
import { FoundationAPI } from "./API";
function FoundationsPage(){
    
    //const [Foundations, setFoundations] = useState<Foundation[]>(MOCK_FoundationS);
    const [Foundations, setFoundations] = useState<FoundationTransfer[]>([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | undefined>(undefined);
    //const [currentPage, setCurrentPage] = useState(1);

    
   /* const handleOnMoreClick = () => {
        if (Foundations.length >= currentPage){
            setCurrentPage((currentPage) => currentPage + 1);
        }
        
    };*/
    useEffect(() =>{
        async function loadFoundations() {
            setLoading(true);
            try{
                const data = await FoundationAPI.get();
                //if (currentPage === 1) {
                    setFoundations(data);
                //} else {
                //    setFoundations((Foundations) => [...Foundations, ...data]);
                //}
            }catch (e){
                if (e instanceof Error){
                    setError(e.message);
                }
            }finally{
                setLoading(false);
            }
        }
        loadFoundations();
    }, [/*currentPage*/]);
    return(
        <Fragment>
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
                <FoundationList foundations={Foundations} />

                {/*!loading && !error && (
                    <div className="row">
                    <div className="col-sm-12">
                        <div className="button-group fluid">
                        <button className="button default" onClick={handleOnMoreClick}>
                            More...
                        </button>
                        </div>
                    </div>
                    </div>
                )*/}
                {loading && ( 
                    <div className="center-page">
                        <span className="spinning primary"></span>
                        <p>Loading...</p>
                    </div>)
                }
        </Fragment>
    ); 
}

export default FoundationsPage;