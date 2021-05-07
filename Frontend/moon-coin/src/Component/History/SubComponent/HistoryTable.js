import React from 'react'
import { DataGrid } from '@material-ui/data-grid';
import axios from 'axios';

export default class HistoryTable  extends React.Component {
    state = {
        loading: true,
        responseData: null,
    }

    async componentDidMount(){
        var rows = [];
        await axios({
            baseURL: 'http://localhost:8079/userinfo',
            method: 'GET',
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Access-Control-Allow-Methods': 'GET,POST,HEAD,PUT,DELETE,PATCH',
                'Access-Control-Allow-Headers': 'access-control-allow-origin, Origin, Content-Type, Accept, Content-Length, Authorization',
            },
        }).then(response => {
            if(response.data.DataLength > 0) {
                for (var i = 0; i < response.data.DataLength; i++) {
                    var respData = response.data.Data; 
                    rows.push({ 
                        id: i + 1, 
                        buyDate: respData[i].BuyDate, 
                        username: respData[i].Username, 
                        thbt: respData[i].THBT, 
                        moon: respData[i].MOON, 
                        rate: respData[i].Rate 
                    })
                }
                console.log(rows);
                this.setState({
                    responseData:rows,
                    loading:false,
                })
            }
        })
    }
    render() {
        
        const columns = [
            { field: 'id', headerName: 'No.', width: 100 },
            { field: 'buyDate', headerName: 'Date and time', width: 200 },
            { field: 'username', headerName: 'ID', width: 200 },
            { field: 'thbt', headerName: 'THBT', width: 130 },
            { field: 'moon', headerName: 'MOON', width: 130 },
            { field: 'rate', headerName: 'RATE', width: 500 }
        ];
        console.log((this.state.loading || !this.state.responseData));
        if(this.state.loading || !this.state.responseData) {
            return (<>Loading...</>);
        }
        else {
            return (<DataGrid rows={this.state.responseData} columns={columns} pageSize={11} />);
        }
    }

    
}