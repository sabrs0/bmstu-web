import { useEffect, useState } from "react";
import { FoundrisingTransfer } from "./Transfer";
import { FoundrisingAPI } from "./API";
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import DoneIcon from '@mui/icons-material/Done';
import Button from '@mui/material/Button';
import Modal from '@mui/material/Modal';
interface FoundrisingChoiceProps {
    handleSetID: (id: string)=> void;
}
function FoundrisingChoice({handleSetID} : FoundrisingChoiceProps){
    const [selectedDescr, setSelectedDescr] = useState('')

    const handleSelectDescr = (descr: string) => {
       setSelectedDescr(descr);
    };
    return(
        <div className="input-row">
            <label htmlFor="descr">{selectedDescr}</label>
            <ModalComponent handleSetDescr={handleSelectDescr} handleSetID={handleSetID} />
        </div>
    )
}
export default FoundrisingChoice;
interface ModalCompProps {
    handleSetID: (id: string)=> void;
    handleSetDescr: (descr: string)=> void;
}

function ModalComponent({handleSetID, handleSetDescr} : ModalCompProps){
  const [open, setOpen] = useState(false);
  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
      <Button className="button-login" variant="contained" onClick={handleOpen}>
        Choose
      </Button>
      <Modal
        open={open}
        onClose={handleClose}
      >
        <div
            className="form-container"
          style={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            width: 400,
            height: 300,
            backgroundColor: 'darkgrey',
            border: '2px solid #000',
            boxShadow: '0',
            
          }}
        >
          <h2 id="modal-modal-title">Choose Foundrising</h2>
          <DescriptionList handleSetDescr={handleSetDescr} handleSetID={handleSetID} handleClose={handleClose}/>
          <Button className="button-login" style={{margin: '10px'}} onClick={handleClose}>
            Close
          </Button>
        </div>
      </Modal>
    </div>
  );
};


interface DescriptionListProps {
    handleSetID: (id: string)=> void;
    handleSetDescr: (descr: string)=> void;
    handleClose: () => void;
}

function DescriptionList({handleSetID, handleSetDescr, handleClose}: DescriptionListProps){
    const [showFull, setShowFull] = useState<string | null>(null);
    const [Foundrisings, setFoundrisings] = useState<FoundrisingTransfer[]>([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | undefined>(undefined);
    const [id, setId] = useState<number | undefined>(undefined)
    const [isEmpty, setIsEmpty] = useState(true)
    
    useEffect(() =>{
        async function loadFoundrisings() {
            setLoading(true);
            try{
                const data = await FoundrisingAPI.get();
                data.length == 0 ? setIsEmpty(true): setIsEmpty(false);
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

    const handleExpand = (description: string | null) =>{
        
    }
    const handleChoose = (id: number | undefined, descr: string | undefined) =>{
        if (id && descr){
            handleSetID(id.toString());
            handleSetDescr(descr.slice(0, 10)+(descr.length > 10 ? '...' : ''));
            handleClose();
            return;
        }else{
            console.log("NILS: ", id, descr);
        }
            
    }
    return (
        <List 
          sx={{
            width: '100%',
            maxWidth: '200px',
            bgcolor: 'grey',
            position: 'relative',
            overflow: 'auto',
            borderRadius: '10px',
            border: '2px solid black',
            maxHeight: 200,
            '& ul': { padding: 0 },
          }}
          subheader={<li />}
        >
          <li key={`Foundrisings`}>
              <ul>
                {!isEmpty && Foundrisings.map((item) => (
                  <ListItem key={`item-${item.id}`} sx={{border: '2px solid black'}}>
                      <div>
                        <ListItemText primary={item.description.slice(0, 10)+(item.description.length > 10 ? '...' : '')} />
                        <button onClick={() =>handleChoose(item.id, item.description)} className="button-login" style={{width:'75px', fontSize: '11px'}}>Choose</button>
                        <ModalExpand  description={item.description}/>
                        {/*<button className="button-login" style={{width:'75px', fontSize: '11px'}}>Expand</button>*/}
                      </div>
                  </ListItem>
                ))}
                {isEmpty && <div>No active Foundrisings</div>}
              </ul>
            </li>
          
        </List>
      );
}


interface ModalExpandProps {
  description: string
}
function ModalExpand({description}: ModalExpandProps){
  const [open, setOpen] = useState(false);
  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
      <button className="button-login" style={{width:'75px', fontSize: '11px'}} onClick={handleOpen}>
        Expand
      </button>
      <Modal
        open={open}
        onClose={handleClose}
      >
        <div
            className="form-container"
          style={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            width: 400,
            height: 300,
            backgroundColor: 'darkgrey',
            border: '2px solid #000',
            boxShadow: '0',
            
          }}
        >
          <div>
            <label htmlFor="description">{description}</label>
          </div>
          <Button className="button-login" style={{margin: '10px'}} onClick={handleClose}>
            Close
          </Button>
        </div>
      </Modal>
    </div>
  );
};

