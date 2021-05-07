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
                                <label>You bought {state.Data.MOON} MOON</label>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                        <div style={commonSyle}>
                                <label>With {state.Data.THBT} THBT</label>
                            </div>
                        </Grid>
                        <Grid item xs={12}>
                            <div style={commonSyle}>
                                <Button variant="contained" color="primary" style={componentWidth} onClick={() => {this.setState({screenState: 'toBuy'})}}>HOME</Button>
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
