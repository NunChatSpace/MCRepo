import './App.css';
import React from 'react';
import Grid from '@material-ui/core/Grid';
import Navbar from './Component/Navbar/Navbar';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import BuyPage from './Pages/BuyPage';
import HistoryPage from './Pages/HistoryPage';
import SuccessPage from './Pages/SuccessPage';
import ErrorPage from './Pages/ErrorPage';
function App() {
  
  if(localStorage.getItem('username') === ''){
    localStorage.setItem('username', generateUsername());
    localStorage.setItem('balance', '100');
  }
  else {
    var txt = localStorage.getItem('balance');
    console.log('init');
    if(txt === 'NaN' || txt === '')
    {
      localStorage.setItem('balance', '100');
    }
    else {
      var balance = parseFloat(txt) + 100;
      localStorage.setItem('balance', balance);
    }
  }
  

  return (
    <>
      <Router>
        <Grid container spacing={0}>
          <Grid item xs={2}>
            <Navbar username={localStorage.getItem('username')}/>
          </Grid>
          <Grid item xs={10}>
            <Switch>
              <Route path="/buy" exact component={BuyPage} />
              <Route path="/history" component={HistoryPage} />
              <Route path="/success" component={SuccessPage} />
              <Route path="/error" component={ErrorPage} />
            </Switch>
          </Grid>
        </Grid>
      </Router>
    </>
  );
}


function generateUsername() {
  var date = new Date();
  return "User" + date.getTime();
}
export default App;
