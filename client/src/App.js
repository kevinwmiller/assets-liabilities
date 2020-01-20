import { Typography } from '@material-ui/core';
import React from 'react';
import './App.css';
import RecordEditForm from './components/RecordEditForm';
import RecordTableContainer from './containers/RecordTableContainer';


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Typography variant='h2'>Assets & Liabilities</Typography>
      </header>
      <RecordTableContainer />
      <RecordEditForm />
    </div>
  );
}

export default App;
