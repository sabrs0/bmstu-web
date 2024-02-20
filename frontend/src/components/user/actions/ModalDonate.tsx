import { Button, Modal } from "@mui/material";
import { useState } from "react";
import UserDonateShortForm from "./DonateShortForm";
import { UserDonate } from "../Transfer";

interface ModalDonateProps {
    user_id: number;
    entity_id: number;
    entity_type: boolean;
}

function ModalDonate({user_id, entity_id, entity_type} : ModalDonateProps){
  const [open, setOpen] = useState(false);
  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
      <button className="button-login" onClick={handleOpen} style={{margin: '10px'}}>
        Donate
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
            height: 600,
            backgroundColor: 'darkgrey',
            border: '2px solid #000',
            boxShadow: '0',
            
          }}
        >
          <h2 id="modal-modal-title">Donate</h2>
          <UserDonateShortForm user_id={user_id} entity_id={entity_id} entity_type={entity_type}/>
          <button className="button-login" style={{margin: '10px'}} onClick={handleClose}>
            Close
          </button>
        </div>
      </Modal>
    </div>
  );
};
export default ModalDonate;