import React, {  Component } from 'react'
import { Grid } from '@material-ui/core';
import HistoryTable from './SubComponent/HistoryTable';
import firebase from '../../firebaseConfig';


export default class HistoryView extends Component {
    intervalID = 0;
    constructor(props) {
        super(props);
        this.state = {
            MoonCoin: 0
        }

        this.intervalFlag = true;
        // console.log("to History view");
    }

    componentDidMount(){
        this.intervalID = setInterval(async () => {
           
            const testRef = await (await firebase.database().ref('MoonCoin').get()).val();
            var rtdbMoon = (1 / testRef.ExchangeRate);
            if(this.state.MoonCoin !== rtdbMoon)
            {
                this.setState({
                    MoonCoin: testRef.Remaining,
                });
            }        
        }, 1500);
    }

    componentWillUnmount() {
        clearInterval(this.intervalID);
    }

    precisionRoundMod(number, precision) {
        var factor = Math.pow(10, precision);
        var n = precision < 0 ? number : 0.01 / factor + number;
        return Math.round( n * factor) / factor;
    }

    render() {
        const style= {
            fontSize: 20,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            padding: 10,
        };
        return (
            <>
                <Grid container spacing={0}>
                    <Grid item xs={12}>
                        <div>
                            <div style={style}>
                                <label>Moon left ≈
                                    <span data-atd="moon-left-label" id="moon-left-label" style={{marginRight:20}}>
                                    {this.precisionRoundMod(this.state.MoonCoin, 8)}
                                    </span>
                                    MOON
                                </label>
                            </div>
                        </div>
                    </Grid>
                    <Grid item xs={12}>
                        <div style={{ height: 700 }}>
                            <HistoryTable />
                        </div>
                    </Grid>
                </Grid>
            </>
    
        )
    }

    
}