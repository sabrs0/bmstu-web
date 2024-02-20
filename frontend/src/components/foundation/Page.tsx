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
            <h1 style={{textAlign: 'center'}}>Foundations</h1>
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
            <div className="wrapper">
                <FoundationList foundations={Foundations} />
                {loading && ( 
                    <div className="center-page">
                        <span className="spinning primary"></span>
                        <p>Loading...</p>
                    </div>)
                }
            </div>
        </Fragment>
    ); 
}

export default FoundationsPage;