
var gvar = 0;
function get_requests() {
    gvar++;
    console.log("check_gvar (get_requests):", gvar);
    return [{
       "method": "GET",
       "url": "https://httpbin.org/get",
       "headers": {
           "User-Agent": "Mozilla/4.0"
       },
       "body": "",
       "meta": {
           "host": "httpbin",
       }
    }, "https://httpbin.org/get"];
}

function parse(response) {
    gvar++;
    console.log("check_gvar (parse):", gvar);
    console.log("check_meta:", response.GetMetaValue("host") == "httpbin");
    return [response.Json()];
}

function process_item(item) {
    gvar++;
    console.log("check_gvar (process_item):", gvar);
    item.pipeline = "script";
    console.log(JSON.stringify(item));
}

function on_scheduled(request) {
    gvar++;
    console.log("check_gvar (on_scheduled):", gvar);
    console.log("on_scheduled:", request.GetURL())
}

function on_parsed_result(result) {
    gvar++;
    console.log("check_gvar (on_parsed_result):", gvar);
    console.log("parsed_result:", JSON.stringify(result));
}
