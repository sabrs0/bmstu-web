import React from 'react';
import { FoundationTransfer } from './Transfer';

interface FoundationDetailProps {
  Foundation: FoundationTransfer;
}
export default function FoundationDetail({ Foundation }: FoundationDetailProps) {
  return (
    <div className='panel-container'>
          <section className="section dark">
          <h5 className="strong">
            <strong>id: {Foundation.id}</strong>
            </h5>
            <p>name : {Foundation.name}</p>
            <p>Current foundrising amount : {Foundation.cur_foudrising_amount.toLocaleString()}</p>
            <p>Closed foundrising amount : {Foundation.closed_foundrising_amount.toLocaleString()}</p>
            <p>Volunteer amount : {Foundation.volunteer_amount.toLocaleString()}</p>
            <p>Country : {Foundation.country}</p>
          </section>
    </div>
  );
}