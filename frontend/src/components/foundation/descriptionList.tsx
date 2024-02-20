import { useEffect, useState } from "react";
import { FoundationTransfer } from "./Transfer";
import { FoundationAPI } from "./API";
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import DoneIcon from '@mui/icons-material/Done';
import Button from '@mui/material/Button';
import Modal from '@mui/material/Modal';
interface FoundationChoiceProps {
    handleSetID: (id: string)=> void;
}
function FoundationChoice({handleSetID} : FoundationChoiceProps){
    const [selectedName, setSelectedName] = useState('')

    const handleSelectName = (name: string) => {
       setSelectedName(name);
    };
    return(
        <div className="input-row">
            <label htmlFor="name">{selectedName}</label>
            <ModalComponent handleSetName={handleSelectName} handleSetID={handleSetID} />
        </div>
    )
}
export default FoundationChoice;
interface ModalCompProps {
    handleSetID: (id: string)=> void;
    handleSetName: (name: string)=> void;
}

function ModalComponent({handleSetID, handleSetName} : ModalCompProps){
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
          <h2 id="modal-modal-title">Choose Foundation</h2>
          <DescriptionList handleSetName={handleSetName} handleSetID={handleSetID} handleClose={handleClose}/>
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
    handleSetName: (name: string)=> void;
    handleClose: () => void;
}

function DescriptionList({handleSetID, handleSetName, handleClose}: DescriptionListProps){
    const [showFull, setShowFull] = useState<string | null>(null);
    const [Foundations, setFoundations] = useState<FoundationTransfer[]>([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | undefined>(undefined);
    const [id, setId] = useState<number | undefined>(undefined)
    const [isEmpty, setIsEmpty] = useState(true)
    
    useEffect(() =>{
        async function loadFoundations() {
            setLoading(true);
            try{
                const data = await FoundationAPI.get();
                data.length == 0 ? setIsEmpty(true): setIsEmpty(false);
                setFoundations(data);
                
            }catch (e){
                if (e instanceof Error){
                    setError(e.message);
                }
            }finally{
                setLoading(false);
            }
        }
        loadFoundations();
    }, []);

    const handleExpand = (nameiption: string | null) =>{
        
    }
    const handleChoose = (id: number | undefined, name: string | undefined) =>{
        if (id && name){
            handleSetID(id.toString());
            handleSetName(name.slice(0, 10)+(name.length > 10 ? '...' : ''));
            handleClose();
            return;
        }else{
            console.log("NILS: ", id, name);
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
          <li key={`Foundations`}>
              <ul>
                {!isEmpty && Foundations.map((item) => (
                  <ListItem key={`item-${item.id}`} sx={{border: '2px solid black'}}>
                      <div>
                        <ListItemText primary={item.name.slice(0, 10)+(item.name.length > 10 ? '...' : '')} />
                        <button onClick={() =>handleChoose(item.id, item.name)} className="button-login" style={{width:'75px', fontSize: '11px'}}>Choose</button>
                        <ModalExpand  name={item.name}/>
                        {/*<button className="button-login" style={{width:'75px', fontSize: '11px'}}>Expand</button>*/}
                      </div>
                  </ListItem>
                ))}
                {isEmpty && <div>No active Foundations</div>}
              </ul>
            </li>
          
        </List>
      );
}


interface ModalExpandProps {
  name: string
}
export function ModalExpand({name}: ModalExpandProps){
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
            <label htmlFor="nameiption">{name}</label>
          </div>
          <Button className="button-login" style={{margin: '10px'}} onClick={handleClose}>
            Close
          </Button>
        </div>
      </Modal>
    </div>
  );
};
