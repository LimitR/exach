import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Post from './components/forms/post';
import Main from './components/pages/main';
import NotFound from './components/pages/404';
import './components/style.css'
import OneThread from './components/pages/one-thread';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path='*' element={<NotFound errorText={''} />}/>
        <Route path='/' element={<Main />} />
        <Route path='/form' element={<Post />} />
        <Route path='/all' element={<OneThread />}/>
      </Routes>
    </Router>
  </React.StrictMode>
);

