import React,{ Fragment, useEffect, useState } from "react";
import FoundrisingList from "./List";
import { FoundrisingTransfer } from "./Transfer";
import { FoundrisingAPI } from "./API";
function FoundrisingsPage(){
    
    const [Foundrisings, setFoundrisings] = useState<FoundrisingTransfer[]>([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | undefined>(undefined);
    
    useEffect(() =>{
        async function loadFoundrisings() {
            setLoading(true);
            try{
                const data = await FoundrisingAPI.get();
                    setFoundrisings(data);
                
            }catch (e){
                if (e instanceof Error){
                    setError(e.message);
                }
            }finally{
                setLoading(false);
            }
        }
        loadFoundrisings();
    }, []);
    return(
        <Fragment>
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

                {loading && ( 
                    <div className="center-page">
                        <span className="spinning primary"></span>
                        <p>Loading...</p>
                    </div>)
                }
        </Fragment>
    ); 
}

export default FoundrisingsPage;