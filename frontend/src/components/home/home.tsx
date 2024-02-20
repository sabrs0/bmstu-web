import React from 'react';
import home from './img/home.svg'
function HomePage() {
    return (
        <div className="home-page">
          <h1 className="foundations-title">Foundations</h1>
          <div className="content-box">
            <p>
              On this site you will find all the necessary information and tools for interacting with the charitable foundation.
              We are committed to working with you to achieve our shared goals of philanthropy and helping those in need.
            </p>
          </div>
          <img className='image' src={home} />
        </div>
      );
}

export default HomePage;