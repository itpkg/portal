function show_code() {
    $("pre code").each(function (i, block) {
        hljs.highlightBlock(block);
    });
}

function to_json(o) {
    return JSON.stringify(o, null, 2)
}

function id2url(id, prefix) {
    return id.substring(prefix.length + 1);
}
