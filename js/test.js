
var TEST_WS = {}

var TEST_READ_CHANNEL = {}

var OrchestratorRequest = {

    "request_option": '',
    "request_target":'',
    "query": ''

}

var OrchestratorResponse = {

    "server_message":'',
    "query_result":''

}

function Test_BodySet(){


    var inject = 
    `
    <input id="test-target" type="text" value="" placeholder="test-target-name">
    <input id="test-query" type="text" value="" placeholder="test:query,input">
    <input type="button" onclick="Test_OrchestratorQuery()" value="run">
    
    `

    $('body').html(inject)

}


function Test_OrchestratorInit(){

    var or_url = window.location.href

    var proto_url = or_url.split('://') 

    var proto = proto_url[0]

    var urlbody_trail = proto_url[1].split('/')

    var urlbody = urlbody_trail[0]

    var trail_path = '/osock/front-test'

    if (proto == 'http'){

        proto = 'ws://'

        var urlbody_port = urlbody.split(':')

        urlbody = urlbody_port[0] + ':' + '7331'

    } else if (proto == 'https'){

        proto = 'wss://'
    }

    var ws_url = proto + urlbody + trail_path

    console.log(ws_url)

    if (!REQ_KEY){
        alert("request key is empty")
        return {}
    }

    var ws = new WebSocket(ws_url)

    ws.onopen = function (){

        oreq = JSON.parse(JSON.stringify(OrchestratorRequest))

        oreq["request_option"] = REQ_KEY

        ws.send(JSON.stringify(oreq))

    }

    ws.onmessage = function (evnt){

        var oresp = evnt.data

        TEST_READ_CHANNEL = JSON.parse(oresp)

    }

    
    return ws

}

async function Test_CheckReadChannelWithTimeOut(limit){


    var counter = 0


    while (counter < limit){

        if (JSON.stringify(TEST_READ_CHANNEL) === '{}'){

            counter += 1 

        } else {
            console.log('read signal received')
            break
        }

        await Test_Delay(1)

    }

    if (counter >= limit){

        alert('time limit exceeded')

        return 1

    }

    return 0

}


async function Test_OrchestratorQuery(){

    var test_target = $('#test-target').val()
    var test_query = $('#test-query').val()

    if(!test_target){
        alert('feed target')
        return
    }

    if(!test_query){
        alert('feed query')
        return
    }

    oreq = JSON.parse(JSON.stringify(OrchestratorRequest))

    oreq["request_target"] = test_target

    oreq["query"] = test_query

    TEST_READ_CHANNEL = {}

    TEST_WS.send(JSON.stringify(oreq))

    var ret_signal = await Test_CheckReadChannelWithTimeOut(10000)

    if (ret_signal == 1){
        return
    }


    console.log(TEST_READ_CHANNEL["server_message"])

    alert(atob(TEST_READ_CHANNEL["query_result"]))


}



function Test_Delay(time) {
    return new Promise(resolve => setTimeout(resolve, time));
}


Test_BodySet()

TEST_WS = Test_OrchestratorInit()

console.log('test ready')