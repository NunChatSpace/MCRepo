import { Button, Grid, TextField } from '@material-ui/core'
import React, { Component } from 'react'
import { Redirect } from 'react-router-dom';
import firebase from '../../firebaseConfig';
import {POST} from '../../Common/RequestAPI';
import {  Cookies } from "react-cookie";

export default class BuyForm extends Component {
    intervalID = 0;
    
    constructor(props) {
        super(props);
        this.username = localStorage.getItem('username');
        this.state = {
            balance: localStorage.getItem('balance'),
            exRateMOON: "0.02",
            exRateTHBT: "0.0",
            thbtValue: "0",
            moonValue: "0",
            slipPage: "100",
            responseData: {},
            ScreenState: 'Normal'
        }
        this.cookies = new Cookies();
        
        this.onBuy = this.onBuy.bind(this);
        this.onTHBTChanged = this.onTHBTChanged.bind(this);
        this.onMOONChanged = this.onMOONChanged.bind(this);
        this.onSlippageChanged = this.onSlippageChanged.bind(this);
        this.buyMOONCoin = this.buyMOONCoin.bind(this);

    }
    componentDidMount() {
        this.intervalID = setInterval(async () => {
            const testRef = await (await firebase.database().ref('MoonCoin').get()).val();
            var rtdbMoon = (1 / testRef.ExchangeRate);
            if(this.state.exRateTHBT !== rtdbMoon)
            {
                this.setState({
                    exRateTHBT: 1 / testRef.ExchangeRate,
                });
            }
        }, 1500);
    }

    componentWillUnmount() {
        clearInterval(this.intervalID);
    }
    onBuy(bs) {
        if (bs.Message === 'Success') {
            var newBalance = parseFloat(this.state.balance) - parseFloat(bs.Data.THBT);
            localStorage.setItem('balance', newBalance);
            this.setState({
                balance: newBalance,
                thbtValue: 0,
                moonValue: 0,
                responseData: bs,
                ScreenState: 'Success'
            });
           
        }
        else {
            this.setState({
                ScreenState: 'Error',
                responseData: bs,
            });
        }
    }

    onTHBTChanged(event) {
        // console.log('THBT Changed');
        if (isNaN(event.target.value)) {
            event.target.value = event.target.value.slice(0, -1)
        }
        if(event.target.value[0] === '0') {
            event.target.value = event.target.value.slice(1)
        }
        if (event.target.value === '') {
            event.target.value = 0;
        }
        if (parseFloat(event.target.value) > parseFloat(this.state.balance))
        {
            event.target.value = parseInt(this.state.balance);
        }
        this.setState({
            thbtValue: event.target.value,
            moonValue: (1 / this.state.exRateTHBT) * parseInt(event.target.value)
        });
    }

    onMOONChanged(event) {
        console.log(event.target.value);
        if (isNaN(event.target.value)) {
            event.target.value = event.target.value.slice(0, -1);
            console.log(event.target.value);
        }
        if (parseInt( event.target.value) > 1000)
        {
            event.target.value = event.target.value.slice(0, -1);
            console.log(event.target.value);
        }
        if (event.target.value === '') {
            event.target.value = 0;
            console.log(event.target.value);
        }
        var thbt = parseFloat(event.target.value) * this.state.exRateTHBT
        if (thbt > parseFloat(this.state.balance))
        {
            event.target.value = parseFloat(parseFloat(this.state.balance)/ this.state.exRateTHBT);
        }

        this.setState({
            thbtValue: thbt,
            moonValue: event.target.value
        });
    }

    onSlippageChanged(event) {
        // console.log(event.target.value);
        if (isNaN(event.target.value)) {
            event.target.value = event.target.value.slice(0, -1)
        }
        if (parseInt( event.target.value) > 100)
        {
            event.target.value = event.target.value.slice(0, -1)
        }
        if (event.target.value === '') {
            event.target.value = 0;
        }
        if(event.target.value[0] === '0') {
            event.target.value = event.target.value.slice(1)
        }
        this.setState({
            slipPage: parseInt(event.target.value),
        });
    }

    precisionRoundMod(number, precision) {
        number = number + '';
        var startIndex = number.indexOf(".");
        if(number.includes("e-"))
        {
            return '0.00';
        }
        if(startIndex < 0) {
            var zero = "";
            for(var i =0; i< precision; i++){
                zero = zero + "0";
            }
            return number + "." + zero;
        }
        else {
            var tmp = number.substring(0, startIndex + precision + 2);
            var isRoundUp = parseInt(tmp[tmp.length - 1]);
            if(isRoundUp >= 5) {
                var lastint = parseInt(tmp[tmp.length - 2]) + 1 ;
                tmp = tmp.substring(0, tmp.length - 2) + lastint;
                return tmp;
            }
            else {
                tmp = number.substring(0, startIndex + precision + 1);
                return tmp;
            }
        }
    }

    buyMOONCoin() {
        const url = 'http://localhost:8080/buy';
        var srMin  = 1 - (parseFloat(this.state.slipPage) / 100.0);
        var srMax  = 1 + (parseFloat(this.state.slipPage) / 100.0);
        const body = {
            Username: localStorage.getItem('username'),
            BuyWith: parseFloat(this.state.thbtValue),
            CurrentExchangeRate: (1 / this.state.exRateTHBT),
            SlippageRateMin: srMin,
            SlippageRateMax: srMax,
        }
        POST(url, body).then(response => {
            this.onBuy(response.data);
        })
    }

    render() {
        var componentWidth = {
            width: 400,
        };
        
        var commonSyle = {
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            padding: 10,
        }

        var style = {
            fontSize: 40,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            padding: 10,
        }
    
        var styleBalance = {
            fontSize: 20,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            padding: 10,
        }

        if (this.state.ScreenState === 'Normal') {
            return (
                <div style={this.componentWidth}>
                    <Grid container spacing={0}>
                        <Grid item xs={12}>
                            <div style={style}>
                                <label>
                                    MOON ≈
                                    <span data-atd="moon-price-label" style={{marginRight:20}}>
                                        {this.precisionRoundMod(this.state.exRateTHBT, 2)}
                                    </span>
                                    THBT
                                </label>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                            <div style={styleBalance}>
                                <label>
                                    You have
                                    <span data-atd="balance-label" style={{marginRight:20}}>
                                        {this.precisionRoundMod(this.state.balance, 2)}
                                    </span>
                                    THBT
                                </label>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                            <div style={commonSyle}>
                                <div style={componentWidth}>
                                <Grid container spacing={0}>
                                    <Grid item xs={12}>
                                        <label style={componentWidth}>Amount to buy (THBT)</label>
                                    </Grid>
                                    <Grid item xs={12}>
                                        <TextField style={componentWidth} id="thbt-input" data-atd="thbt-input" onChange={this.onTHBTChanged} value={this.state.thbtValue}></TextField>
                                    </Grid>
                                </Grid>
                                </div>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                            <div style={commonSyle}>
                                <div style={componentWidth}>
                                <Grid container spacing={0}>
                                    <Grid item xs={12}>
                                        <label style={componentWidth}>Amount MOON</label>
                                    </Grid>
                                    <Grid item xs={12}>
                                        <TextField style={componentWidth} id="moon-input"  data-atd="moon-input" onChange={this.onMOONChanged} value={this.state.moonValue}></TextField>
                                    </Grid>
                                </Grid>
                                </div>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                            <div style={commonSyle}>
                                <div style={componentWidth}>
                                    <div style={{ paddingTop: 30}}>
                                    <Grid container spacing={0}>
                                        <Grid item xs={12}>
                                            <label style={componentWidth}>Slippage Tolerance (%)</label>
                                        </Grid>
                                        <Grid item xs={12}>
                                            <TextField style={componentWidth} id="slippage-input"  data-atd="slippage-input" onChange={this.onSlippageChanged} value={this.state.slipPage}></TextField>
                                        </Grid>
                                    </Grid>
                                    </div>
                                </div>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                            <div style={commonSyle}>
                            <Button variant="contained" color="primary" id="buy-btn" data-atd="buy-btn" style={componentWidth} onClick={this.buyMOONCoin}> Buy</Button>
                            </div>
                           
                        </Grid>
                    </Grid>
                </div>
            );
        }
        else if (this.state.ScreenState === 'Success') {
            return (<Redirect to={{
                pathname: '/success',
                state: this.state.responseData
                }}/>)
        }
        else {
            return (<Redirect to={{
                pathname: '/error',
                state: this.state.responseData
            }}/>);
        }
    }
}
