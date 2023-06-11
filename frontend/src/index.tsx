import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Post from './components/forms/post';
import Main from './pages/main';
import NotFound from './pages/404';
import './components/style.css'
import OneThread from './pages/one-thread';
import Admin from './components/forms/admin';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  
  <>
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path='*' element={<NotFound errorText={''} />}/>
        <Route path='/' element={<Main page={undefined}/>} />
        <Route path='/form' element={<Post />} />
        <Route path='/all' element={<OneThread />}/>
        <Route path='/admin' element={<Main page={<Admin/>}/>} />
      </Routes>
    </Router>
  </React.StrictMode>
  <link rel="stylesheet" id="style-direction" href="./index.css"></link>
  </>
);

