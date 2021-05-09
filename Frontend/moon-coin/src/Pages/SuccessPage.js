import { Button, Grid } from '@material-ui/core';
import React, { Component } from 'react'
import { Redirect } from 'react-router';

export default class SuccessPage extends Component {
    constructor(props) {
        super(props)
        this.state = {
            screenState: 'Normal',
        }
    }
    render() {
        const { state } = this.props.location
        var componentWidth = {
            width: 400,
        };
        var commonSyle = {
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            padding: 10,
        }
        if(this.state.screenState === 'Normal') {
            return (
                <div>
                    <Grid container spacing={0}>
                        <Grid item xs={12}>
                            <div style={commonSyle}>
                                <label>Success</label>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                            <div style={commonSyle}>
                                <label>You bought 
                                    <span data-atd="success-moon-label" style={{marginRight:20}}>
                                        {state.Data.MOON}
                                    </span>
                                     MOON
                                </label>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                        <div style={commonSyle}>
                                <label>With 
                                    <span data-atd="success-thbt-label" style={{marginRight:20}}>
                                        {state.Data.THBT} 
                                    </span>
                                    THBT
                                </label>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                            <div style={commonSyle}>
                                <Button variant="contained" color="primary" id="back-btn" data-atd="back-btn" style={componentWidth} onClick={() => {this.setState({screenState: 'toBuy'})}}>HOME</Button>
                            </div>
                        </Grid>
                    </Grid>
                </div>
            )
        }
        else {
            return (<Redirect to={{pathname: '/buy'}}/>)
        }
        
    }
}
