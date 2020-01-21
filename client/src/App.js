import { Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/styles';
import React from 'react';
import './App.css';
import RecordTableContainer from './containers/RecordTableContainer';

const GlobalCss = withStyles(theme => ({
  // @global is handled by jss-plugin-global.
  '@global': {
    '.MuiTypography-h1': {
      fontSize: '32px',
      padding: 0,
      fontWeight: 300,
      lineHeight: '1.3333',
      // opacity: 0.9,
      margin: 0,
    },
    '.MuiTypography-h2': {
      fontSize: '26px',
      lineHeight: '1.3333',
      // fontWeight: 600,
      // opacity: 0.9,
    },
    '.MuiTypography-h3': {
      fontSize: '22px',
      fontWeight: 600,
      lineHeight: '1.25',
      // opacity: 0.9,
    },
    '.MuiTypography-body1': {
      fontSize: '15px',
      lineHeight: '1.5',
      fontWeight: 400,
      opacity: 0.9,
      color: '#393b41',
      textRendering: 'optimizeLegibility',
    },
    '.MuiTypography-subtitle1': {
      fontSize: '14px',
      fontWeight: 400,
      // opacity: 0.9,
    },
    '.MuiTypography-subtitle2': {
      fontSize: '14px',
      fontWeight: 400,
      // opacity: 0.7,
    },
    '.MuiTypography-caption': {
      fontSize: '13px',
      fontWeight: 400,
    },
    '.MuiButton-label': {
      textTransform: 'none',
      fontWeight: 500,
      fontSize: '15px',
      textRendering: 'optimizeLegibility',
    },
    '.MuiStepIcon-root.MuiStepIcon-active': {
      // color: theme.palette.secondary.main,
    },
    '.MuiMenuItem-root': {
      fontFamily: 'Open Sans',
    },
  },
}))(() => null);


function App() {
  return (
    <div className="App">
      <GlobalCss />
      <header className="App-header">
        <Typography variant='h1'>Assets & Liabilities</Typography>
      </header>
      <div className="mainContent">
        <RecordTableContainer />
      </div>
    </div>
  );
}

export default App;
