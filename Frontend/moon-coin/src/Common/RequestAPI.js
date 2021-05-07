import axios from 'axios';

export const GET = async (url) => {
    var result = {}
    var rows = []
    await axios({
        baseURL: url,
        method: 'GET',
        headers: {
            'Access-Control-Allow-Origin': '*',
            'Access-Control-Allow-Methods': 'GET,POST,HEAD,PUT,DELETE,PATCH',
            'Access-Control-Allow-Headers': 'access-control-allow-origin, Origin, Content-Type, Accept, Content-Length, Authorization',
        },
    }).then(response => {
        if(response.data.DataLength > 0) {
            // console.log(response.data.Data);
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
            result = rows;
        }
        console.log(rows);
    })
    return result;
}

export const POST = async (url, body) => {
    return axios({
        baseURL: url,
        method: 'POST',
        headers: {
            'Access-Control-Allow-Origin': '*',
            'Access-Control-Allow-Methods': 'GET,POST,HEAD,PUT,DELETE,PATCH',
            'Access-Control-Allow-Headers': 'access-control-allow-origin, Origin, Content-Type, Accept, Content-Length, Authorization',
        },
        data: body
    })
}