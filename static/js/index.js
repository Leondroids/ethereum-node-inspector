$(document).ready(function () {
    console.log("ready");
    $('#myTabs').tabs();
    loadStatus();
    loadInfo();
    loadAccounts()
});

function loadStatus() {
    console.log("loading...");
    $.get("/node/status", function (data, status) {
        body = $('#tbodyStatus')
        body.empty();

        console.log(status)
        console.log(data)

        json = data;

        body.append(createRow('Is Mining', json.isMining, ''));
        body.append(createRow('Hashrate', json.hashrate + '/s', ''));
        body.append(createRow('Is Listening', json.isListening, ''));
        body.append(createRow('Peer Count', json.peerCount, ''));
        body.append(createRow('Current Block', nc(json.syncStatus.currentBlock), ''));
        body.append(createRow('Highest Block', nc(json.syncStatus.highestBlock), ''));
        body.append(createRow('Starting Block', nc(json.syncStatus.startingBlock), ''));
        body.append(createRow('Gas Price', json.gasPrice, ''));

    });
}

function loadInfo() {
    console.log("loading info...");
    $.get("/node/info", function (data, status) {
        body = $('#tbodyInfo')
        body.empty();

        console.log(status)
        console.log(data)

        json = data;

        body.append(createRow('Client Version', json.clientVersion, ''));
        body.append(createRow('Protocol Version', json.protocolVersion, ''));
        body.append(createRow('NetworkID', json.networkId, ''));
    });
}
function loadAccounts() {
    console.log("loading accounts...");
    $.get("/node/accounts", function (data, status) {
        body = $('#tbodyAccounts')
        body.empty();

        console.log(status)
        console.log(data)

        json = data;

        body.append(createRow('Coinbase', json.coinbase, ''));
        body.append(createRow('................', '..........................................', ''));
        json.accounts.forEach(function (value, index) {
            body.append(createRow('Account ' + index, value, ''));
        });
    });
}

function createRow(rowName, value) {
    row =
        '<tr>' +
        '<td class="left row">' + rowName + '</td>' +
        '<td class="right">' + value + '</td>' +
        '</tr>'

    return row
}

function openIndex(id) {
    console.log(id)
    //window.location.replace("/create.html/" + id);

}
