import React, { useState } from 'react';
import './App.css';
import { BrowserRouter as Router, Routes, Route, NavLink } from 'react-router-dom';
import HomePage from './components/home/home';
import FoundationsPage from './components/foundation/Page';
import FoundationSinglePage from './components/foundation/Single';
import FoundrisingsPage from './components/foundrising/Page';
import LoginForm from './components/login/LoginPage';
import FoundationDashboard from './components/foundation/Dashboard';
import FoundrisingAddForm from './components/foundation/actions/FoundrisingAddForm';
import UserDashboard from './components/user/Dashboard';
import { ButtonGroup } from '@mui/material';
import Button from '@mui/material';
import { BorderBottom } from '@mui/icons-material';
function App() {
  //const [role, setRole] = useState('')  
  return (
      <Router>
        <header style={{borderBottom: '2px solid black'}}>
          <div className='nav-container'>
          <NavLink to='/' className='button-nav'>
            Home
          </NavLink>
          <NavLink to='/login' className='button-nav'>
            Login
          </NavLink>
          <NavLink to='/foundations' className='button-nav'>
            Foundations
          </NavLink>
          </div>
        </header>
        <div className='container'>
          <Routes>
            <Route path='/' element={<HomePage />} ></Route>
            <Route path='/login' element={<LoginForm/>} ></Route>
            <Route path='/foundations' element={<FoundationsPage />} ></Route>
            <Route path='/foundations/:id' element={<FoundationSinglePage />}></Route>
            <Route path='/foundrisings' element={<FoundrisingsPage />} ></Route>
            <Route path='/foundation_dash/:login' element={<FoundationDashboard />} ></Route>
            <Route path='/user_dash/:login' element={<UserDashboard /> } ></Route>
            
            
          </Routes>
        </div>
      </Router>
  )
}

export default App;
